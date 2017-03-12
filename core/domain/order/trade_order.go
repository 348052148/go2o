package order

import (
	"errors"
	"github.com/jsix/gof/util"
	"go2o/core/domain/interface/member"
	"go2o/core/domain/interface/merchant"
	"go2o/core/domain/interface/merchant/shop"
	"go2o/core/domain/interface/order"
	"go2o/core/domain/interface/payment"
	"go2o/core/domain/interface/valueobject"
	"strconv"
	"strings"
	"time"
)

var _ order.IOrder = new(tradeOrderImpl)
var _ order.ITradeOrder = new(tradeOrderImpl)

type tradeOrderImpl struct {
	*baseOrderImpl
	value        *order.TradeOrder
	paymentOrder payment.IPaymentOrder
	payRepo      payment.IPaymentRepo
	mchRepo      merchant.IMerchantRepo
	valueRepo    valueobject.IValueRepo
}

func newTradeOrder(base *baseOrderImpl, payRepo payment.IPaymentRepo,
	mchRepo merchant.IMerchantRepo, valueRepo valueobject.IValueRepo) order.IOrder {
	o := &tradeOrderImpl{
		baseOrderImpl: base,
		payRepo:       payRepo,
		mchRepo:       mchRepo,
		valueRepo:     valueRepo,
	}
	return o.init()
}

func (o *tradeOrderImpl) init() order.IOrder {
	o.getValue()
	return o
}

func (o *tradeOrderImpl) getValue() *order.TradeOrder {
	if o.value == nil {
		id := o.GetAggregateRootId()
		if id > 0 {
			o.value = o.repo.GetTradeOrder("order_id=?", id)
		}
	}
	return o.value
}

// 复合的订单信息
func (o *tradeOrderImpl) Complex() *order.ComplexOrder {
	v := o.getValue()
	co := o.baseOrderImpl.Complex()
	co.SubOrderId = 0
	co.VendorId = v.VendorId
	co.ShopId = v.ShopId
	co.Subject = v.Subject
	co.DiscountAmount = v.DiscountAmount
	co.ItemAmount = v.OrderAmount
	co.FinalAmount = v.FinalAmount
	co.IsBreak = 0
	co.UpdateTime = v.UpdateTime
	co.Extend["TicketImage"] = v.TicketImage
	co.Extend["TradeRate"] = strconv.FormatFloat(v.TradeRate, 'g', 2, 64)
	co.Extend["CashPay"] = strconv.Itoa(int(v.CashPay))
	return co
}

// 从订单信息中拷贝相应的数据
func (o *tradeOrderImpl) Set(v *order.ComplexOrder, rate float64) error {
	err := o.parseOrder(v, rate)
	if err == nil {
		err = o.checkRate()
	}
	return err
}

// 转换为订单相关对象
func (w *tradeOrderImpl) parseOrder(v *order.ComplexOrder, rate float64) error {
	if w.GetAggregateRootId() > 0 {
		panic(errors.New("trade order must copy before creating!"))
	}
	if v.VendorId <= 0 {
		return merchant.ErrNoSuchMerchant
	}
	if v.ShopId <= 0 {
		return shop.ErrNoSuchShop
	}
	if v.Subject == "" {
		return order.ErrMissingSubject
	}
	if v.ItemAmount <= 0 {
		return member.ErrIncorrectAmount
	}
	w.value = &order.TradeOrder{
		ID:             0,
		OrderId:        v.OrderId,
		VendorId:       v.VendorId,
		ShopId:         v.ShopId,
		Subject:        v.Subject,
		OrderAmount:    v.ItemAmount,
		DiscountAmount: v.DiscountAmount,
		FinalAmount:    0,
		TradeRate:      rate,
		State:          w.baseValue.State,
	}
	//计算最终金额
	w.fixFinalAmount()
	return nil
}

// 检查结算比例
func (w *tradeOrderImpl) checkRate() error {
	if w.value.TradeRate < 0 {
		return order.ErrTradeRateLessZero
	}
	if w.value.TradeRate > 1 {
		return order.ErrTradeRateMoreThan100
	}
	return nil
}

