/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2013-12-12 17:16
 * description :
 * history :
 */

package repository

import (
	"encoding/json"
	"fmt"
	"github.com/jsix/gof/db"
	"github.com/jsix/gof/db/orm"
	"github.com/jsix/gof/storage"
	"go2o/core/domain/interface/merchant"
	"go2o/core/domain/interface/merchant/shop"
)

var _ shop.IShopRepo = new(shopRepo)

type shopRepo struct {
	db.Connector
	storage storage.Interface
}

func NewShopRepo(c db.Connector, storage storage.Interface) shop.IShopRepo {
	return &shopRepo{
		Connector: c,
		storage:   storage,
	}
}

// 商店别名是否存在
func (s *shopRepo) ShopAliasExists(alias string, shopId int32) bool {
	num := 0
	s.Connector.ExecScalar(`SELECT COUNT(0) FROM mch_online_shop WHERE
		alias=? AND shop_id<>?`, &num, alias, shopId)
	return num > 0
}

// 获取线上商店
func (s *shopRepo) GetOnlineShop(shopId int32) *shop.OnlineShop {
	e := shop.OnlineShop{}
	if s.GetOrm().Get(shopId, &e) != nil {
		return nil
	}
	return &e
}

// 保存线上商店
func (s *shopRepo) SaveOnlineShop(v *shop.OnlineShop, create bool) error {
	var err error
	if create {
		_, _, err = s.GetOrm().Save(nil, v)
	} else {
		_, _, err = s.GetOrm().Save(v.ShopId, v)
	}
	return err
}

// 获取线下商店
func (s *shopRepo) GetOfflineShop(shopId int32) *shop.OfflineShop {
	e := shop.OfflineShop{}
	if s.GetOrm().Get(shopId, &e) != nil {
		return nil
	}
	return &e
}

// 保存线下商店
func (s *shopRepo) SaveOfflineShop(v *shop.OfflineShop, create bool) error {
	var err error
	if create {
		_, _, err = s.GetOrm().Save(nil, v)
	} else {
		_, _, err = s.GetOrm().Save(v.ShopId, v)
	}
	return err
}

// 保存API信息
func (s *shopRepo) SaveApiInfo(v *merchant.ApiInfo) error {
	_, err := orm.Save(s.GetOrm(), v, int(v.MerchantId))
	return err
}

// 获取API信息
func (s *shopRepo) GetApiInfo(mchId int32) *merchant.ApiInfo {
	var d *merchant.ApiInfo = new(merchant.ApiInfo)
	if err := s.GetOrm().Get(mchId, d); err == nil {
		return d
	}
	return nil
}

func (s *shopRepo) SaveShop(v *shop.Shop) (int32, error) {
	id, err := orm.I32(orm.Save(s.GetOrm(), v, int(v.Id)))
	if err == nil {
		s.delCache(v.VendorId)
	}
	return id, err
}

func (s *shopRepo) GetValueShop(mchId, shopId int32) *shop.Shop {
	var v *shop.Shop = new(shop.Shop)
	err := s.Connector.GetOrm().Get(shopId, v)
	if err == nil &&
		v.VendorId == mchId {
		return v
	} else {
		handleError(err)
	}
	return nil
}

func (s *shopRepo) delCache(mchId int32) {
	PrefixDel(s.storage, fmt.Sprintf("go2o:rep:shop:%d:*", mchId))
}

func (s *shopRepo) getShopCacheKey(mchId int32) string {
	return fmt.Sprintf("go2o:rep:shop:%d:shops", mchId)
}

func (s *shopRepo) GetShopsOfMerchant(mchId int32) []shop.Shop {
	shops := []shop.Shop{}
	key := s.getShopCacheKey(mchId)
	jsonStr, err := s.storage.GetString(key)
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &shops)
	}
	if err != nil {
		err = s.Connector.GetOrm().SelectByQuery(&shops,
			"SELECT * FROM mch_shop WHERE mch_id=?", mchId)
		if err != nil {
			handleError(err)
			return nil
		} else {
			b, _ := json.Marshal(shops)
			s.storage.Set(key, string(b))
		}
	}
	return shops
}

func (s *shopRepo) deleteShop(mchId, shopId int32) error {
	_, err := s.Connector.GetOrm().Delete(shop.Shop{},
		"mch_id=? AND id=?", mchId, shopId)
	s.delCache(mchId)
	return err
}

// 删除线上商店
func (s *shopRepo) DeleteOnlineShop(mchId, shopId int32) error {
	err := s.deleteShop(mchId, shopId)
	if err == nil {
		err = s.Connector.GetOrm().DeleteByPk(shop.OnlineShop{}, shopId)
		s.delCache(mchId)
	}
	return err
}

// 删除线下门店
func (s *shopRepo) DeleteOfflineShop(mchId, shopId int32) error {
	err := s.deleteShop(mchId, shopId)
	if err == nil {
		err = s.Connector.GetOrm().DeleteByPk(shop.OfflineShop{}, shopId)
		s.delCache(mchId)
	}
	return err
}
