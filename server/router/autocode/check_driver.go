package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CheckDriverRouter struct {
}

// InitCheckDriverRouter 初始化 CheckDriver 路由信息
func (s *CheckDriverRouter) InitCheckDriverRouter(Router *gin.RouterGroup) {
	checkDriverRouter := Router.Group("checkDriver").Use(middleware.OperationRecord())
	checkDriverRouterWithoutRecord := Router.Group("checkDriver")
	var checkDriverApi = v1.ApiGroupApp.AutoCodeApiGroup.CheckDriverApi
	{
		checkDriverRouter.POST("createCheckDriver", checkDriverApi.CreateCheckDriver)   // 新建CheckDriver
		checkDriverRouter.DELETE("deleteCheckDriver", checkDriverApi.DeleteCheckDriver) // 删除CheckDriver
		checkDriverRouter.DELETE("deleteCheckDriverByIds", checkDriverApi.DeleteCheckDriverByIds) // 批量删除CheckDriver
		checkDriverRouter.PUT("updateCheckDriver", checkDriverApi.UpdateCheckDriver)    // 更新CheckDriver
	}
	{
		checkDriverRouterWithoutRecord.GET("findCheckDriver", checkDriverApi.FindCheckDriver)        // 根据ID获取CheckDriver
		checkDriverRouterWithoutRecord.GET("getCheckDriverList", checkDriverApi.GetCheckDriverList)  // 获取CheckDriver列表
	}
}
