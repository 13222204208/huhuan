package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MinappSetRouter struct {
}

// InitMinappSetRouter 初始化 MinappSet 路由信息
func (s *MinappSetRouter) InitMinappSetRouter(Router *gin.RouterGroup) {
	minappSetRouter := Router.Group("minappSet").Use(middleware.OperationRecord())
	minappSetRouterWithoutRecord := Router.Group("minappSet")
	var minappSetApi = v1.ApiGroupApp.AutoCodeApiGroup.MinappSetApi
	{
		minappSetRouter.POST("createMinappSet", minappSetApi.CreateMinappSet)   // 新建MinappSet
		minappSetRouter.DELETE("deleteMinappSet", minappSetApi.DeleteMinappSet) // 删除MinappSet
		minappSetRouter.DELETE("deleteMinappSetByIds", minappSetApi.DeleteMinappSetByIds) // 批量删除MinappSet
		minappSetRouter.PUT("updateMinappSet", minappSetApi.UpdateMinappSet)    // 更新MinappSet
	}
	{
		minappSetRouterWithoutRecord.GET("findMinappSet", minappSetApi.FindMinappSet)        // 根据ID获取MinappSet
		minappSetRouterWithoutRecord.GET("getMinappSetList", minappSetApi.GetMinappSetList)  // 获取MinappSet列表
	}
}
