package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ReleaseRecordRouter struct {
}

// InitReleaseRecordRouter 初始化 ReleaseRecord 路由信息
func (s *ReleaseRecordRouter) InitReleaseRecordRouter(Router *gin.RouterGroup) {
	releaseRecordRouter := Router.Group("releaseRecord").Use(middleware.OperationRecord())
	releaseRecordRouterWithoutRecord := Router.Group("releaseRecord")
	var releaseRecordApi = v1.ApiGroupApp.AutoCodeApiGroup.ReleaseRecordApi
	{
		releaseRecordRouter.POST("createReleaseRecord", releaseRecordApi.CreateReleaseRecord)   // 新建ReleaseRecord
		releaseRecordRouter.DELETE("deleteReleaseRecord", releaseRecordApi.DeleteReleaseRecord) // 删除ReleaseRecord
		releaseRecordRouter.DELETE("deleteReleaseRecordByIds", releaseRecordApi.DeleteReleaseRecordByIds) // 批量删除ReleaseRecord
		releaseRecordRouter.PUT("updateReleaseRecord", releaseRecordApi.UpdateReleaseRecord)    // 更新ReleaseRecord
	}
	{
		releaseRecordRouterWithoutRecord.GET("findReleaseRecord", releaseRecordApi.FindReleaseRecord)        // 根据ID获取ReleaseRecord
		releaseRecordRouterWithoutRecord.GET("getReleaseRecordList", releaseRecordApi.GetReleaseRecordList)  // 获取ReleaseRecord列表
	}
}
