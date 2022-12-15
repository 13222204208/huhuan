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

type HelpIncidentallyApi struct {
}

var helpIncidentallyService = service.ServiceGroupApp.AutoCodeServiceGroup.HelpIncidentallyService


// CreateHelpIncidentally 创建HelpIncidentally
// @Tags HelpIncidentally
// @Summary 创建HelpIncidentally
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpIncidentally true "创建HelpIncidentally"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpIncidentally/createHelpIncidentally [post]
func (helpIncidentallyApi *HelpIncidentallyApi) CreateHelpIncidentally(c *gin.Context) {
	var helpIncidentally autocode.HelpIncidentally
	_ = c.ShouldBindJSON(&helpIncidentally)
	if err := helpIncidentallyService.CreateHelpIncidentally(helpIncidentally); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHelpIncidentally 删除HelpIncidentally
// @Tags HelpIncidentally
// @Summary 删除HelpIncidentally
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpIncidentally true "删除HelpIncidentally"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpIncidentally/deleteHelpIncidentally [delete]
func (helpIncidentallyApi *HelpIncidentallyApi) DeleteHelpIncidentally(c *gin.Context) {
	var helpIncidentally autocode.HelpIncidentally
	_ = c.ShouldBindJSON(&helpIncidentally)
	if err := helpIncidentallyService.DeleteHelpIncidentally(helpIncidentally); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHelpIncidentallyByIds 批量删除HelpIncidentally
// @Tags HelpIncidentally
// @Summary 批量删除HelpIncidentally
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpIncidentally"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /helpIncidentally/deleteHelpIncidentallyByIds [delete]
func (helpIncidentallyApi *HelpIncidentallyApi) DeleteHelpIncidentallyByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := helpIncidentallyService.DeleteHelpIncidentallyByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHelpIncidentally 更新HelpIncidentally
// @Tags HelpIncidentally
// @Summary 更新HelpIncidentally
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpIncidentally true "更新HelpIncidentally"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpIncidentally/updateHelpIncidentally [put]
func (helpIncidentallyApi *HelpIncidentallyApi) UpdateHelpIncidentally(c *gin.Context) {
	var helpIncidentally autocode.HelpIncidentally
	_ = c.ShouldBindJSON(&helpIncidentally)
	if err := helpIncidentallyService.UpdateHelpIncidentally(helpIncidentally); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHelpIncidentally 用id查询HelpIncidentally
// @Tags HelpIncidentally
// @Summary 用id查询HelpIncidentally
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.HelpIncidentally true "用id查询HelpIncidentally"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpIncidentally/findHelpIncidentally [get]
func (helpIncidentallyApi *HelpIncidentallyApi) FindHelpIncidentally(c *gin.Context) {
	var helpIncidentally autocode.HelpIncidentally
	_ = c.ShouldBindQuery(&helpIncidentally)
	if err, rehelpIncidentally := helpIncidentallyService.GetHelpIncidentally(helpIncidentally.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehelpIncidentally": rehelpIncidentally}, c)
	}
}

// GetHelpIncidentallyList 分页获取HelpIncidentally列表
// @Tags HelpIncidentally
// @Summary 分页获取HelpIncidentally列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.HelpIncidentallySearch true "分页获取HelpIncidentally列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpIncidentally/getHelpIncidentallyList [get]
func (helpIncidentallyApi *HelpIncidentallyApi) GetHelpIncidentallyList(c *gin.Context) {
	var pageInfo autocodeReq.HelpIncidentallySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := helpIncidentallyService.GetHelpIncidentallyInfoList(pageInfo); err != nil {
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
