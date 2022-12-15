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

type CheckMerchantApi struct {
}

var checkMerchantService = service.ServiceGroupApp.AutoCodeServiceGroup.CheckMerchantService


// CreateCheckMerchant 创建CheckMerchant
// @Tags CheckMerchant
// @Summary 创建CheckMerchant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CheckMerchant true "创建CheckMerchant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkMerchant/createCheckMerchant [post]
func (checkMerchantApi *CheckMerchantApi) CreateCheckMerchant(c *gin.Context) {
	var checkMerchant autocode.CheckMerchant
	_ = c.ShouldBindJSON(&checkMerchant)
	if err := checkMerchantService.CreateCheckMerchant(checkMerchant); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCheckMerchant 删除CheckMerchant
// @Tags CheckMerchant
// @Summary 删除CheckMerchant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CheckMerchant true "删除CheckMerchant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /checkMerchant/deleteCheckMerchant [delete]
func (checkMerchantApi *CheckMerchantApi) DeleteCheckMerchant(c *gin.Context) {
	var checkMerchant autocode.CheckMerchant
	_ = c.ShouldBindJSON(&checkMerchant)
	if err := checkMerchantService.DeleteCheckMerchant(checkMerchant); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCheckMerchantByIds 批量删除CheckMerchant
// @Tags CheckMerchant
// @Summary 批量删除CheckMerchant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CheckMerchant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /checkMerchant/deleteCheckMerchantByIds [delete]
func (checkMerchantApi *CheckMerchantApi) DeleteCheckMerchantByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := checkMerchantService.DeleteCheckMerchantByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCheckMerchant 更新CheckMerchant
// @Tags CheckMerchant
// @Summary 更新CheckMerchant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CheckMerchant true "更新CheckMerchant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /checkMerchant/updateCheckMerchant [put]
func (checkMerchantApi *CheckMerchantApi) UpdateCheckMerchant(c *gin.Context) {
	var checkMerchant autocode.CheckMerchant
	_ = c.ShouldBindJSON(&checkMerchant)
	if err := checkMerchantService.UpdateCheckMerchant(checkMerchant); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCheckMerchant 用id查询CheckMerchant
// @Tags CheckMerchant
// @Summary 用id查询CheckMerchant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CheckMerchant true "用id查询CheckMerchant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /checkMerchant/findCheckMerchant [get]
func (checkMerchantApi *CheckMerchantApi) FindCheckMerchant(c *gin.Context) {
	var checkMerchant autocode.CheckMerchant
	_ = c.ShouldBindQuery(&checkMerchant)
	if err, recheckMerchant := checkMerchantService.GetCheckMerchant(checkMerchant.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recheckMerchant": recheckMerchant}, c)
	}
}

// GetCheckMerchantList 分页获取CheckMerchant列表
// @Tags CheckMerchant
// @Summary 分页获取CheckMerchant列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CheckMerchantSearch true "分页获取CheckMerchant列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkMerchant/getCheckMerchantList [get]
func (checkMerchantApi *CheckMerchantApi) GetCheckMerchantList(c *gin.Context) {
	var pageInfo autocodeReq.CheckMerchantSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := checkMerchantService.GetCheckMerchantInfoList(pageInfo); err != nil {
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
