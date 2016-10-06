/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2014-02-11 21:15
 * description :
 * history :
 */
package format

import (
	"go2o/core/infrastructure"
	"go2o/core/variable"
	"strconv"
	"strings"
)

var (
	imageServe   string
	noPicUrl     string
	_noPicUrl    string
	picCfgLoaded bool
)

//todo: not used
// 格式化商品编号，不足位用０补齐
func FormatGoodsNo(d int) string {
	const l int = 6
	s := strconv.Itoa(d)
	sl := len(s)
	if sl >= 6 {
		return s
	}
	return strings.Repeat("0", l-sl) + s
}

// 获取无图片地址
func GetNoPicPath() string {
	if len(_noPicUrl) == 0 {
		ctx := infrastructure.GetApp()
		_noPicUrl = ctx.Config().GetString(variable.NoPicPath)
	}
	return _noPicUrl
}

// 获取商品图片地址
func GetGoodsImageUrl(image string) string {
	if !picCfgLoaded {
		ctx := infrastructure.GetApp()
		if len(imageServe) == 0 {
			imageServe = ctx.Config().GetString(variable.ImageServer)
		}

		if len(noPicUrl) == 0 {
			noPicUrl = imageServe + "/" + ctx.Config().GetString(variable.NoPicPath)
		}
		picCfgLoaded = true
	}

	if len(image) == 0 {
		return noPicUrl
	}

	if strings.HasPrefix(image, "http://") || strings.HasPrefix(image, "https://") {
		return image
	}
	return imageServe + "/" + image
}

// 获取资源地址
func GetResUrl(image string) string {
	if !picCfgLoaded {
		ctx := infrastructure.GetApp()
		if len(imageServe) == 0 {
			imageServe = ctx.Config().GetString(variable.ImageServer)
		}

		if len(noPicUrl) == 0 {
			noPicUrl = imageServe + "/" + ctx.Config().GetString(variable.NoPicPath)
		}
		picCfgLoaded = true
	}

	if len(image) == 0 {
		return noPicUrl
	}

	if strings.HasPrefix(image, "http://") || strings.HasPrefix(image, "https://") {
		return image
	}
	return imageServe + "/" + image
}
