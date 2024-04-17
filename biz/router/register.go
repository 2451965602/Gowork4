// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	interact "tiktok/biz/router/interact"
	model "tiktok/biz/router/model"
	social "tiktok/biz/router/social"
	user "tiktok/biz/router/user"
	video "tiktok/biz/router/video"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	video.Register(r)

	user.Register(r)

	model.Register(r)

	social.Register(r)

	interact.Register(r)
}
