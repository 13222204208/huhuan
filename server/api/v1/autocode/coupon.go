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

type CouponApi struct {
}

var couponService = service.ServiceGroupApp.AutoCodeServiceGroup.CouponService

// CreateCoupon 创建Coupon
// @Tags Coupon
// @Summary 创建Coupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Coupon true "创建Coupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /coupon/createCoupon [post]
func (couponApi *CouponApi) CreateCoupon(c *gin.Context) {
	var coupon autocode.Coupon
	_ = c.ShouldBindJSON(&coupon)
	if err := couponService.CreateCoupon(coupon); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCoupon 删除Coupon
// @Tags Coupon
// @Summary 删除Coupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Coupon true "删除Coupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /coupon/deleteCoupon [delete]
func (couponApi *CouponApi) DeleteCoupon(c *gin.Context) {
	var coupon autocode.Coupon
	_ = c.ShouldBindJSON(&coupon)
	if err := couponService.DeleteCoupon(coupon); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCouponByIds 批量删除Coupon
// @Tags Coupon
// @Summary 批量删除Coupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Coupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /coupon/deleteCouponByIds [delete]
func (couponApi *CouponApi) DeleteCouponByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := couponService.DeleteCouponByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCoupon 更新Coupon
// @Tags Coupon
// @Summary 更新Coupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Coupon true "更新Coupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /coupon/updateCoupon [put]
func (couponApi *CouponApi) UpdateCoupon(c *gin.Context) {
	var coupon autocode.Coupon
	_ = c.ShouldBindJSON(&coupon)
	if err := couponService.UpdateCoupon(coupon); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCoupon 用id查询Coupon
// @Tags Coupon
// @Summary 用id查询Coupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Coupon true "用id查询Coupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /coupon/findCoupon [get]
func (couponApi *CouponApi) FindCoupon(c *gin.Context) {
	var coupon autocode.Coupon
	_ = c.ShouldBindQuery(&coupon)
	if err, recoupon := couponService.GetCoupon(coupon.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recoupon": recoupon}, c)
	}
}

// GetCouponList 分页获取Coupon列表
// @Tags Coupon
// @Summary 分页获取Coupon列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CouponSearch true "分页获取Coupon列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /coupon/getCouponList [get]
func (couponApi *CouponApi) GetCouponList(c *gin.Context) {
	var pageInfo autocodeReq.CouponSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := couponService.GetCouponInfoList(pageInfo); err != nil {
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

//获取所有优惠券的名称和ID
func (couponApi *CouponApi) GetAllCouponsTitle(c *gin.Context) {
	if err, list := couponService.GetAllCouponsTitle(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"allCoupons": list}, "获取成功", c)
	}
}
