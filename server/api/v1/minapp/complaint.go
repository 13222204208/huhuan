package minapp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	minAppRequest "github.com/flipped-aurora/gin-vue-admin/server/model/minapp/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ComplaintApi struct{}

func (a *ComplaintApi) Labels(c *gin.Context) {
	if err, result := complaintService.GetAllComplaintLabels(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"labels": result}, "获取成功", c)
	}
}

func (a *ComplaintApi) Create(c *gin.Context) {
	var o minAppRequest.ComplaintOrder
	_ = c.ShouldBindJSON(&o)
	if err := utils.Verify(o, utils.MinAppComplaintOrder); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := c.MustGet("id").(uint) //投诉人ID
	complaint := &autocode.Complaint{
		Label:           o.Label,
		Contents:        o.Contents,
		OrderNum:        o.OrderNum,
		ComplainantId:   uid,
		ByComplainantId: o.ByComplainantId, //被投诉人ID
	}

	if err := complaintService.CreateComplaintOrder(*complaint); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithMessage("成功", c)
	}
}
