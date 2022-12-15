package minapp

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type NotifyRouter struct{}

func (r *NotifyRouter) InitNotifyRouter(Router *gin.RouterGroup) {
	notifyPublicRouter := Router.Group("pay")

	notifyApi := v1.ApiGroupApp.MinAppApiGroup.NotifyApi

	//支付回调地址
	notifyPublicRouter.POST("notify", notifyApi.PayNotify)
}
