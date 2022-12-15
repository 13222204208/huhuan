package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PayCouponApi struct {
}

var payCouponService = service.ServiceGroupApp.AutoCodeServiceGroup.PayCouponService

// CreatePayCoupon 创建PayCoupon
// @Tags PayCoupon
// @Summary 创建PayCoupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.PayCoupon true "创建PayCoupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payCoupon/createPayCoupon [post]
func (payCouponApi *PayCouponApi) CreatePayCoupon(c *gin.Context) {
	var payCoupon autocode.PayCoupon
	_ = c.ShouldBindJSON(&payCoupon)
	//查询优惠券信息
	if err, result := couponService.GetCoupon(payCoupon.CouponId); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		payCoupon.Title = result.Title
		payCoupon.Condition = result.Condition
		payCoupon.Start = result.Start
		payCoupon.Over = result.Over
		payCoupon.Way = result.Way
		if err := payCouponService.CreatePayCoupon(payCoupon); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
		} else {
			response.OkWithMessage("创建成功", c)
		}
	}

}

// DeletePayCoupon 删除PayCoupon
// @Tags PayCoupon
// @Summary 删除PayCoupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.PayCoupon true "删除PayCoupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payCoupon/deletePayCoupon [delete]
func (payCouponApi *PayCouponApi) DeletePayCoupon(c *gin.Context) {
	var payCoupon autocode.PayCoupon
	_ = c.ShouldBindJSON(&payCoupon)
	if err := payCouponService.DeletePayCoupon(payCoupon); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePayCouponByIds 批量删除PayCoupon
// @Tags PayCoupon
// @Summary 批量删除PayCoupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayCoupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /payCoupon/deletePayCouponByIds [delete]
func (payCouponApi *PayCouponApi) DeletePayCouponByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := payCouponService.DeletePayCouponByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePayCoupon 更新PayCoupon
// @Tags PayCoupon
// @Summary 更新PayCoupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.PayCoupon true "更新PayCoupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payCoupon/updatePayCoupon [put]
func (payCouponApi *PayCouponApi) UpdatePayCoupon(c *gin.Context) {
	var payCoupon autocode.PayCoupon
	_ = c.ShouldBindJSON(&payCoupon)
	if err, result := couponService.GetCoupon(payCoupon.CouponId); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		payCoupon.Title = result.Title
		payCoupon.Condition = result.Condition
		payCoupon.Start = result.Start
		payCoupon.Over = result.Over
		payCoupon.Way = result.Way
		if err := payCouponService.UpdatePayCoupon(payCoupon); err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Error(err))
			response.FailWithMessage("更新失败", c)
		} else {
			response.OkWithMessage("更新成功", c)
		}
	}
}

// FindPayCoupon 用id查询PayCoupon
// @Tags PayCoupon
// @Summary 用id查询PayCoupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.PayCoupon true "用id查询PayCoupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payCoupon/findPayCoupon [get]
func (payCouponApi *PayCouponApi) FindPayCoupon(c *gin.Context) {
	var payCoupon autocode.PayCoupon
	_ = c.ShouldBindQuery(&payCoupon)
	if err, repayCoupon := payCouponService.GetPayCoupon(payCoupon.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repayCoupon": repayCoupon}, c)
	}
}

// GetPayCouponList 分页获取PayCoupon列表
// @Tags PayCoupon
// @Summary 分页获取PayCoupon列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.PayCouponSearch true "分页获取PayCoupon列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payCoupon/getPayCouponList [get]
func (payCouponApi *PayCouponApi) GetPayCouponList(c *gin.Context) {
	var pageInfo autocodeReq.PayCouponSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := payCouponService.GetPayCouponInfoList(pageInfo); err != nil {
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
