package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WithdrawRouter struct {
}

// InitWithdrawRouter 初始化 Withdraw 路由信息
func (s *WithdrawRouter) InitWithdrawRouter(Router *gin.RouterGroup) {
	withdrawRouter := Router.Group("withdraw").Use(middleware.OperationRecord())
	withdrawRouterWithoutRecord := Router.Group("withdraw")
	var withdrawApi = v1.ApiGroupApp.AutoCodeApiGroup.WithdrawApi
	{
		withdrawRouter.POST("createWithdraw", withdrawApi.CreateWithdraw)   // 新建Withdraw
		withdrawRouter.DELETE("deleteWithdraw", withdrawApi.DeleteWithdraw) // 删除Withdraw
		withdrawRouter.DELETE("deleteWithdrawByIds", withdrawApi.DeleteWithdrawByIds) // 批量删除Withdraw
		withdrawRouter.PUT("updateWithdraw", withdrawApi.UpdateWithdraw)    // 更新Withdraw
	}
	{
		withdrawRouterWithoutRecord.GET("findWithdraw", withdrawApi.FindWithdraw)        // 根据ID获取Withdraw
		withdrawRouterWithoutRecord.GET("getWithdrawList", withdrawApi.GetWithdrawList)  // 获取Withdraw列表
	}
}
