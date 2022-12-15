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

type HelpWorkApi struct {
}

var helpWorkService = service.ServiceGroupApp.AutoCodeServiceGroup.HelpWorkService


// CreateHelpWork 创建HelpWork
// @Tags HelpWork
// @Summary 创建HelpWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpWork true "创建HelpWork"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpWork/createHelpWork [post]
func (helpWorkApi *HelpWorkApi) CreateHelpWork(c *gin.Context) {
	var helpWork autocode.HelpWork
	_ = c.ShouldBindJSON(&helpWork)
	if err := helpWorkService.CreateHelpWork(helpWork); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHelpWork 删除HelpWork
// @Tags HelpWork
// @Summary 删除HelpWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpWork true "删除HelpWork"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpWork/deleteHelpWork [delete]
func (helpWorkApi *HelpWorkApi) DeleteHelpWork(c *gin.Context) {
	var helpWork autocode.HelpWork
	_ = c.ShouldBindJSON(&helpWork)
	if err := helpWorkService.DeleteHelpWork(helpWork); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHelpWorkByIds 批量删除HelpWork
// @Tags HelpWork
// @Summary 批量删除HelpWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpWork"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /helpWork/deleteHelpWorkByIds [delete]
func (helpWorkApi *HelpWorkApi) DeleteHelpWorkByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := helpWorkService.DeleteHelpWorkByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHelpWork 更新HelpWork
// @Tags HelpWork
// @Summary 更新HelpWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HelpWork true "更新HelpWork"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpWork/updateHelpWork [put]
func (helpWorkApi *HelpWorkApi) UpdateHelpWork(c *gin.Context) {
	var helpWork autocode.HelpWork
	_ = c.ShouldBindJSON(&helpWork)
	if err := helpWorkService.UpdateHelpWork(helpWork); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHelpWork 用id查询HelpWork
// @Tags HelpWork
// @Summary 用id查询HelpWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.HelpWork true "用id查询HelpWork"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpWork/findHelpWork [get]
func (helpWorkApi *HelpWorkApi) FindHelpWork(c *gin.Context) {
	var helpWork autocode.HelpWork
	_ = c.ShouldBindQuery(&helpWork)
	if err, rehelpWork := helpWorkService.GetHelpWork(helpWork.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehelpWork": rehelpWork}, c)
	}
}

// GetHelpWorkList 分页获取HelpWork列表
// @Tags HelpWork
// @Summary 分页获取HelpWork列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.HelpWorkSearch true "分页获取HelpWork列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpWork/getHelpWorkList [get]
func (helpWorkApi *HelpWorkApi) GetHelpWorkList(c *gin.Context) {
	var pageInfo autocodeReq.HelpWorkSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := helpWorkService.GetHelpWorkInfoList(pageInfo); err != nil {
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
