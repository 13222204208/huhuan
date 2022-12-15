package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ComplaintLabelRouter struct {
}

// InitComplaintLabelRouter 初始化 ComplaintLabel 路由信息
func (s *ComplaintLabelRouter) InitComplaintLabelRouter(Router *gin.RouterGroup) {
	complaintLabelRouter := Router.Group("complaintLabel").Use(middleware.OperationRecord())
	complaintLabelRouterWithoutRecord := Router.Group("complaintLabel")
	var complaintLabelApi = v1.ApiGroupApp.AutoCodeApiGroup.ComplaintLabelApi
	{
		complaintLabelRouter.POST("createComplaintLabel", complaintLabelApi.CreateComplaintLabel)   // 新建ComplaintLabel
		complaintLabelRouter.DELETE("deleteComplaintLabel", complaintLabelApi.DeleteComplaintLabel) // 删除ComplaintLabel
		complaintLabelRouter.DELETE("deleteComplaintLabelByIds", complaintLabelApi.DeleteComplaintLabelByIds) // 批量删除ComplaintLabel
		complaintLabelRouter.PUT("updateComplaintLabel", complaintLabelApi.UpdateComplaintLabel)    // 更新ComplaintLabel
	}
	{
		complaintLabelRouterWithoutRecord.GET("findComplaintLabel", complaintLabelApi.FindComplaintLabel)        // 根据ID获取ComplaintLabel
		complaintLabelRouterWithoutRecord.GET("getComplaintLabelList", complaintLabelApi.GetComplaintLabelList)  // 获取ComplaintLabel列表
	}
}
