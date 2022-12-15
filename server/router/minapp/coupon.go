package minapp

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CouponRouter struct{}

func (r *CouponRouter) InitCouponRouter(Router *gin.RouterGroup) {
	couponPrivateRouter := Router.Group("coupon").Use(middleware.JWTAuthMiddleware())
	couponApi := v1.ApiGroupApp.MinAppApiGroup.CouponApi
	{
		couponPrivateRouter.GET("my-coupon", couponApi.MyCoupon)
		//充值送的优惠券列表
		couponPrivateRouter.GET("pay-give", couponApi.GetPayCouponList)

		//邀请用户赠送优惠券
		couponPrivateRouter.GET("invite", couponApi.GetGiveCoupon)
	}
}
