package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CreditRatingRouter struct {
}

// InitCreditRatingRouter 初始化 CreditRating 路由信息
func (s *CreditRatingRouter) InitCreditRatingRouter(Router *gin.RouterGroup) {
	creditRatingRouter := Router.Group("creditRating").Use(middleware.OperationRecord())
	creditRatingRouterWithoutRecord := Router.Group("creditRating")
	var creditRatingApi = v1.ApiGroupApp.AutoCodeApiGroup.CreditRatingApi
	{
		creditRatingRouter.POST("createCreditRating", creditRatingApi.CreateCreditRating)   // 新建CreditRating
		creditRatingRouter.DELETE("deleteCreditRating", creditRatingApi.DeleteCreditRating) // 删除CreditRating
		creditRatingRouter.DELETE("deleteCreditRatingByIds", creditRatingApi.DeleteCreditRatingByIds) // 批量删除CreditRating
		creditRatingRouter.PUT("updateCreditRating", creditRatingApi.UpdateCreditRating)    // 更新CreditRating
	}
	{
		creditRatingRouterWithoutRecord.GET("findCreditRating", creditRatingApi.FindCreditRating)        // 根据ID获取CreditRating
		creditRatingRouterWithoutRecord.GET("getCreditRatingList", creditRatingApi.GetCreditRatingList)  // 获取CreditRating列表
	}
}