// 提交订单。如遇拆单,需均摊优惠抵扣金额到商品
func (o *tradeOrderImpl) Submit() error {
	if o.GetAggregateRootId() > 0 {
		return errors.New("订单不允许重复提交")
	}
	err := o.checkBuyer()
	if err == nil {
		err = o.checkRate()
	}
	if err != nil {
		return err
	}
	// 提交订单
	err = o.baseOrderImpl.Submit()
	if err == nil {
		// 保存订单信息到常规订单
		o.value.OrderId = o.GetAggregateRootId()
		o.value.State = int32(order.StatAwaitingPayment)
		o.value.CreateTime = o.baseValue.CreateTime
		o.value.UpdateTime = o.baseValue.CreateTime
		// 保存订单
		o.value.ID, err = util.I64Err(o.repo.SaveTradeOrder(o.value))
		if err == nil {
			// 生成支付单
			err = o.createPaymentForOrder()
		}
	}
	return err
}

// 检查买家及收货地址
func (o *tradeOrderImpl) checkBuyer() error {
	buyer := o.Buyer()
	if buyer == nil {
		return member.ErrNoSuchMember
	}
	if buyer.GetValue().State == 0 {
		return member.ErrMemberDisabled
	}
	return nil
}

// 计算折扣
func (o *tradeOrderImpl) applyGroupDiscount() {
	//todo: 随机立减
}

// 修正订单实际金额
func (o *tradeOrderImpl) fixFinalAmount() {
	o.value.FinalAmount = o.value.OrderAmount - o.value.DiscountAmount
}

// 生成支付单
func (o *tradeOrderImpl) createPaymentForOrder() error {
	v := o.baseOrderImpl.createPaymentOrder()
	v.VendorId = o.value.VendorId
	v.TotalFee = float32(o.value.FinalAmount)
	v.CouponDiscount = 0
	v.IntegralDiscount = 0
	v.FinalAmount = v.TotalFee - v.SubAmount - v.SystemDiscount -
		v.IntegralDiscount - v.BalanceDiscount
	o.paymentOrder = o.payRepo.CreatePaymentOrder(v)
	_, err := o.paymentOrder.Commit()
	return err
}

// 获取支付单
func (o *tradeOrderImpl) GetPaymentOrder() payment.IPaymentOrder {
	if o.paymentOrder == nil {
		id := o.GetAggregateRootId()
		if id <= 0 {
			panic(" Get payment order error ; because of order no yet created!")
		}
		o.paymentOrder = o.payRepo.GetPaymentBySalesOrderId(id)
	}
	return o.paymentOrder
}

// 现金支付
func (o *tradeOrderImpl) CashPay() error {
	py := o.GetPaymentOrder()
	pv := py.GetValue()
	switch int(pv.State) {
	case payment.StateHasCancel:
		return payment.ErrOrderHasCancel
	case payment.StateFinishPayment:
		return payment.ErrOrderPayed
	}
	v := o.getValue()
	// 商家收取现金，从商家账户扣除交易费
	vp := (1 - v.TradeRate) * float64(pv.TotalFee)
	if vp > 0 {
		va := o.mchRepo.GetMerchant(v.VendorId)
		err := va.Account().TakePayment(o.OrderNo(), vp, 0,
			"交易费-"+o.value.Subject)
		if err != nil {
			return err
		}
	}
	err := py.PaymentFinish("现金支付", "000000000")
	if err == nil {
		o.getValue()
		o.value.CashPay = 1
		return o.saveTradeOrder()
	}
	return err
}

// 在线支付交易完成,交易单付款后直接完成
func (o *tradeOrderImpl) TradePaymentFinish() error {
	o.getValue()
	if o.value.State == order.StatAwaitingPayment {
		conf := o.valueRepo.GetGlobMchSaleConf()
		// 如果交易单需要上传发票，则变为待确认。否则直接完成
		if conf.TradeOrderRequireTicket {
			if o.value.TicketImage != "" {
				return o.updateOrderComplete()
			}
			o.value.State = order.StatAwaitingConfirm
			return o.saveTradeOrder()
		}
		return o.updateOrderComplete()
	}
	return order.ErrOrderPayed
}

