package minapp

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SignRouter struct{}

func (r *SignRouter) InitSignRouter(Router *gin.RouterGroup) {
	signPrivateRouter := Router.Group("sign").Use(middleware.JWTAuthMiddleware())
	signApi := v1.ApiGroupApp.MinAppApiGroup.SignApi
	{
		//更改是否弹出签到通知
		signPrivateRouter.POST("change-state", signApi.ChangeSignInState)

		//用户签到
		signPrivateRouter.POST("in", signApi.CreateUserSignIn)

		//签到页面，可用积分兑换的优惠券数
		signPrivateRouter.GET("coupon", signApi.GetIntegralCouponList)

		//积分兑换优惠券
		signPrivateRouter.POST("integral-exchange-coupon", signApi.ExchangeCoupon)

		//我的签到记录
		signPrivateRouter.GET("record", signApi.GetMySignInRecords)
	}
}
