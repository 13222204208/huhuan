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

type HelpSecondApi struct {
}

var helpSecondService = service.ServiceGroupApp.AutoCodeServiceGroup.HelpSecondService


// CreateHelpSecond 创建HelpSecond
// @Tags HelpSecond
// @Summary 创建HelpSecond
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpSecond true "创建HelpSecond"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpSecond/createHelpSecond [post]
func (helpSecondApi *HelpSecondApi) CreateHelpSecond(c *gin.Context) {
	var helpSecond autocode.HelpSecond
	_ = c.ShouldBindJSON(&helpSecond)
	if err := helpSecondService.CreateHelpSecond(helpSecond); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHelpSecond 删除HelpSecond
// @Tags HelpSecond
// @Summary 删除HelpSecond
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpSecond true "删除HelpSecond"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpSecond/deleteHelpSecond [delete]
func (helpSecondApi *HelpSecondApi) DeleteHelpSecond(c *gin.Context) {
	var helpSecond autocode.HelpSecond
	_ = c.ShouldBindJSON(&helpSecond)
	if err := helpSecondService.DeleteHelpSecond(helpSecond); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHelpSecondByIds 批量删除HelpSecond
// @Tags HelpSecond
// @Summary 批量删除HelpSecond
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpSecond"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /helpSecond/deleteHelpSecondByIds [delete]
func (helpSecondApi *HelpSecondApi) DeleteHelpSecondByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := helpSecondService.DeleteHelpSecondByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHelpSecond 更新HelpSecond
// @Tags HelpSecond
// @Summary 更新HelpSecond
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpSecond true "更新HelpSecond"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpSecond/updateHelpSecond [put]
func (helpSecondApi *HelpSecondApi) UpdateHelpSecond(c *gin.Context) {
	var helpSecond autocode.HelpSecond
	_ = c.ShouldBindJSON(&helpSecond)
	if err := helpSecondService.UpdateHelpSecond(helpSecond); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHelpSecond 用id查询HelpSecond
// @Tags HelpSecond
// @Summary 用id查询HelpSecond
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.HelpSecond true "用id查询HelpSecond"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpSecond/findHelpSecond [get]
func (helpSecondApi *HelpSecondApi) FindHelpSecond(c *gin.Context) {
	var helpSecond autocode.HelpSecond
	_ = c.ShouldBindQuery(&helpSecond)
	if err, rehelpSecond := helpSecondService.GetHelpSecond(helpSecond.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehelpSecond": rehelpSecond}, c)
	}
}

// GetHelpSecondList 分页获取HelpSecond列表
// @Tags HelpSecond
// @Summary 分页获取HelpSecond列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.HelpSecondSearch true "分页获取HelpSecond列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpSecond/getHelpSecondList [get]
func (helpSecondApi *HelpSecondApi) GetHelpSecondList(c *gin.Context) {
	var pageInfo autocodeReq.HelpSecondSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := helpSecondService.GetHelpSecondInfoList(pageInfo); err != nil {
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
