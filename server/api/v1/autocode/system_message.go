package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/minapp"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type SystemMessageApi struct {
}

var systemMessageService = service.ServiceGroupApp.AutoCodeServiceGroup.SystemMessageService

// CreateSystemMessage 创建SystemMessage
// @Tags SystemMessage
// @Summary 创建SystemMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SystemMessage true "创建SystemMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /systemMessage/createSystemMessage [post]
func (systemMessageApi *SystemMessageApi) CreateSystemMessage(c *gin.Context) {
	var r autocodeReq.SendSystemMessageRequest
	var message []autocode.SystemMessage
	_ = c.ShouldBindJSON(&r)
	if r.Resource == "1" {
		for _, i := range r.Uid {
			uid, _ := strconv.Atoi(i)
			message = append(message, autocode.SystemMessage{
				Uid:      uid,
				Contents: r.Contents,
			})
		}
		_ = minapp.UnreadMessageCount(r.Uid)
		//minapp.SystemSendMessage(r.Uid, r.Contents)
	}
	if err := systemMessageService.CreateSystemMessage(message); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSystemMessage 删除SystemMessage
// @Tags SystemMessage
// @Summary 删除SystemMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SystemMessage true "删除SystemMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /systemMessage/deleteSystemMessage [delete]
func (systemMessageApi *SystemMessageApi) DeleteSystemMessage(c *gin.Context) {
	var systemMessage autocode.SystemMessage
	_ = c.ShouldBindJSON(&systemMessage)
	if err := systemMessageService.DeleteSystemMessage(systemMessage); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSystemMessageByIds 批量删除SystemMessage
// @Tags SystemMessage
// @Summary 批量删除SystemMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SystemMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /systemMessage/deleteSystemMessageByIds [delete]
func (systemMessageApi *SystemMessageApi) DeleteSystemMessageByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := systemMessageService.DeleteSystemMessageByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSystemMessage 更新SystemMessage
// @Tags SystemMessage
// @Summary 更新SystemMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SystemMessage true "更新SystemMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /systemMessage/updateSystemMessage [put]
func (systemMessageApi *SystemMessageApi) UpdateSystemMessage(c *gin.Context) {
	var systemMessage autocode.SystemMessage
	_ = c.ShouldBindJSON(&systemMessage)
	if err := systemMessageService.UpdateSystemMessage(systemMessage); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSystemMessage 用id查询SystemMessage
// @Tags SystemMessage
// @Summary 用id查询SystemMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.SystemMessage true "用id查询SystemMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /systemMessage/findSystemMessage [get]
func (systemMessageApi *SystemMessageApi) FindSystemMessage(c *gin.Context) {
	var systemMessage autocode.SystemMessage
	_ = c.ShouldBindQuery(&systemMessage)
	if err, resystemMessage := systemMessageService.GetSystemMessage(systemMessage.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resystemMessage": resystemMessage}, c)
	}
}

// GetSystemMessageList 分页获取SystemMessage列表
// @Tags SystemMessage
// @Summary 分页获取SystemMessage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.SystemMessageSearch true "分页获取SystemMessage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /systemMessage/getSystemMessageList [get]
func (systemMessageApi *SystemMessageApi) GetSystemMessageList(c *gin.Context) {
	var pageInfo autocodeReq.SystemMessageSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := systemMessageService.GetSystemMessageInfoList(pageInfo); err != nil {
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
