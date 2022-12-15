package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ComplaintRouter struct {
}

// InitComplaintRouter 初始化 Complaint 路由信息
func (s *ComplaintRouter) InitComplaintRouter(Router *gin.RouterGroup) {
	complaintRouter := Router.Group("complaint").Use(middleware.OperationRecord())
	complaintRouterWithoutRecord := Router.Group("complaint")
	var complaintApi = v1.ApiGroupApp.AutoCodeApiGroup.ComplaintApi
	{
		complaintRouter.POST("createComplaint", complaintApi.CreateComplaint)   // 新建Complaint
		complaintRouter.DELETE("deleteComplaint", complaintApi.DeleteComplaint) // 删除Complaint
		complaintRouter.DELETE("deleteComplaintByIds", complaintApi.DeleteComplaintByIds) // 批量删除Complaint
		complaintRouter.PUT("updateComplaint", complaintApi.UpdateComplaint)    // 更新Complaint
	}
	{
		complaintRouterWithoutRecord.GET("findComplaint", complaintApi.FindComplaint)        // 根据ID获取Complaint
		complaintRouterWithoutRecord.GET("getComplaintList", complaintApi.GetComplaintList)  // 获取Complaint列表
	}
}
