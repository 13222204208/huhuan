package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type BannerApi struct {
}

var bannerService = service.ServiceGroupApp.AutoCodeServiceGroup.BannerService


// CreateBanner 创建Banner
// @Tags Banner
// @Summary 创建Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Banner true "创建Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /banner/createBanner [post]
func (bannerApi *BannerApi) CreateBanner(c *gin.Context) {
	var banner autocode.Banner
	_ = c.ShouldBindJSON(&banner)
	if err := bannerService.CreateBanner(banner); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBanner 删除Banner
// @Tags Banner
// @Summary 删除Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Banner true "删除Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /banner/deleteBanner [delete]
func (bannerApi *BannerApi) DeleteBanner(c *gin.Context) {
	var banner autocode.Banner
	_ = c.ShouldBindJSON(&banner)
	if err := bannerService.DeleteBanner(banner); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBannerByIds 批量删除Banner
// @Tags Banner
// @Summary 批量删除Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /banner/deleteBannerByIds [delete]
func (bannerApi *BannerApi) DeleteBannerByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bannerService.DeleteBannerByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBanner 更新Banner
// @Tags Banner
// @Summary 更新Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Banner true "更新Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /banner/updateBanner [put]
func (bannerApi *BannerApi) UpdateBanner(c *gin.Context) {
	var banner autocode.Banner
	_ = c.ShouldBindJSON(&banner)
	if err := bannerService.UpdateBanner(banner); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBanner 用id查询Banner
// @Tags Banner
// @Summary 用id查询Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Banner true "用id查询Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /banner/findBanner [get]
func (bannerApi *BannerApi) FindBanner(c *gin.Context) {
	var banner autocode.Banner
	_ = c.ShouldBindQuery(&banner)
	if err, rebanner := bannerService.GetBanner(banner.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebanner": rebanner}, c)
	}
}

// GetBannerList 分页获取Banner列表
// @Tags Banner
// @Summary 分页获取Banner列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.BannerSearch true "分页获取Banner列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /banner/getBannerList [get]
func (bannerApi *BannerApi) GetBannerList(c *gin.Context) {
	var pageInfo autocodeReq.BannerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := bannerService.GetBannerInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
