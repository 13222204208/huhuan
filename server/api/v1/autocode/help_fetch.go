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

type HelpFetchApi struct {
}

var helpFetchService = service.ServiceGroupApp.AutoCodeServiceGroup.HelpFetchService


// CreateHelpFetch 创建HelpFetch
// @Tags HelpFetch
// @Summary 创建HelpFetch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpFetch true "创建HelpFetch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpFetch/createHelpFetch [post]
func (helpFetchApi *HelpFetchApi) CreateHelpFetch(c *gin.Context) {
	var helpFetch autocode.HelpFetch
	_ = c.ShouldBindJSON(&helpFetch)
	if err := helpFetchService.CreateHelpFetch(helpFetch); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHelpFetch 删除HelpFetch
// @Tags HelpFetch
// @Summary 删除HelpFetch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpFetch true "删除HelpFetch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpFetch/deleteHelpFetch [delete]
func (helpFetchApi *HelpFetchApi) DeleteHelpFetch(c *gin.Context) {
	var helpFetch autocode.HelpFetch
	_ = c.ShouldBindJSON(&helpFetch)
	if err := helpFetchService.DeleteHelpFetch(helpFetch); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHelpFetchByIds 批量删除HelpFetch
// @Tags HelpFetch
// @Summary 批量删除HelpFetch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpFetch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /helpFetch/deleteHelpFetchByIds [delete]
func (helpFetchApi *HelpFetchApi) DeleteHelpFetchByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := helpFetchService.DeleteHelpFetchByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHelpFetch 更新HelpFetch
// @Tags HelpFetch
// @Summary 更新HelpFetch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpFetch true "更新HelpFetch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpFetch/updateHelpFetch [put]
func (helpFetchApi *HelpFetchApi) UpdateHelpFetch(c *gin.Context) {
	var helpFetch autocode.HelpFetch
	_ = c.ShouldBindJSON(&helpFetch)
	if err := helpFetchService.UpdateHelpFetch(helpFetch); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHelpFetch 用id查询HelpFetch
// @Tags HelpFetch
// @Summary 用id查询HelpFetch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.HelpFetch true "用id查询HelpFetch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpFetch/findHelpFetch [get]
func (helpFetchApi *HelpFetchApi) FindHelpFetch(c *gin.Context) {
	var helpFetch autocode.HelpFetch
	_ = c.ShouldBindQuery(&helpFetch)
	if err, rehelpFetch := helpFetchService.GetHelpFetch(helpFetch.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehelpFetch": rehelpFetch}, c)
	}
}

// GetHelpFetchList 分页获取HelpFetch列表
// @Tags HelpFetch
// @Summary 分页获取HelpFetch列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.HelpFetchSearch true "分页获取HelpFetch列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpFetch/getHelpFetchList [get]
func (helpFetchApi *HelpFetchApi) GetHelpFetchList(c *gin.Context) {
	var pageInfo autocodeReq.HelpFetchSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := helpFetchService.GetHelpFetchInfoList(pageInfo); err != nil {
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
