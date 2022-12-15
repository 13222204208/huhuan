package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CheckMerchantRouter struct {
}

// InitCheckMerchantRouter 初始化 CheckMerchant 路由信息
func (s *CheckMerchantRouter) InitCheckMerchantRouter(Router *gin.RouterGroup) {
	checkMerchantRouter := Router.Group("checkMerchant").Use(middleware.OperationRecord())
	checkMerchantRouterWithoutRecord := Router.Group("checkMerchant")
	var checkMerchantApi = v1.ApiGroupApp.AutoCodeApiGroup.CheckMerchantApi
	{
		checkMerchantRouter.POST("createCheckMerchant", checkMerchantApi.CreateCheckMerchant)   // 新建CheckMerchant
		checkMerchantRouter.DELETE("deleteCheckMerchant", checkMerchantApi.DeleteCheckMerchant) // 删除CheckMerchant
		checkMerchantRouter.DELETE("deleteCheckMerchantByIds", checkMerchantApi.DeleteCheckMerchantByIds) // 批量删除CheckMerchant
		checkMerchantRouter.PUT("updateCheckMerchant", checkMerchantApi.UpdateCheckMerchant)    // 更新CheckMerchant
	}
	{
		checkMerchantRouterWithoutRecord.GET("findCheckMerchant", checkMerchantApi.FindCheckMerchant)        // 根据ID获取CheckMerchant
		checkMerchantRouterWithoutRecord.GET("getCheckMerchantList", checkMerchantApi.GetCheckMerchantList)  // 获取CheckMerchant列表
	}
}
