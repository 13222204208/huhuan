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

type HelpMeApi struct {
}

var helpMeService = service.ServiceGroupApp.AutoCodeServiceGroup.HelpMeService

// CreateHelpMe 创建HelpMe
// @Tags HelpMe
// @Summary 创建HelpMe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpMe true "创建HelpMe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpMe/createHelpMe [post]
func (helpMeApi *HelpMeApi) CreateHelpMe(c *gin.Context) {
	var helpMe autocode.HelpMe
	_ = c.ShouldBindJSON(&helpMe)
	if err := helpMeService.CreateHelpMe(helpMe); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHelpMe 删除HelpMe
// @Tags HelpMe
// @Summary 删除HelpMe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpMe true "删除HelpMe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpMe/deleteHelpMe [delete]
func (helpMeApi *HelpMeApi) DeleteHelpMe(c *gin.Context) {
	var helpMe autocode.HelpMe
	_ = c.ShouldBindJSON(&helpMe)
	if err := helpMeService.DeleteHelpMe(helpMe); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHelpMeByIds 批量删除HelpMe
// @Tags HelpMe
// @Summary 批量删除HelpMe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpMe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /helpMe/deleteHelpMeByIds [delete]
func (helpMeApi *HelpMeApi) DeleteHelpMeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := helpMeService.DeleteHelpMeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHelpMe 更新HelpMe
// @Tags HelpMe
// @Summary 更新HelpMe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpMe true "更新HelpMe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpMe/updateHelpMe [put]
func (helpMeApi *HelpMeApi) UpdateHelpMe(c *gin.Context) {
	var helpMe autocode.HelpMe
	_ = c.ShouldBindJSON(&helpMe)
	if err := helpMeService.UpdateHelpMe(helpMe); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHelpMe 用id查询HelpMe
// @Tags HelpMe
// @Summary 用id查询HelpMe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.HelpMe true "用id查询HelpMe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpMe/findHelpMe [get]
func (helpMeApi *HelpMeApi) FindHelpMe(c *gin.Context) {
	var helpMe autocode.HelpMe
	_ = c.ShouldBindQuery(&helpMe)
	if err, rehelpMe := helpMeService.GetHelpMe(helpMe.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehelpMe": rehelpMe}, c)
	}
}

// GetHelpMeList 分页获取HelpMe列表
// @Tags HelpMe
// @Summary 分页获取HelpMe列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.HelpMeSearch true "分页获取HelpMe列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpMe/getHelpMeList [get]
func (helpMeApi *HelpMeApi) GetHelpMeList(c *gin.Context) {
	var pageInfo autocodeReq.HelpMeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := helpMeService.GetHelpMeInfoList(pageInfo); err != nil {
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

func (helpMeApi *HelpMeApi) GetHelpMeTitle(c *gin.Context) {
	if err, result := helpMeService.GetHelpMeTitle(); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithDetailed(gin.H{"HelpMeTitle": result}, "成功", c)
	}
}
