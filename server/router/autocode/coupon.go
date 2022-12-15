package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CouponRouter struct {
}

// InitCouponRouter 初始化 Coupon 路由信息
func (s *CouponRouter) InitCouponRouter(Router *gin.RouterGroup) {
	couponRouter := Router.Group("coupon").Use(middleware.OperationRecord())
	couponRouterWithoutRecord := Router.Group("coupon")
	var couponApi = v1.ApiGroupApp.AutoCodeApiGroup.CouponApi
	{
		couponRouter.POST("createCoupon", couponApi.CreateCoupon)             // 新建Coupon
		couponRouter.DELETE("deleteCoupon", couponApi.DeleteCoupon)           // 删除Coupon
		couponRouter.DELETE("deleteCouponByIds", couponApi.DeleteCouponByIds) // 批量删除Coupon
		couponRouter.PUT("updateCoupon", couponApi.UpdateCoupon)              // 更新Coupon
	}
	{
		couponRouterWithoutRecord.GET("findCoupon", couponApi.FindCoupon)       // 根据ID获取Coupon
		couponRouterWithoutRecord.GET("getCouponList", couponApi.GetCouponList) // 获取Coupon列表

		//获取所有的优惠券
		couponRouterWithoutRecord.GET("getAllCoupons", couponApi.GetAllCouponsTitle)
	}
}
