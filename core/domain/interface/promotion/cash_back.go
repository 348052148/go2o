/**
 * Copyright 2015 @ z3q.net.
 * name : cash_back
 * author : jarryliu
 * date : -- :
 * description :
 * history :
 */
package promotion

// 返现促销
type ICashBackPromotion interface {
	// 获取领域编号
	GetDomainId() int

	// 设置详细的促销信息
	SetDetailsValue(*ValueCashBack) error

	// 获取自定义数据
	GetDataTag() map[string]string
}
