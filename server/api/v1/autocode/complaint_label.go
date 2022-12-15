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

type ComplaintLabelApi struct {
}

var complaintLabelService = service.ServiceGroupApp.AutoCodeServiceGroup.ComplaintLabelService


// CreateComplaintLabel 创建ComplaintLabel
// @Tags ComplaintLabel
// @Summary 创建ComplaintLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ComplaintLabel true "创建ComplaintLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /complaintLabel/createComplaintLabel [post]
func (complaintLabelApi *ComplaintLabelApi) CreateComplaintLabel(c *gin.Context) {
	var complaintLabel autocode.ComplaintLabel
	_ = c.ShouldBindJSON(&complaintLabel)
	if err := complaintLabelService.CreateComplaintLabel(complaintLabel); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteComplaintLabel 删除ComplaintLabel
// @Tags ComplaintLabel
// @Summary 删除ComplaintLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ComplaintLabel true "删除ComplaintLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /complaintLabel/deleteComplaintLabel [delete]
func (complaintLabelApi *ComplaintLabelApi) DeleteComplaintLabel(c *gin.Context) {
	var complaintLabel autocode.ComplaintLabel
	_ = c.ShouldBindJSON(&complaintLabel)
	if err := complaintLabelService.DeleteComplaintLabel(complaintLabel); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteComplaintLabelByIds 批量删除ComplaintLabel
// @Tags ComplaintLabel
// @Summary 批量删除ComplaintLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ComplaintLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /complaintLabel/deleteComplaintLabelByIds [delete]
func (complaintLabelApi *ComplaintLabelApi) DeleteComplaintLabelByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := complaintLabelService.DeleteComplaintLabelByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateComplaintLabel 更新ComplaintLabel
// @Tags ComplaintLabel
// @Summary 更新ComplaintLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ComplaintLabel true "更新ComplaintLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /complaintLabel/updateComplaintLabel [put]
func (complaintLabelApi *ComplaintLabelApi) UpdateComplaintLabel(c *gin.Context) {
	var complaintLabel autocode.ComplaintLabel
	_ = c.ShouldBindJSON(&complaintLabel)
	if err := complaintLabelService.UpdateComplaintLabel(complaintLabel); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindComplaintLabel 用id查询ComplaintLabel
// @Tags ComplaintLabel
// @Summary 用id查询ComplaintLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ComplaintLabel true "用id查询ComplaintLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /complaintLabel/findComplaintLabel [get]
func (complaintLabelApi *ComplaintLabelApi) FindComplaintLabel(c *gin.Context) {
	var complaintLabel autocode.ComplaintLabel
	_ = c.ShouldBindQuery(&complaintLabel)
	if err, recomplaintLabel := complaintLabelService.GetComplaintLabel(complaintLabel.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recomplaintLabel": recomplaintLabel}, c)
	}
}

// GetComplaintLabelList 分页获取ComplaintLabel列表
// @Tags ComplaintLabel
// @Summary 分页获取ComplaintLabel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ComplaintLabelSearch true "分页获取ComplaintLabel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /complaintLabel/getComplaintLabelList [get]
func (complaintLabelApi *ComplaintLabelApi) GetComplaintLabelList(c *gin.Context) {
	var pageInfo autocodeReq.ComplaintLabelSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := complaintLabelService.GetComplaintLabelInfoList(pageInfo); err != nil {
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
