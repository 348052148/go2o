/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2013-12-08 10:23
 * description :
 * history :
 */

package shopping

import (
	"go2o/core/domain/interface/member"
	"go2o/core/domain/interface/merchant"
	"go2o/core/domain/interface/sale"
)

type ICart interface {
	GetDomainId() int

	// 获取购物车值
	GetValue() ValueCart

	// 获取购物车中的商品
	GetCartGoods() []sale.IGoods

	// 结算数据持久化
	SettlePersist(shopId, paymentOpt, deliverOpt, deliverId int) error

	// 获取结算数据
	GetSettleData() (s merchant.IShop, d member.IDeliver, paymentOpt, deliverOpt int)

	// 设置购买会员
	SetBuyer(buyerId int) error

	// 添加项
	AddItem(goodsId, num int) (*ValueCartItem, error)

	// 移出项
	RemoveItem(goodsId, num int) error

	// 合并购物车，并返回新的购物车
	Combine(ICart) (ICart, error)

	// 保存购物车
	Save() (int, error)

	// 销毁购物车
	Destroy() error

	// 获取汇总信息
	GetSummary() string

	// 获取Json格式的商品数据
	GetJsonItems() []byte

	// 获取金额
	GetFee() (totalFee float32, orderFee float32)
}
