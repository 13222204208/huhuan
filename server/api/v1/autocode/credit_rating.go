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

type CreditRatingApi struct {
}

var creditRatingService = service.ServiceGroupApp.AutoCodeServiceGroup.CreditRatingService


// CreateCreditRating 创建CreditRating
// @Tags CreditRating
// @Summary 创建CreditRating
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CreditRating true "创建CreditRating"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /creditRating/createCreditRating [post]
func (creditRatingApi *CreditRatingApi) CreateCreditRating(c *gin.Context) {
	var creditRating autocode.CreditRating
	_ = c.ShouldBindJSON(&creditRating)
	if err := creditRatingService.CreateCreditRating(creditRating); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCreditRating 删除CreditRating
// @Tags CreditRating
// @Summary 删除CreditRating
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CreditRating true "删除CreditRating"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /creditRating/deleteCreditRating [delete]
func (creditRatingApi *CreditRatingApi) DeleteCreditRating(c *gin.Context) {
	var creditRating autocode.CreditRating
	_ = c.ShouldBindJSON(&creditRating)
	if err := creditRatingService.DeleteCreditRating(creditRating); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCreditRatingByIds 批量删除CreditRating
// @Tags CreditRating
// @Summary 批量删除CreditRating
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CreditRating"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /creditRating/deleteCreditRatingByIds [delete]
func (creditRatingApi *CreditRatingApi) DeleteCreditRatingByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := creditRatingService.DeleteCreditRatingByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCreditRating 更新CreditRating
// @Tags CreditRating
// @Summary 更新CreditRating
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CreditRating true "更新CreditRating"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /creditRating/updateCreditRating [put]
func (creditRatingApi *CreditRatingApi) UpdateCreditRating(c *gin.Context) {
	var creditRating autocode.CreditRating
	_ = c.ShouldBindJSON(&creditRating)
	if err := creditRatingService.UpdateCreditRating(creditRating); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCreditRating 用id查询CreditRating
// @Tags CreditRating
// @Summary 用id查询CreditRating
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CreditRating true "用id查询CreditRating"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /creditRating/findCreditRating [get]
func (creditRatingApi *CreditRatingApi) FindCreditRating(c *gin.Context) {
	var creditRating autocode.CreditRating
	_ = c.ShouldBindQuery(&creditRating)
	if err, recreditRating := creditRatingService.GetCreditRating(creditRating.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recreditRating": recreditRating}, c)
	}
}

// GetCreditRatingList 分页获取CreditRating列表
// @Tags CreditRating
// @Summary 分页获取CreditRating列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CreditRatingSearch true "分页获取CreditRating列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /creditRating/getCreditRatingList [get]
func (creditRatingApi *CreditRatingApi) GetCreditRatingList(c *gin.Context) {
	var pageInfo autocodeReq.CreditRatingSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := creditRatingService.GetCreditRatingInfoList(pageInfo); err != nil {
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
