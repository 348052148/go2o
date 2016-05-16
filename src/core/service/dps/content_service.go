/**
 * Copyright 2015 @ z3q.net.
 * name : content_service
 * author : jarryliu
 * date : -- :
 * description :
 * history :
 */
package dps

import (
	"go2o/src/core/domain/interface/content"
	"go2o/src/core/domain/interface/merchant"
	"go2o/src/core/query"
)

type contentService struct {
	_contentRep content.IContentRep
	_query      *query.ContentQuery
}

func NewContentService(rep content.IContentRep, q *query.ContentQuery) *contentService {
	return &contentService{
		_contentRep: rep,
		_query:      q,
	}
}

// 获取页面
func (this *contentService) GetPage(partnerId, id int) *content.ValuePage {
	c := this._contentRep.GetContent(partnerId)
	page := c.GetPage(id)
	if page != nil {
		return page.GetValue()
	}
	return nil
}

// 根据标识获取页面
func (this *contentService) GetPageByIndent(partnerId int, indent string) *content.ValuePage {
	c := this._contentRep.GetContent(partnerId)
	page := c.GetPageByStringIndent(indent)
	if page != nil {
		return page.GetValue()
	}
	return nil
}

// 保存页面
func (this *contentService) SavePage(partnerId int, v *content.ValuePage) (int, error) {
	c := this._contentRep.GetContent(partnerId)
	var page content.IPage

	if v.MerchantId != partnerId {
		return -1, merchant.ErrPartnerNotMatch
	}

	if v.Id > 0 {
		page = c.GetPage(v.Id)
		page.SetValue(v)
	} else {
		page = c.CreatePage(v)
	}

	return page.Save()
}

// 删除页面
func (this *contentService) DeletePage(partnerId int, pageId int) error {
	c := this._contentRep.GetContent(partnerId)
	return c.DeletePage(pageId)
}
