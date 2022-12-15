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

type HelpCarpoolApi struct {
}

var helpCarpoolService = service.ServiceGroupApp.AutoCodeServiceGroup.HelpCarpoolService


// CreateHelpCarpool 创建HelpCarpool
// @Tags HelpCarpool
// @Summary 创建HelpCarpool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpCarpool true "创建HelpCarpool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpCarpool/createHelpCarpool [post]
func (helpCarpoolApi *HelpCarpoolApi) CreateHelpCarpool(c *gin.Context) {
	var helpCarpool autocode.HelpCarpool
	_ = c.ShouldBindJSON(&helpCarpool)
	if err := helpCarpoolService.CreateHelpCarpool(helpCarpool); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHelpCarpool 删除HelpCarpool
// @Tags HelpCarpool
// @Summary 删除HelpCarpool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpCarpool true "删除HelpCarpool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpCarpool/deleteHelpCarpool [delete]
func (helpCarpoolApi *HelpCarpoolApi) DeleteHelpCarpool(c *gin.Context) {
	var helpCarpool autocode.HelpCarpool
	_ = c.ShouldBindJSON(&helpCarpool)
	if err := helpCarpoolService.DeleteHelpCarpool(helpCarpool); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHelpCarpoolByIds 批量删除HelpCarpool
// @Tags HelpCarpool
// @Summary 批量删除HelpCarpool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpCarpool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /helpCarpool/deleteHelpCarpoolByIds [delete]
func (helpCarpoolApi *HelpCarpoolApi) DeleteHelpCarpoolByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := helpCarpoolService.DeleteHelpCarpoolByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHelpCarpool 更新HelpCarpool
// @Tags HelpCarpool
// @Summary 更新HelpCarpool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpCarpool true "更新HelpCarpool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpCarpool/updateHelpCarpool [put]
func (helpCarpoolApi *HelpCarpoolApi) UpdateHelpCarpool(c *gin.Context) {
	var helpCarpool autocode.HelpCarpool
	_ = c.ShouldBindJSON(&helpCarpool)
	if err := helpCarpoolService.UpdateHelpCarpool(helpCarpool); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHelpCarpool 用id查询HelpCarpool
// @Tags HelpCarpool
// @Summary 用id查询HelpCarpool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.HelpCarpool true "用id查询HelpCarpool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpCarpool/findHelpCarpool [get]
func (helpCarpoolApi *HelpCarpoolApi) FindHelpCarpool(c *gin.Context) {
	var helpCarpool autocode.HelpCarpool
	_ = c.ShouldBindQuery(&helpCarpool)
	if err, rehelpCarpool := helpCarpoolService.GetHelpCarpool(helpCarpool.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehelpCarpool": rehelpCarpool}, c)
	}
}

// GetHelpCarpoolList 分页获取HelpCarpool列表
// @Tags HelpCarpool
// @Summary 分页获取HelpCarpool列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.HelpCarpoolSearch true "分页获取HelpCarpool列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpCarpool/getHelpCarpoolList [get]
func (helpCarpoolApi *HelpCarpoolApi) GetHelpCarpoolList(c *gin.Context) {
	var pageInfo autocodeReq.HelpCarpoolSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := helpCarpoolService.GetHelpCarpoolInfoList(pageInfo); err != nil {
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
