package minapp

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MineRouter struct{}

func (r *MineRouter) InitMineRouter(Router *gin.RouterGroup) {
	minePrivateRouter := Router.Group("mine").Use(middleware.JWTAuthMiddleware())
	mineApi := v1.ApiGroupApp.MinAppApiGroup.MineApi
	{
		//我发布的二手闲置
		minePrivateRouter.GET("second-order", mineApi.MyHelpSecond)
		minePrivateRouter.PUT("second", mineApi.SetMyHelpSecond)

		//我发布的团购接龙
		minePrivateRouter.GET("groupon", mineApi.MyHelpGroupon)
		//我发布的组团拼车
		minePrivateRouter.GET("carpool", mineApi.MyHelpCarpool)
		minePrivateRouter.PUT("carpool", mineApi.SetMyHelpCarpool)

		//我发布的帮我取订单
		minePrivateRouter.GET("fetch", mineApi.MyHelpFetch)
		minePrivateRouter.PUT("fetch", mineApi.SetMyHelpFetch)

		//我发布的帮我买订单
		minePrivateRouter.GET("buy", mineApi.MyHelpBuy)
		minePrivateRouter.PUT("buy", mineApi.SetMyHelpBuy)

		//我的发布，帮我做订单
		minePrivateRouter.GET("work", mineApi.MyHelpWork)
		minePrivateRouter.PUT("work", mineApi.SetMyHelpWork)

		//我的发布，帮我修订单
		minePrivateRouter.GET("repair", mineApi.MyHelpRepair)
		minePrivateRouter.PUT("repair", mineApi.SetMyHelpRepair)

		//我的发布国内捎带订单
		minePrivateRouter.GET("incidentally", mineApi.MyHelpIncidentally)
		minePrivateRouter.PUT("incidentally", mineApi.SetMyHelpIncidentally)
	}
}
