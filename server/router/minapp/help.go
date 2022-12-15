package minapp

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HelpRouter struct{}

func (r *HelpRouter) InitHelpRouter(Router *gin.RouterGroup) {

	helpPrivateRouter := Router.Group("help").Use(middleware.JWTAuthMiddleware())
	helpPublicRouter := Router.Group("help")
	helpApi := v1.ApiGroupApp.MinAppApiGroup.HelpApi

	{
		//创建帮我修订单
		helpPrivateRouter.POST("create-help-repair", helpApi.CreateHelpRepair)

		//创建帮我做订单
		helpPrivateRouter.POST("create-help-work", helpApi.CreateHelpWork)

		//帮我取订单
		helpPrivateRouter.POST("create-help-fetch", helpApi.CreateHelpFetch)

		//帮我买订单
		helpPrivateRouter.POST("create-help-buy", helpApi.CreateHelpBuy)

		//创建国内捎带订单
		helpPrivateRouter.POST("create-help-incidentally", helpApi.CreateHelpIncidentally)

		//创建组团拼车
		helpPrivateRouter.POST("create-help-carpool", helpApi.CreateHelpCarpool)

		//发布团购接龙商品
		helpPrivateRouter.POST("create-help-groupon", helpApi.CreateHelpGroupon)

		//发布二手闲置商品发布
		helpPrivateRouter.POST("create-help-second", helpApi.CreateHelpSecond)

		//二手闲置的收藏和取消收藏
		helpPrivateRouter.POST("collect", helpApi.Collect)

		//二手闲置评论及留言
		helpPrivateRouter.POST("second-message", helpApi.SecondMessage)

		//获取二手闲置订单所有的留言
		helpPrivateRouter.GET("order-message", helpApi.GetSecondOrderMessage)

		//报名占位团购接龙订单
		helpPrivateRouter.POST("apply-groupon", helpApi.ApplyHelpGroupon)

		//获取二手闲置商品
		helpPrivateRouter.GET("get-help-second", helpApi.GetHelpSecond)
	}

	{
		//获取帮我修订单
		helpPublicRouter.GET("get-help-repair", helpApi.GetHelpRepair)
		//获取帮我做订单
		helpPublicRouter.GET("get-help-work", helpApi.GetHelpWork)

		//根据订单号获取订单详情
		helpPublicRouter.GET("order", helpApi.GetOrder)

		//获取帮我取订单
		helpPublicRouter.GET("get-help-fetch", helpApi.GetHelpFetch)

		//获取帮我买的订单
		helpPublicRouter.GET("get-help-buy", helpApi.GetHelpBuy)

		//获取国内捎带订单
		helpPublicRouter.GET("get-help-incidentally", helpApi.GetHelpIncidentally)

		//获取组团拼车订单
		helpPublicRouter.GET("get-help-carpool", helpApi.GetHelpCarpool)

		//获取团购接龙商品
		helpPublicRouter.GET("get-help-groupon", helpApi.GetHelpGroupon)

		//获取二手闲置分类列表
		helpPublicRouter.GET("get-category", helpApi.GetCategoryList)

		//获取团购商品信息
		helpPublicRouter.GET("category", helpApi.GetGroupByCategory)
	}

}
