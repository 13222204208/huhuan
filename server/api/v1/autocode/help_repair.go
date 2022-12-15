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

type HelpRepairApi struct {
}

var helpRepairService = service.ServiceGroupApp.AutoCodeServiceGroup.HelpRepairService


// CreateHelpRepair 创建HelpRepair
// @Tags HelpRepair
// @Summary 创建HelpRepair
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpRepair true "创建HelpRepair"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpRepair/createHelpRepair [post]
func (helpRepairApi *HelpRepairApi) CreateHelpRepair(c *gin.Context) {
	var helpRepair autocode.HelpRepair
	_ = c.ShouldBindJSON(&helpRepair)
	if err := helpRepairService.CreateHelpRepair(helpRepair); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHelpRepair 删除HelpRepair
// @Tags HelpRepair
// @Summary 删除HelpRepair
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpRepair true "删除HelpRepair"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpRepair/deleteHelpRepair [delete]
func (helpRepairApi *HelpRepairApi) DeleteHelpRepair(c *gin.Context) {
	var helpRepair autocode.HelpRepair
	_ = c.ShouldBindJSON(&helpRepair)
	if err := helpRepairService.DeleteHelpRepair(helpRepair); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHelpRepairByIds 批量删除HelpRepair
// @Tags HelpRepair
// @Summary 批量删除HelpRepair
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpRepair"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /helpRepair/deleteHelpRepairByIds [delete]
func (helpRepairApi *HelpRepairApi) DeleteHelpRepairByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := helpRepairService.DeleteHelpRepairByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHelpRepair 更新HelpRepair
// @Tags HelpRepair
// @Summary 更新HelpRepair
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpRepair true "更新HelpRepair"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpRepair/updateHelpRepair [put]
func (helpRepairApi *HelpRepairApi) UpdateHelpRepair(c *gin.Context) {
	var helpRepair autocode.HelpRepair
	_ = c.ShouldBindJSON(&helpRepair)
	if err := helpRepairService.UpdateHelpRepair(helpRepair); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHelpRepair 用id查询HelpRepair
// @Tags HelpRepair
// @Summary 用id查询HelpRepair
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.HelpRepair true "用id查询HelpRepair"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpRepair/findHelpRepair [get]
func (helpRepairApi *HelpRepairApi) FindHelpRepair(c *gin.Context) {
	var helpRepair autocode.HelpRepair
	_ = c.ShouldBindQuery(&helpRepair)
	if err, rehelpRepair := helpRepairService.GetHelpRepair(helpRepair.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehelpRepair": rehelpRepair}, c)
	}
}

// GetHelpRepairList 分页获取HelpRepair列表
// @Tags HelpRepair
// @Summary 分页获取HelpRepair列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.HelpRepairSearch true "分页获取HelpRepair列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpRepair/getHelpRepairList [get]
func (helpRepairApi *HelpRepairApi) GetHelpRepairList(c *gin.Context) {
	var pageInfo autocodeReq.HelpRepairSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := helpRepairService.GetHelpRepairInfoList(pageInfo); err != nil {
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
