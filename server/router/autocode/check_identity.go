package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CheckIdentityRouter struct {
}

// InitCheckIdentityRouter 初始化 CheckIdentity 路由信息
func (s *CheckIdentityRouter) InitCheckIdentityRouter(Router *gin.RouterGroup) {
	checkIdentityRouter := Router.Group("checkIdentity").Use(middleware.OperationRecord())
	checkIdentityRouterWithoutRecord := Router.Group("checkIdentity")
	var checkIdentityApi = v1.ApiGroupApp.AutoCodeApiGroup.CheckIdentityApi
	{
		checkIdentityRouter.POST("createCheckIdentity", checkIdentityApi.CreateCheckIdentity)   // 新建CheckIdentity
		checkIdentityRouter.DELETE("deleteCheckIdentity", checkIdentityApi.DeleteCheckIdentity) // 删除CheckIdentity
		checkIdentityRouter.DELETE("deleteCheckIdentityByIds", checkIdentityApi.DeleteCheckIdentityByIds) // 批量删除CheckIdentity
		checkIdentityRouter.PUT("updateCheckIdentity", checkIdentityApi.UpdateCheckIdentity)    // 更新CheckIdentity
	}
	{
		checkIdentityRouterWithoutRecord.GET("findCheckIdentity", checkIdentityApi.FindCheckIdentity)        // 根据ID获取CheckIdentity
		checkIdentityRouterWithoutRecord.GET("getCheckIdentityList", checkIdentityApi.GetCheckIdentityList)  // 获取CheckIdentity列表
	}
}
