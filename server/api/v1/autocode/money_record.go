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

type MoneyRecordApi struct {
}

var moneyRecordService = service.ServiceGroupApp.AutoCodeServiceGroup.MoneyRecordService


// CreateMoneyRecord 创建MoneyRecord
// @Tags MoneyRecord
// @Summary 创建MoneyRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.MoneyRecord true "创建MoneyRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneyRecord/createMoneyRecord [post]
func (moneyRecordApi *MoneyRecordApi) CreateMoneyRecord(c *gin.Context) {
	var moneyRecord autocode.MoneyRecord
	_ = c.ShouldBindJSON(&moneyRecord)
	if err := moneyRecordService.CreateMoneyRecord(moneyRecord); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMoneyRecord 删除MoneyRecord
// @Tags MoneyRecord
// @Summary 删除MoneyRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.MoneyRecord true "删除MoneyRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moneyRecord/deleteMoneyRecord [delete]
func (moneyRecordApi *MoneyRecordApi) DeleteMoneyRecord(c *gin.Context) {
	var moneyRecord autocode.MoneyRecord
	_ = c.ShouldBindJSON(&moneyRecord)
	if err := moneyRecordService.DeleteMoneyRecord(moneyRecord); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMoneyRecordByIds 批量删除MoneyRecord
// @Tags MoneyRecord
// @Summary 批量删除MoneyRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MoneyRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /moneyRecord/deleteMoneyRecordByIds [delete]
func (moneyRecordApi *MoneyRecordApi) DeleteMoneyRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := moneyRecordService.DeleteMoneyRecordByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMoneyRecord 更新MoneyRecord
// @Tags MoneyRecord
// @Summary 更新MoneyRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.MoneyRecord true "更新MoneyRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /moneyRecord/updateMoneyRecord [put]
func (moneyRecordApi *MoneyRecordApi) UpdateMoneyRecord(c *gin.Context) {
	var moneyRecord autocode.MoneyRecord
	_ = c.ShouldBindJSON(&moneyRecord)
	if err := moneyRecordService.UpdateMoneyRecord(moneyRecord); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMoneyRecord 用id查询MoneyRecord
// @Tags MoneyRecord
// @Summary 用id查询MoneyRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.MoneyRecord true "用id查询MoneyRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /moneyRecord/findMoneyRecord [get]
func (moneyRecordApi *MoneyRecordApi) FindMoneyRecord(c *gin.Context) {
	var moneyRecord autocode.MoneyRecord
	_ = c.ShouldBindQuery(&moneyRecord)
	if err, remoneyRecord := moneyRecordService.GetMoneyRecord(moneyRecord.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remoneyRecord": remoneyRecord}, c)
	}
}

// GetMoneyRecordList 分页获取MoneyRecord列表
// @Tags MoneyRecord
// @Summary 分页获取MoneyRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.MoneyRecordSearch true "分页获取MoneyRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneyRecord/getMoneyRecordList [get]
func (moneyRecordApi *MoneyRecordApi) GetMoneyRecordList(c *gin.Context) {
	var pageInfo autocodeReq.MoneyRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := moneyRecordService.GetMoneyRecordInfoList(pageInfo); err != nil {
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
