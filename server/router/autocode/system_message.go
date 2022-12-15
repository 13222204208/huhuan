package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SystemMessageRouter struct {
}

// InitSystemMessageRouter 初始化 SystemMessage 路由信息
func (s *SystemMessageRouter) InitSystemMessageRouter(Router *gin.RouterGroup) {
	systemMessageRouter := Router.Group("systemMessage").Use(middleware.OperationRecord())
	systemMessageRouterWithoutRecord := Router.Group("systemMessage")
	var systemMessageApi = v1.ApiGroupApp.AutoCodeApiGroup.SystemMessageApi
	{
		systemMessageRouter.POST("createSystemMessage", systemMessageApi.CreateSystemMessage)   // 新建SystemMessage
		systemMessageRouter.DELETE("deleteSystemMessage", systemMessageApi.DeleteSystemMessage) // 删除SystemMessage
		systemMessageRouter.DELETE("deleteSystemMessageByIds", systemMessageApi.DeleteSystemMessageByIds) // 批量删除SystemMessage
		systemMessageRouter.PUT("updateSystemMessage", systemMessageApi.UpdateSystemMessage)    // 更新SystemMessage
	}
	{
		systemMessageRouterWithoutRecord.GET("findSystemMessage", systemMessageApi.FindSystemMessage)        // 根据ID获取SystemMessage
		systemMessageRouterWithoutRecord.GET("getSystemMessageList", systemMessageApi.GetSystemMessageList)  // 获取SystemMessage列表
	}
}
