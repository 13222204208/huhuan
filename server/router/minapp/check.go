package minapp

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CheckRouter struct{}

func (r *CheckRouter) InitCheckRouter(Router *gin.RouterGroup) {
	checkPrivateRouter := Router.Group("check").Use(middleware.JWTAuthMiddleware())
	checkApi := v1.ApiGroupApp.MinAppApiGroup.CheckApi
	{
		//司机认证
		checkPrivateRouter.POST("driver", checkApi.CreateCheckDriver)
		//商家认证
		checkPrivateRouter.POST("merchant", checkApi.CreateCheckMerchant)

		//身份认证
		checkPrivateRouter.POST("identity", checkApi.CreateCheckIdentity)
	}
}
