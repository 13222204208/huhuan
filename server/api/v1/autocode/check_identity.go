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

type CheckIdentityApi struct {
}

var checkIdentityService = service.ServiceGroupApp.AutoCodeServiceGroup.CheckIdentityService


// CreateCheckIdentity 创建CheckIdentity
// @Tags CheckIdentity
// @Summary 创建CheckIdentity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CheckIdentity true "创建CheckIdentity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkIdentity/createCheckIdentity [post]
func (checkIdentityApi *CheckIdentityApi) CreateCheckIdentity(c *gin.Context) {
	var checkIdentity autocode.CheckIdentity
	_ = c.ShouldBindJSON(&checkIdentity)
	if err := checkIdentityService.CreateCheckIdentity(checkIdentity); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCheckIdentity 删除CheckIdentity
// @Tags CheckIdentity
// @Summary 删除CheckIdentity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CheckIdentity true "删除CheckIdentity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /checkIdentity/deleteCheckIdentity [delete]
func (checkIdentityApi *CheckIdentityApi) DeleteCheckIdentity(c *gin.Context) {
	var checkIdentity autocode.CheckIdentity
	_ = c.ShouldBindJSON(&checkIdentity)
	if err := checkIdentityService.DeleteCheckIdentity(checkIdentity); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCheckIdentityByIds 批量删除CheckIdentity
// @Tags CheckIdentity
// @Summary 批量删除CheckIdentity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CheckIdentity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /checkIdentity/deleteCheckIdentityByIds [delete]
func (checkIdentityApi *CheckIdentityApi) DeleteCheckIdentityByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := checkIdentityService.DeleteCheckIdentityByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCheckIdentity 更新CheckIdentity
// @Tags CheckIdentity
// @Summary 更新CheckIdentity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CheckIdentity true "更新CheckIdentity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /checkIdentity/updateCheckIdentity [put]
func (checkIdentityApi *CheckIdentityApi) UpdateCheckIdentity(c *gin.Context) {
	var checkIdentity autocode.CheckIdentity
	_ = c.ShouldBindJSON(&checkIdentity)
	if err := checkIdentityService.UpdateCheckIdentity(checkIdentity); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCheckIdentity 用id查询CheckIdentity
// @Tags CheckIdentity
// @Summary 用id查询CheckIdentity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CheckIdentity true "用id查询CheckIdentity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /checkIdentity/findCheckIdentity [get]
func (checkIdentityApi *CheckIdentityApi) FindCheckIdentity(c *gin.Context) {
	var checkIdentity autocode.CheckIdentity
	_ = c.ShouldBindQuery(&checkIdentity)
	if err, recheckIdentity := checkIdentityService.GetCheckIdentity(checkIdentity.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recheckIdentity": recheckIdentity}, c)
	}
}

// GetCheckIdentityList 分页获取CheckIdentity列表
// @Tags CheckIdentity
// @Summary 分页获取CheckIdentity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CheckIdentitySearch true "分页获取CheckIdentity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkIdentity/getCheckIdentityList [get]
func (checkIdentityApi *CheckIdentityApi) GetCheckIdentityList(c *gin.Context) {
	var pageInfo autocodeReq.CheckIdentitySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := checkIdentityService.GetCheckIdentityInfoList(pageInfo); err != nil {
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
