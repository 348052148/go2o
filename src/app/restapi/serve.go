/**
 * Copyright 2015 @ z3q.net.
 * name : rest_server.go
 * author : jarryliu
 * date : -- :
 * description :
 * history :
 */
package restapi

import (
	"github.com/jsix/gof"
	"go2o/src/core/variable"
	"gopkg.in/labstack/echo.v1"
	mw "gopkg.in/labstack/echo.v1/middleware"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	API_DOMAIN   string
	API_HOST_CHK bool = false // 必须匹配Host
	PathPrefix        = "/go2o_api_v1"
	sto          gof.Storage
	serve        *echo.Echo
)

func init() {
	serve = newServe()
}

func newServe() *echo.Echo {
	serve := echo.New()
	serve.Use(mw.Recover())
	serve.Use(beforeRequest())
	serve.Hook(splitPath) // 获取新的路径,在请求之前发生
	registerRoutes(serve)
	return serve
}

// 获取服务实例
func GetServe() *echo.Echo {
	return serve
}

func Run(app gof.App, port int) {
	sto = app.Storage()
	API_DOMAIN = app.Config().GetString(variable.ApiDomain)
	log.Println("** [ Go2o][ API][ Booted] - Api server running on port " +
		strconv.Itoa(port))
	serve.Run(":" + strconv.Itoa(port)) //启动服务
}

func registerRoutes(s *echo.Echo) {
	pc := &partnerC{}
	mc := &MemberC{}
	gc := &getC{}

	s.Get("/", ApiTest)
	s.Get("/get/invite_qr", gc.Invite_qr) // 获取二维码
	s.Get("/get/gen_qr", gc.GenQr)        //生成二维码
	s.Post("/mm_login", mc.Login)         // 会员登陆接口
	s.Post("/mm_register", mc.Register)   // 会员注册接口
	s.Post("/partner/get_ad", pc.Get_ad)  // 商户广告接口
	//s.Post("/member/*",mc)  // 会员接口
}

func beforeRequest() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx *echo.Context) error {
			host := ctx.Request().URL.Host
			path := ctx.Request().URL.Path
			// todo: path compare
			if API_HOST_CHK && host != API_DOMAIN {
				return ctx.String(http.StatusNotFound, "no such file")
			}

			if path != "/" {
				//检查商户接口权限
				ctx.Request().ParseForm()
				if !chkPartnerApiSecret(ctx) {
					return ctx.String(http.StatusOK, "{error:'secret incorrent'}")
				}
				//检查会员会话
				if strings.HasPrefix(path, "/member") && !checkMemberToken(ctx) {
					return ctx.String(http.StatusOK, "{error:'incorrent session'}")
				}
			}

			return h(ctx)
		}
	}
}

func splitPath(w http.ResponseWriter, r *http.Request) {
	preLen := len(PathPrefix)
	if len(r.URL.Path) > preLen {
		r.URL.Path = r.URL.Path[preLen:]
	}
}
