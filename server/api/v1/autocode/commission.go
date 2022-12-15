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

type CommissionApi struct {
}

var commissionService = service.ServiceGroupApp.AutoCodeServiceGroup.CommissionService


// CreateCommission 创建Commission
// @Tags Commission
// @Summary 创建Commission
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Commission true "创建Commission"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /commission/createCommission [post]
func (commissionApi *CommissionApi) CreateCommission(c *gin.Context) {
	var commission autocode.Commission
	_ = c.ShouldBindJSON(&commission)
	if err := commissionService.CreateCommission(commission); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCommission 删除Commission
// @Tags Commission
// @Summary 删除Commission
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Commission true "删除Commission"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /commission/deleteCommission [delete]
func (commissionApi *CommissionApi) DeleteCommission(c *gin.Context) {
	var commission autocode.Commission
	_ = c.ShouldBindJSON(&commission)
	if err := commissionService.DeleteCommission(commission); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCommissionByIds 批量删除Commission
// @Tags Commission
// @Summary 批量删除Commission
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Commission"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /commission/deleteCommissionByIds [delete]
func (commissionApi *CommissionApi) DeleteCommissionByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := commissionService.DeleteCommissionByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCommission 更新Commission
// @Tags Commission
// @Summary 更新Commission
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Commission true "更新Commission"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /commission/updateCommission [put]
func (commissionApi *CommissionApi) UpdateCommission(c *gin.Context) {
	var commission autocode.Commission
	_ = c.ShouldBindJSON(&commission)
	if err := commissionService.UpdateCommission(commission); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCommission 用id查询Commission
// @Tags Commission
// @Summary 用id查询Commission
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Commission true "用id查询Commission"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /commission/findCommission [get]
func (commissionApi *CommissionApi) FindCommission(c *gin.Context) {
	var commission autocode.Commission
	_ = c.ShouldBindQuery(&commission)
	if err, recommission := commissionService.GetCommission(commission.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recommission": recommission}, c)
	}
}

// GetCommissionList 分页获取Commission列表
// @Tags Commission
// @Summary 分页获取Commission列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CommissionSearch true "分页获取Commission列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /commission/getCommissionList [get]
func (commissionApi *CommissionApi) GetCommissionList(c *gin.Context) {
	var pageInfo autocodeReq.CommissionSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := commissionService.GetCommissionInfoList(pageInfo); err != nil {
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
