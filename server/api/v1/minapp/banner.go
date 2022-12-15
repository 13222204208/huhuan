package minapp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	minAppRequest "github.com/flipped-aurora/gin-vue-admin/server/model/minapp/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BannerApi struct{}

func (bannerApi *BannerApi) GetBannerList(c *gin.Context) {
	var b minAppRequest.BannerList
	_ = c.ShouldBindQuery(&b)

	if err := utils.Verify(b, utils.MinAppBannerVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, result := bannerService.GetBannerList(b.City, b.Class); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"bannerList": result}, "获取成功", c)
	}
}
