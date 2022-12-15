package minapp

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ComplaintRouter struct{}

func (r *ComplaintRouter) InitComplaintRouter(Router *gin.RouterGroup) {
	complaintPublicRouter := Router.Group("complaint")
	complaintPrivateRouter := Router.Group("complaint").Use(middleware.JWTAuthMiddleware())
	complaintApi := v1.ApiGroupApp.MinAppApiGroup.ComplaintApi
	{
		//投诉标签
		complaintPublicRouter.GET("labels", complaintApi.Labels)
	}

	{
		//投诉订单
		complaintPrivateRouter.POST("create", complaintApi.Create)
	}
}
