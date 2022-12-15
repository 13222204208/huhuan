package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MinappUserRouter struct {
}

// InitMinappUserRouter 初始化 MinappUser 路由信息
func (s *MinappUserRouter) InitMinappUserRouter(Router *gin.RouterGroup) {
	minappUserRouter := Router.Group("minappUser").Use(middleware.OperationRecord())
	minappUserRouterWithoutRecord := Router.Group("minappUser")
	var minappUserApi = v1.ApiGroupApp.AutoCodeApiGroup.MinappUserApi
	{
		minappUserRouter.POST("createMinappUser", minappUserApi.CreateMinappUser)             // 新建MinappUser
		minappUserRouter.DELETE("deleteMinappUser", minappUserApi.DeleteMinappUser)           // 删除MinappUser
		minappUserRouter.DELETE("deleteMinappUserByIds", minappUserApi.DeleteMinappUserByIds) // 批量删除MinappUser
		minappUserRouter.PUT("updateMinappUser", minappUserApi.UpdateMinappUser)              // 更新MinappUser
		//获取新增用户的每天数据
		minappUserRouter.GET("getNewUserCountList", minappUserApi.GetNewUserCount)
	}
	{
		minappUserRouterWithoutRecord.GET("findMinappUser", minappUserApi.FindMinappUser)       // 根据ID获取MinappUser
		minappUserRouterWithoutRecord.GET("getMinappUserList", minappUserApi.GetMinappUserList) // 获取MinappUser列表
	}
}
