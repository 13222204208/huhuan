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

type HelpGrouponApi struct {
}

var helpGrouponService = service.ServiceGroupApp.AutoCodeServiceGroup.HelpGrouponService


// CreateHelpGroupon 创建HelpGroupon
// @Tags HelpGroupon
// @Summary 创建HelpGroupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpGroupon true "创建HelpGroupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpGroupon/createHelpGroupon [post]
func (helpGrouponApi *HelpGrouponApi) CreateHelpGroupon(c *gin.Context) {
	var helpGroupon autocode.HelpGroupon
	_ = c.ShouldBindJSON(&helpGroupon)
	if err := helpGrouponService.CreateHelpGroupon(helpGroupon); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHelpGroupon 删除HelpGroupon
// @Tags HelpGroupon
// @Summary 删除HelpGroupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpGroupon true "删除HelpGroupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpGroupon/deleteHelpGroupon [delete]
func (helpGrouponApi *HelpGrouponApi) DeleteHelpGroupon(c *gin.Context) {
	var helpGroupon autocode.HelpGroupon
	_ = c.ShouldBindJSON(&helpGroupon)
	if err := helpGrouponService.DeleteHelpGroupon(helpGroupon); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHelpGrouponByIds 批量删除HelpGroupon
// @Tags HelpGroupon
// @Summary 批量删除HelpGroupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpGroupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /helpGroupon/deleteHelpGrouponByIds [delete]
func (helpGrouponApi *HelpGrouponApi) DeleteHelpGrouponByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := helpGrouponService.DeleteHelpGrouponByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHelpGroupon 更新HelpGroupon
// @Tags HelpGroupon
// @Summary 更新HelpGroupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpGroupon true "更新HelpGroupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpGroupon/updateHelpGroupon [put]
func (helpGrouponApi *HelpGrouponApi) UpdateHelpGroupon(c *gin.Context) {
	var helpGroupon autocode.HelpGroupon
	_ = c.ShouldBindJSON(&helpGroupon)
	if err := helpGrouponService.UpdateHelpGroupon(helpGroupon); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHelpGroupon 用id查询HelpGroupon
// @Tags HelpGroupon
// @Summary 用id查询HelpGroupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.HelpGroupon true "用id查询HelpGroupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpGroupon/findHelpGroupon [get]
func (helpGrouponApi *HelpGrouponApi) FindHelpGroupon(c *gin.Context) {
	var helpGroupon autocode.HelpGroupon
	_ = c.ShouldBindQuery(&helpGroupon)
	if err, rehelpGroupon := helpGrouponService.GetHelpGroupon(helpGroupon.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehelpGroupon": rehelpGroupon}, c)
	}
}

// GetHelpGrouponList 分页获取HelpGroupon列表
// @Tags HelpGroupon
// @Summary 分页获取HelpGroupon列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.HelpGrouponSearch true "分页获取HelpGroupon列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpGroupon/getHelpGrouponList [get]
func (helpGrouponApi *HelpGrouponApi) GetHelpGrouponList(c *gin.Context) {
	var pageInfo autocodeReq.HelpGrouponSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := helpGrouponService.GetHelpGrouponInfoList(pageInfo); err != nil {
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
