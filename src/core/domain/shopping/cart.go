package shopping

import (
	"bytes"
	"encoding/json"
	"go2o/src/core/domain/interface/member"
	"go2o/src/core/domain/interface/merchant"
	"go2o/src/core/domain/interface/sale"
	"go2o/src/core/domain/interface/shopping"
	"go2o/src/core/domain/interface/valueobject"
	"go2o/src/core/infrastructure/domain"
	"strconv"
	"time"
)

type Cart struct {
	_value       *shopping.ValueCart
	_saleRep     sale.ISaleRep
	_goodsRep    sale.IGoodsRep
	_shoppingRep shopping.IShoppingRep
	_partnerRep  merchant.IMerchantRep
	_memberRep   member.IMemberRep
	_partnerId   int
	_summary     string
	_shop        merchant.IShop
	_deliver     member.IDeliver
}

func createCart(partnerRep merchant.IMerchantRep, memberRep member.IMemberRep, saleRep sale.ISaleRep,
	goodsRep sale.IGoodsRep, shoppingRep shopping.IShoppingRep, partnerId int,
	val *shopping.ValueCart) shopping.ICart {
	return (&Cart{
		_value:       val,
		_partnerId:   partnerId,
		_partnerRep:  partnerRep,
		_memberRep:   memberRep,
		_shoppingRep: shoppingRep,
		_saleRep:     saleRep,
		_goodsRep:    goodsRep,
	}).init()
}

//todo: partnerId 应去掉，可能在多个商家买东西
func newCart(partnerRep merchant.IMerchantRep, memberRep member.IMemberRep, saleRep sale.ISaleRep,
	goodsRep sale.IGoodsRep, shoppingRep shopping.IShoppingRep, partnerId int, buyerId int) shopping.ICart {
	unix := time.Now().Unix()
	cartKey := domain.GenerateCartKey(unix, time.Now().Nanosecond())
	value := &shopping.ValueCart{
		CartKey:    cartKey,
		BuyerId:    buyerId,
		ShopId:     0,
		DeliverId:  0,
		PaymentOpt: 1,
		DeliverOpt: 1,
		CreateTime: unix,
		UpdateTime: unix,
		Items:      nil,
	}

	return (&Cart{
		_value:       value,
		_partnerRep:  partnerRep,
		_memberRep:   memberRep,
		_partnerId:   partnerId,
		_shoppingRep: shoppingRep,
		_saleRep:     saleRep,
		_goodsRep:    goodsRep,
	}).init()
}

func (this *Cart) init() shopping.ICart {
	// 初始化购物车的信息
	if this._value != nil && this._value.Items != nil {
		this.setAttachGoodsInfo(this._value.Items)
	}
	return this
}

// 设置附加的商品信息
func (this *Cart) setAttachGoodsInfo(items []*shopping.ValueCartItem) {
	if items != nil {
		l := len(items)
		if l == 0 {
			return
		}
		var ids []int = make([]int, l)
		for i, v := range items {
			ids[i] = v.GoodsId
		}

		// 设置附加的值
		goods, err := this._goodsRep.GetGoodsByIds(ids...)
		if err == nil {
			var goodsMap = make(map[int]*valueobject.Goods, len(goods))
			for _, v := range goods {
				goodsMap[v.GoodsId] = v
			}

			var level int
			var goods sale.IGoods
			var sl sale.ISale

			//  更新登陆后的优惠价
			if this._value.BuyerId > 0 {
				sl = this._saleRep.GetSale(this._partnerId)
				m := this._memberRep.GetMember(this._value.BuyerId)
				if m != nil {
					level = m.GetValue().Level
				}
			}

			for _, v := range items {
				gv, ok := goodsMap[v.GoodsId]
				if level > 0 {
					goods = sl.CreateGoodsByItem(
						sl.CreateItem(sale.ParseToPartialValueItem(gv)),
						sale.ParseToValueGoods(gv),
					)
					if p := goods.GetPromotionPrice(level); p < gv.SalePrice {
						gv.SalePrice = p
					}
				}
				if ok {
					v.Name = gv.Name
					v.Price = gv.Price
					v.GoodsNo = gv.GoodsNo
					v.Image = gv.Image
					v.SalePrice = gv.SalePrice
				}
			}
		}
	}
}

