package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PayCouponRouter struct {
}

// InitPayCouponRouter 初始化 PayCoupon 路由信息
func (s *PayCouponRouter) InitPayCouponRouter(Router *gin.RouterGroup) {
	payCouponRouter := Router.Group("payCoupon").Use(middleware.OperationRecord())
	payCouponRouterWithoutRecord := Router.Group("payCoupon")
	var payCouponApi = v1.ApiGroupApp.AutoCodeApiGroup.PayCouponApi
	{
		payCouponRouter.POST("createPayCoupon", payCouponApi.CreatePayCoupon)   // 新建PayCoupon
		payCouponRouter.DELETE("deletePayCoupon", payCouponApi.DeletePayCoupon) // 删除PayCoupon
		payCouponRouter.DELETE("deletePayCouponByIds", payCouponApi.DeletePayCouponByIds) // 批量删除PayCoupon
		payCouponRouter.PUT("updatePayCoupon", payCouponApi.UpdatePayCoupon)    // 更新PayCoupon
	}
	{
		payCouponRouterWithoutRecord.GET("findPayCoupon", payCouponApi.FindPayCoupon)        // 根据ID获取PayCoupon
		payCouponRouterWithoutRecord.GET("getPayCouponList", payCouponApi.GetPayCouponList)  // 获取PayCoupon列表
	}
}
