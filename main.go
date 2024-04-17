// Code generated by hertz generator.

package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/opensergo/sentinel/adapter"
	"tiktok/biz/middleware/jwt"
	"tiktok/biz/router/websock"
	"tiktok/pkg/cfg"
	"tiktok/pkg/errmsg"
	"tiktok/pkg/sync"
)

func Init() error {
	err := cfg.Init()
	if err != nil {
		return err
	}
	jwt.Init()
	go sync.StartSync()
	return nil
}

func main() {

	err := Init()
	if err != nil {
		hlog.Error(err)
		return
	}

	h := server.Default(server.WithHostPorts("0.0.0.0:10001"))
	h.Use(adapter.SentinelServerMiddleware(
		adapter.WithServerResourceExtractor(func(c context.Context, ctx *app.RequestContext) string {
			return "default"
		}),
		adapter.WithServerBlockFallback(func(c context.Context, ctx *app.RequestContext) {
			ctx.AbortWithStatusJSON(429, utils.H{
				"base": utils.H{
					"code": errmsg.SentinelBlockCode,
					"msg":  errmsg.SentinelBlockMsg,
				},
			})
		}),
	))

	ws := server.Default(server.WithHostPorts("0.0.0.0:10000"))
	ws.NoHijackConnPool = true

	register(h)

	websock.WebsocketRegister(ws)
	go ws.Spin()
	h.Spin()
}