func (this *Cart) GetDomainId() int {
	return this._value.Id
}

func (this *Cart) GetValue() shopping.ValueCart {
	return *this._value
}

// 获取购物车中的商品
func (this *Cart) GetCartGoods() []sale.IGoods {
	sl := this._saleRep.GetSale(this._partnerId)
	var gs []sale.IGoods = make([]sale.IGoods, len(this._value.Items))
	for i, v := range this._value.Items {
		gs[i] = sl.GetGoods(v.GoodsId)
	}
	return gs
}

// 添加项
func (this *Cart) AddItem(goodsId, num int) (*shopping.ValueCartItem, error) {
	var err error
	if this._value.Items == nil {
		this._value.Items = []*shopping.ValueCartItem{}
	}
	sl := this._saleRep.GetSale(this._partnerId)
	goods := sl.GetGoods(goodsId)
	if goods == nil {
		return nil, sale.ErrNoSuchGoods // 没有商品
	}

	if !goods.GetItem().IsOnShelves() {
		return nil, sale.ErrNotOnShelves //未上架
	}

	stockNum := goods.GetValue().StockNum
	if stockNum == 0 {
		return nil, sale.ErrFullOfStock // 已经卖完了
	}

	// 添加数量
	for _, v := range this._value.Items {
		if v.GoodsId == goodsId {
			if v.Quantity+num > stockNum {
				return v, sale.ErrOutOfStock // 库存不足
			}
			v.Quantity += num
			return v, err
		}
	}

	gv := goods.GetPackedValue()
	snap := goods.GetLatestSnapshot()

	if snap == nil {
		return nil, sale.ErrNoSuchSnapshot
	}

	v := &shopping.ValueCartItem{
		CartId:     this.GetDomainId(),
		SnapshotId: snap.Id,
		GoodsId:    goodsId,
		Quantity:   num,
		Name:       gv.Name,
		GoodsNo:    gv.GoodsNo,
		Image:      gv.Image,
		Price:      gv.Price,
		SalePrice:  gv.PromPrice, // 使用优惠价
	}
	this._value.Items = append(this._value.Items, v)
	return v, err
}

// 移出项
func (this *Cart) RemoveItem(goodsId, num int) error {
	if this._value.Items == nil {
		return shopping.ErrEmptyShoppingCart
	}

	// 删除数量
	for _, v := range this._value.Items {
		if v.GoodsId == goodsId {
			if newNum := v.Quantity - num; newNum <= 0 {
				// 移出购物车
				//this.value.Items = append(this.value.Items[:i],this.value.Items[i+1:]...)
				v.Quantity = 0
			} else {
				v.Quantity = newNum
			}
			break
		}
	}
	return nil
}

// 合并购物车，并返回新的购物车
func (this *Cart) Combine(c shopping.ICart) (shopping.ICart, error) {
	if c.GetDomainId() != this.GetDomainId() {
		for _, v := range c.GetValue().Items {
			this.AddItem(v.GoodsId, v.Quantity)
		}
	}
	return this, nil
}

// 设置购买会员
func (this *Cart) SetBuyer(buyerId int) error {
	if this._value.BuyerId > 0 {
		return shopping.ErrCartBuyerBinded
	}
	this._value.BuyerId = buyerId
	_, err := this.Save()
	return err
}

