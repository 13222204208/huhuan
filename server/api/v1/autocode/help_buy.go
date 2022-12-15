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

type HelpBuyApi struct {
}

var helpBuyService = service.ServiceGroupApp.AutoCodeServiceGroup.HelpBuyService


// CreateHelpBuy 创建HelpBuy
// @Tags HelpBuy
// @Summary 创建HelpBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpBuy true "创建HelpBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpBuy/createHelpBuy [post]
func (helpBuyApi *HelpBuyApi) CreateHelpBuy(c *gin.Context) {
	var helpBuy autocode.HelpBuy
	_ = c.ShouldBindJSON(&helpBuy)
	if err := helpBuyService.CreateHelpBuy(helpBuy); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHelpBuy 删除HelpBuy
// @Tags HelpBuy
// @Summary 删除HelpBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpBuy true "删除HelpBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpBuy/deleteHelpBuy [delete]
func (helpBuyApi *HelpBuyApi) DeleteHelpBuy(c *gin.Context) {
	var helpBuy autocode.HelpBuy
	_ = c.ShouldBindJSON(&helpBuy)
	if err := helpBuyService.DeleteHelpBuy(helpBuy); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHelpBuyByIds 批量删除HelpBuy
// @Tags HelpBuy
// @Summary 批量删除HelpBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /helpBuy/deleteHelpBuyByIds [delete]
func (helpBuyApi *HelpBuyApi) DeleteHelpBuyByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := helpBuyService.DeleteHelpBuyByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHelpBuy 更新HelpBuy
// @Tags HelpBuy
// @Summary 更新HelpBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpBuy true "更新HelpBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpBuy/updateHelpBuy [put]
func (helpBuyApi *HelpBuyApi) UpdateHelpBuy(c *gin.Context) {
	var helpBuy autocode.HelpBuy
	_ = c.ShouldBindJSON(&helpBuy)
	if err := helpBuyService.UpdateHelpBuy(helpBuy); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHelpBuy 用id查询HelpBuy
// @Tags HelpBuy
// @Summary 用id查询HelpBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.HelpBuy true "用id查询HelpBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpBuy/findHelpBuy [get]
func (helpBuyApi *HelpBuyApi) FindHelpBuy(c *gin.Context) {
	var helpBuy autocode.HelpBuy
	_ = c.ShouldBindQuery(&helpBuy)
	if err, rehelpBuy := helpBuyService.GetHelpBuy(helpBuy.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehelpBuy": rehelpBuy}, c)
	}
}

// GetHelpBuyList 分页获取HelpBuy列表
// @Tags HelpBuy
// @Summary 分页获取HelpBuy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.HelpBuySearch true "分页获取HelpBuy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpBuy/getHelpBuyList [get]
func (helpBuyApi *HelpBuyApi) GetHelpBuyList(c *gin.Context) {
	var pageInfo autocodeReq.HelpBuySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := helpBuyService.GetHelpBuyInfoList(pageInfo); err != nil {
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
