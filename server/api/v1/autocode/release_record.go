package autocode

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ReleaseRecordApi struct {
}

var releaseRecordService = service.ServiceGroupApp.AutoCodeServiceGroup.ReleaseRecordService

// CreateReleaseRecord 创建ReleaseRecord
// @Tags ReleaseRecord
// @Summary 创建ReleaseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ReleaseRecord true "创建ReleaseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /releaseRecord/createReleaseRecord [post]
func (releaseRecordApi *ReleaseRecordApi) CreateReleaseRecord(c *gin.Context) {
	var r autocodeReq.SendReleaseRecord
	_ = c.ShouldBindJSON(&r)
	var record []autocode.ReleaseRecord

	//查询优惠券信息
	if err, result := couponService.GetCoupon(r.CouponId); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		for _, i := range r.Uid {
			record = append(record, autocode.ReleaseRecord{
				Uid:       uint(i),
				CouponId:  r.CouponId,
				Title:     result.Title,
				Condition: result.Condition,
				Start:     result.Start,
				Over:      result.Over,
				Way:       result.Way,
				Total:     &r.Total,
				Status:    &r.Status,
			})
		}
	}

	fmt.Println("数据", record)
	if err := releaseRecordService.CreateReleaseRecord(record); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteReleaseRecord 删除ReleaseRecord
// @Tags ReleaseRecord
// @Summary 删除ReleaseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ReleaseRecord true "删除ReleaseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /releaseRecord/deleteReleaseRecord [delete]
func (releaseRecordApi *ReleaseRecordApi) DeleteReleaseRecord(c *gin.Context) {
	var releaseRecord autocode.ReleaseRecord
	_ = c.ShouldBindJSON(&releaseRecord)
	if err := releaseRecordService.DeleteReleaseRecord(releaseRecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteReleaseRecordByIds 批量删除ReleaseRecord
// @Tags ReleaseRecord
// @Summary 批量删除ReleaseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ReleaseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /releaseRecord/deleteReleaseRecordByIds [delete]
func (releaseRecordApi *ReleaseRecordApi) DeleteReleaseRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := releaseRecordService.DeleteReleaseRecordByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateReleaseRecord 更新ReleaseRecord
// @Tags ReleaseRecord
// @Summary 更新ReleaseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ReleaseRecord true "更新ReleaseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /releaseRecord/updateReleaseRecord [put]
func (releaseRecordApi *ReleaseRecordApi) UpdateReleaseRecord(c *gin.Context) {
	var releaseRecord autocode.ReleaseRecord
	_ = c.ShouldBindJSON(&releaseRecord)
	if err := releaseRecordService.UpdateReleaseRecord(releaseRecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindReleaseRecord 用id查询ReleaseRecord
// @Tags ReleaseRecord
// @Summary 用id查询ReleaseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ReleaseRecord true "用id查询ReleaseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /releaseRecord/findReleaseRecord [get]
func (releaseRecordApi *ReleaseRecordApi) FindReleaseRecord(c *gin.Context) {
	var releaseRecord autocode.ReleaseRecord
	_ = c.ShouldBindQuery(&releaseRecord)
	if err, rereleaseRecord := releaseRecordService.GetReleaseRecord(releaseRecord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rereleaseRecord": rereleaseRecord}, c)
	}
}

// GetReleaseRecordList 分页获取ReleaseRecord列表
// @Tags ReleaseRecord
// @Summary 分页获取ReleaseRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ReleaseRecordSearch true "分页获取ReleaseRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /releaseRecord/getReleaseRecordList [get]
func (releaseRecordApi *ReleaseRecordApi) GetReleaseRecordList(c *gin.Context) {
	var pageInfo autocodeReq.ReleaseRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := releaseRecordService.GetReleaseRecordInfoList(pageInfo); err != nil {
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