// 结算数据持久化
func (this *Cart) SettlePersist(shopId, paymentOpt, deliverOpt, deliverId int) error {
	var shop merchant.IShop
	var deliver member.IDeliver
	var err error

	if shopId > 0 {
		var pt merchant.IMerchant
		pt, err = this._partnerRep.GetMerchant(this._partnerId)
		if err != nil {
			return err
		}
		shop = pt.GetShop(shopId)
		if shop == nil {
			return merchant.ErrNoSuchShop
		}
		this._shop = shop
		this._value.ShopId = shopId
	}

	if this._value.BuyerId > 0 && deliverId > 0 {
		var m member.IMember = this._memberRep.GetMember(this._value.BuyerId)
		if m == nil {
			return member.ErrNoSuchMember
		}
		deliver = m.GetDeliver(deliverId)
		if deliver == nil {
			return member.ErrInvalidSession
		}
		this._deliver = deliver
		this._value.DeliverId = deliverId
	}

	this._value.PaymentOpt = paymentOpt
	this._value.DeliverOpt = deliverOpt
	return nil
}

// 获取结算数据
func (this *Cart) GetSettleData() (s merchant.IShop, d member.IDeliver, paymentOpt, deliverOpt int) {
	var err error
	if this._value.ShopId > 0 && this._shop == nil {
		var pt merchant.IMerchant
		pt, err = this._partnerRep.GetMerchant(this._partnerId)
		if err == nil {
			this._shop = pt.GetShop(this._value.ShopId)
		}
	}
	if this._value.DeliverId > 0 && this._deliver == nil {
		var m member.IMember
		m = this._memberRep.GetMember(this._value.BuyerId)
		if m != nil {
			this._deliver = m.GetDeliver(this._value.DeliverId)
		}
	}
	return this._shop, this._deliver, this._value.PaymentOpt, this._value.DeliverOpt
}

// 保存购物车
func (this *Cart) Save() (int, error) {
	rep := this._shoppingRep
	this._value.UpdateTime = time.Now().Unix()
	id, err := rep.SaveShoppingCart(this._value)
	this._value.Id = id

	if this._value.Items != nil {
		for _, v := range this._value.Items {
			if v.Quantity <= 0 {
				rep.RemoveCartItem(v.Id)
			} else {
				i, err := rep.SaveCartItem(v)
				if err != nil {
					v.Id = i
				}
			}
		}
	}

	return id, err
}

// 销毁购物车
func (this *Cart) Destroy() (err error) {
	if err = this._shoppingRep.EmptyCartItems(this.GetDomainId()); err == nil {
		return this._shoppingRep.DeleteCart(this.GetDomainId())
	}
	return err
}

// 获取总览信息
func (this *Cart) GetSummary() string {
	if len(this._summary) != 0 {
		return this._summary
	}
	buf := bytes.NewBufferString("")
	length := len(this._value.Items)

	var snap *sale.GoodsSnapshot
	for i, v := range this._value.Items {

		snap = this._saleRep.GetGoodsSnapshot(v.SnapshotId)
		if snap != nil {
			buf.WriteString(snap.GoodsName)
			if len(snap.SmallTitle) != 0 {
				buf.WriteString("(" + snap.SmallTitle + ")")
			}
			buf.WriteString("*" + strconv.Itoa(v.Quantity))
			if i < length-1 {
				buf.WriteString("\n")
			}
		}
	}
	return buf.String()
}

// 获取Json格式的商品数据
func (this *Cart) GetJsonItems() []byte {
	var goods []*shopping.OrderGoods = make([]*shopping.OrderGoods, len(this._value.Items))
	for i, v := range this._value.Items {
		goods[i] = &shopping.OrderGoods{
			GoodsId:    v.GoodsId,
			GoodsImage: v.Image,
			Quantity:   v.Quantity,
			Name:       v.Name,
		}
	}
	d, _ := json.Marshal(goods)
	return d
}

// 获取订单金额,返回totalFee为总额，
// orderFee为实际订单的金额(扣去促销优惠等后的金额)
func (this *Cart) GetFee() (totalFee float32, orderFee float32) {
	var qua float32
	for _, v := range this._value.Items {
		qua = float32(v.Quantity)
		totalFee = totalFee + v.Price*qua
		orderFee = orderFee + v.SalePrice*qua
	}
	return totalFee, orderFee
}
