/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2013-12-08 10:31
 * description :
 * history :
 */

package shopping

type ValueCart struct {
	Id      int    `db:"id" pk:"yes" auto:"yes"`
	CartKey string `db:"cart_key"`
	BuyerId int    `db:"buyer_id"`
	//OrderNo    string           `db:"order_no"`
	//IsBought   int              `db:"is_bought"`
	PaymentOpt int              `db:"payment_opt"`
	DeliverOpt int              `db:"deliver_opt"`
	DeliverId  int              `db:"deliver_id"`
	ShopId     int              `db:"shop_id"`
	CreateTime int64            `db:"create_time"`
	UpdateTime int64            `db:"update_time"`
	Items      []*ValueCartItem `db:"-"`
}

type ValueCartItem struct {
	Id         int     `db:"id" pk:"yes" auto:"yes"`
	CartId     int     `db:"cart_id"`
	GoodsId    int     `db:"goods_id"`
	SnapshotId int     `db:"snap_id"`
	Quantity   int     `db:"quantity"`
	Sku        string  `db:"-"`
	Price      float32 `db:"-"`
	SalePrice  float32 `db:"-"`
	Name       string  `db:"-"`
	GoodsNo    string  `db:"-"`
	SmallTitle string  `db:"-"`
	Image      string  `db:"-"`
}
