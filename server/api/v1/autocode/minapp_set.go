package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MinappSetApi struct {
}

var minappSetService = service.ServiceGroupApp.AutoCodeServiceGroup.MinappSetService

// CreateMinappSet 创建MinappSet
// @Tags MinappSet
// @Summary 创建MinappSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.MinappSet true "创建MinappSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /minappSet/createMinappSet [post]
func (minappSetApi *MinappSetApi) CreateMinappSet(c *gin.Context) {
	var minappSet autocode.MinappSet
	_ = c.ShouldBindJSON(&minappSet)

	if err := minappSetService.CreateMinappSet(minappSet); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMinappSet 删除MinappSet
// @Tags MinappSet
// @Summary 删除MinappSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.MinappSet true "删除MinappSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /minappSet/deleteMinappSet [delete]
func (minappSetApi *MinappSetApi) DeleteMinappSet(c *gin.Context) {
	var minappSet autocode.MinappSet
	_ = c.ShouldBindJSON(&minappSet)
	if err := minappSetService.DeleteMinappSet(minappSet); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMinappSetByIds 批量删除MinappSet
// @Tags MinappSet
// @Summary 批量删除MinappSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MinappSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /minappSet/deleteMinappSetByIds [delete]
func (minappSetApi *MinappSetApi) DeleteMinappSetByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := minappSetService.DeleteMinappSetByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMinappSet 更新MinappSet
// @Tags MinappSet
// @Summary 更新MinappSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.MinappSet true "更新MinappSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /minappSet/updateMinappSet [put]
func (minappSetApi *MinappSetApi) UpdateMinappSet(c *gin.Context) {
	var minappSet autocode.MinappSet
	_ = c.ShouldBindJSON(&minappSet)

	if err := minappSetService.UpdateMinappSet(minappSet); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMinappSet 用id查询MinappSet
// @Tags MinappSet
// @Summary 用id查询MinappSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.MinappSet true "用id查询MinappSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /minappSet/findMinappSet [get]
func (minappSetApi *MinappSetApi) FindMinappSet(c *gin.Context) {
	var minappSet autocode.MinappSet
	_ = c.ShouldBindQuery(&minappSet)
	if err, reminappSet := minappSetService.GetMinappSet(minappSet.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reminappSet": reminappSet}, c)
	}
}

// GetMinappSetList 分页获取MinappSet列表
// @Tags MinappSet
// @Summary 分页获取MinappSet列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.MinappSetSearch true "分页获取MinappSet列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /minappSet/getMinappSetList [get]
func (minappSetApi *MinappSetApi) GetMinappSetList(c *gin.Context) {
	var pageInfo autocodeReq.MinappSetSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := minappSetService.GetMinappSetInfoList(pageInfo); err != nil {
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
