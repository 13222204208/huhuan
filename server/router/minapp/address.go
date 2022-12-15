package minapp

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AddressRouter struct{}

func (r *AddressRouter) InitAddressRouter(Router *gin.RouterGroup) {
	addressRouter := Router.Group("address").Use(middleware.JWTAuthMiddleware())

	addressApi := v1.ApiGroupApp.MinAppApiGroup.AddressApi

	{
		//地址的增删改查
		addressRouter.POST("create", addressApi.CreateAddress)
		addressRouter.POST("update", addressApi.UpdateAddress)
		addressRouter.POST("delete", addressApi.DeleteAddress)
		addressRouter.GET("list", addressApi.ListAddress)
	}
}
