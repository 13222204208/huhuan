package minapp

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BannerRouter struct{}

func (r *BannerRouter) InitBannerRouter(Router *gin.RouterGroup) {
	bannerPublicRouter := Router.Group("banner")
	bannerApi := v1.ApiGroupApp.MinAppApiGroup.BannerApi

	{
		//轮播图
		bannerPublicRouter.GET("list", bannerApi.GetBannerList)
	}
}
