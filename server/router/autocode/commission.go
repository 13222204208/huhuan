package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CommissionRouter struct {
}

// InitCommissionRouter 初始化 Commission 路由信息
func (s *CommissionRouter) InitCommissionRouter(Router *gin.RouterGroup) {
	commissionRouter := Router.Group("commission").Use(middleware.OperationRecord())
	commissionRouterWithoutRecord := Router.Group("commission")
	var commissionApi = v1.ApiGroupApp.AutoCodeApiGroup.CommissionApi
	{
		commissionRouter.POST("createCommission", commissionApi.CreateCommission)   // 新建Commission
		commissionRouter.DELETE("deleteCommission", commissionApi.DeleteCommission) // 删除Commission
		commissionRouter.DELETE("deleteCommissionByIds", commissionApi.DeleteCommissionByIds) // 批量删除Commission
		commissionRouter.PUT("updateCommission", commissionApi.UpdateCommission)    // 更新Commission
	}
	{
		commissionRouterWithoutRecord.GET("findCommission", commissionApi.FindCommission)        // 根据ID获取Commission
		commissionRouterWithoutRecord.GET("getCommissionList", commissionApi.GetCommissionList)  // 获取Commission列表
	}
}
