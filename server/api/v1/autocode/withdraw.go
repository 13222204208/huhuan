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

type WithdrawApi struct {
}

var withdrawService = service.ServiceGroupApp.AutoCodeServiceGroup.WithdrawService


// CreateWithdraw 创建Withdraw
// @Tags Withdraw
// @Summary 创建Withdraw
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Withdraw true "创建Withdraw"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /withdraw/createWithdraw [post]
func (withdrawApi *WithdrawApi) CreateWithdraw(c *gin.Context) {
	var withdraw autocode.Withdraw
	_ = c.ShouldBindJSON(&withdraw)
	if err := withdrawService.CreateWithdraw(withdraw); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWithdraw 删除Withdraw
// @Tags Withdraw
// @Summary 删除Withdraw
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Withdraw true "删除Withdraw"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /withdraw/deleteWithdraw [delete]
func (withdrawApi *WithdrawApi) DeleteWithdraw(c *gin.Context) {
	var withdraw autocode.Withdraw
	_ = c.ShouldBindJSON(&withdraw)
	if err := withdrawService.DeleteWithdraw(withdraw); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWithdrawByIds 批量删除Withdraw
// @Tags Withdraw
// @Summary 批量删除Withdraw
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Withdraw"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /withdraw/deleteWithdrawByIds [delete]
func (withdrawApi *WithdrawApi) DeleteWithdrawByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := withdrawService.DeleteWithdrawByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWithdraw 更新Withdraw
// @Tags Withdraw
// @Summary 更新Withdraw
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Withdraw true "更新Withdraw"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /withdraw/updateWithdraw [put]
func (withdrawApi *WithdrawApi) UpdateWithdraw(c *gin.Context) {
	var withdraw autocode.Withdraw
	_ = c.ShouldBindJSON(&withdraw)
	if err := withdrawService.UpdateWithdraw(withdraw); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWithdraw 用id查询Withdraw
// @Tags Withdraw
// @Summary 用id查询Withdraw
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Withdraw true "用id查询Withdraw"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /withdraw/findWithdraw [get]
func (withdrawApi *WithdrawApi) FindWithdraw(c *gin.Context) {
	var withdraw autocode.Withdraw
	_ = c.ShouldBindQuery(&withdraw)
	if err, rewithdraw := withdrawService.GetWithdraw(withdraw.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewithdraw": rewithdraw}, c)
	}
}

// GetWithdrawList 分页获取Withdraw列表
// @Tags Withdraw
// @Summary 分页获取Withdraw列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.WithdrawSearch true "分页获取Withdraw列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /withdraw/getWithdrawList [get]
func (withdrawApi *WithdrawApi) GetWithdrawList(c *gin.Context) {
	var pageInfo autocodeReq.WithdrawSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := withdrawService.GetWithdrawInfoList(pageInfo); err != nil {
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
