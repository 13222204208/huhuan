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

type CheckDriverApi struct {
}

var checkDriverService = service.ServiceGroupApp.AutoCodeServiceGroup.CheckDriverService


// CreateCheckDriver 创建CheckDriver
// @Tags CheckDriver
// @Summary 创建CheckDriver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CheckDriver true "创建CheckDriver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkDriver/createCheckDriver [post]
func (checkDriverApi *CheckDriverApi) CreateCheckDriver(c *gin.Context) {
	var checkDriver autocode.CheckDriver
	_ = c.ShouldBindJSON(&checkDriver)
	if err := checkDriverService.CreateCheckDriver(checkDriver); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCheckDriver 删除CheckDriver
// @Tags CheckDriver
// @Summary 删除CheckDriver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CheckDriver true "删除CheckDriver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /checkDriver/deleteCheckDriver [delete]
func (checkDriverApi *CheckDriverApi) DeleteCheckDriver(c *gin.Context) {
	var checkDriver autocode.CheckDriver
	_ = c.ShouldBindJSON(&checkDriver)
	if err := checkDriverService.DeleteCheckDriver(checkDriver); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCheckDriverByIds 批量删除CheckDriver
// @Tags CheckDriver
// @Summary 批量删除CheckDriver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CheckDriver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /checkDriver/deleteCheckDriverByIds [delete]
func (checkDriverApi *CheckDriverApi) DeleteCheckDriverByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := checkDriverService.DeleteCheckDriverByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCheckDriver 更新CheckDriver
// @Tags CheckDriver
// @Summary 更新CheckDriver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CheckDriver true "更新CheckDriver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /checkDriver/updateCheckDriver [put]
func (checkDriverApi *CheckDriverApi) UpdateCheckDriver(c *gin.Context) {
	var checkDriver autocode.CheckDriver
	_ = c.ShouldBindJSON(&checkDriver)
	if err := checkDriverService.UpdateCheckDriver(checkDriver); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCheckDriver 用id查询CheckDriver
// @Tags CheckDriver
// @Summary 用id查询CheckDriver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CheckDriver true "用id查询CheckDriver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /checkDriver/findCheckDriver [get]
func (checkDriverApi *CheckDriverApi) FindCheckDriver(c *gin.Context) {
	var checkDriver autocode.CheckDriver
	_ = c.ShouldBindQuery(&checkDriver)
	if err, recheckDriver := checkDriverService.GetCheckDriver(checkDriver.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recheckDriver": recheckDriver}, c)
	}
}

// GetCheckDriverList 分页获取CheckDriver列表
// @Tags CheckDriver
// @Summary 分页获取CheckDriver列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CheckDriverSearch true "分页获取CheckDriver列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkDriver/getCheckDriverList [get]
func (checkDriverApi *CheckDriverApi) GetCheckDriverList(c *gin.Context) {
	var pageInfo autocodeReq.CheckDriverSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := checkDriverService.GetCheckDriverInfoList(pageInfo); err != nil {
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