// 更新发票数据
func (o *tradeOrderImpl) UpdateTicket(img string) error {
	o.getValue()
	img = strings.TrimSpace(img)
	if len(img) < 10 {
		return order.ErrTicketImage
	}
	o.value.TicketImage = img
	if o.State() == order.StatAwaitingConfirm {
		return o.updateOrderComplete()
	}
	return o.saveTradeOrder()
}

func (o *tradeOrderImpl) updateOrderComplete() (err error) {
	if o.State() != order.StatCompleted {
		o.value.State = order.StatCompleted
		err := o.saveTradeOrder()
		if err == nil {
			err = o.onOrderComplete()
		}
	}
	return err
}

// 完成订单
func (o *tradeOrderImpl) onOrderComplete() error {
	err := o.updateAccountForOrder()
	if err == nil {
		err = o.vendorSettle()
	}
	return err
}

// 保存订单
func (o *tradeOrderImpl) saveTradeOrder() error {
	unix := time.Now().Unix()
	o.value.UpdateTime = unix
	if o.getValue().ID <= 0 {
		panic("please use Submit() to create new wholesale order!")
	}
	_, err := o.repo.SaveTradeOrder(o.value)
	if err == nil {
		o.syncOrderState()
	}
	return err
}

// 同步订单状态
func (o *tradeOrderImpl) syncOrderState() {
	if o.State() != order.StatBreak {
		o.saveOrderState(order.OrderState(o.value.State))
	}
}

// 商户结算
func (o *tradeOrderImpl) vendorSettle() error {
	if o.value.CashPay == 1 {
		return nil
		panic("交易单使用现金支付，不需要对商户结算!")
	}
	v := o.getValue()
	vendor := o.mchRepo.GetMerchant(v.VendorId)
	if vendor != nil {
		return o.vendorSettleByRate(vendor, v.TradeRate)
	}
	return nil
}

// 根据比例进行商户结算
func (s *tradeOrderImpl) vendorSettleByRate(vendor merchant.IMerchant, rate float64) error {
	v := s.getValue()
	sAmount := float32(v.FinalAmount * rate)
	if sAmount > 0 {
		return vendor.Account().SettleOrder(s.OrderNo(),
			sAmount, 0, 0, "交易单结算-"+v.Subject)
	}
	return nil
}

// 更新账户
func (o *tradeOrderImpl) updateAccountForOrder() error {
	if o.value.State != order.StatCompleted {
		return order.ErrUnusualOrderStat
	}
	m := o.Buyer()
	var err error
	ov := o.getValue()
	conf := o.valueRepo.GetGlobNumberConf()
	registry := o.valueRepo.GetRegistry()
	amount := ov.FinalAmount
	acc := m.GetAccount()

	// 增加经验
	if registry.MemberExperienceEnabled {
		rate := conf.ExperienceRateByOrder
		if exp := int32(amount * float64(rate)); exp > 0 {
			if err = m.AddExp(exp); err != nil {
				return err
			}
		}
	}

	// 增加积分
	//todo: 增加阶梯的返积分,比如订单满30送100积分
	iRate := float64(conf.IntegralRateByConsumption)
	integral := int64(amount*iRate) + conf.IntegralBackExtra
	// 赠送积分
	if integral > 0 {
		err = m.GetAccount().AddIntegral(member.TypeIntegralShoppingPresent,
			o.OrderNo(), integral, "")
		if err != nil {
			return err
		}
	}
	acv := acc.GetValue()
	acv.TotalExpense += float32(ov.FinalAmount)
	acv.TotalPay += float32(ov.FinalAmount)
	acv.UpdateTime = time.Now().Unix()
	_, err = acc.Save()
	return err
}
