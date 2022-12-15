package minapp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	minAppRequest "github.com/flipped-aurora/gin-vue-admin/server/model/minapp/request"
	"github.com/gin-gonic/gin"
)

type CouponApi struct{}

func (couponApi *CouponApi) MyCoupon(c *gin.Context) {
	var coupon minAppRequest.GetOneCoupon
	_ = c.ShouldBindQuery(&coupon)
	if coupon.CouponId > 0 {
		if err, result := AutoCodeCouponService.GetCoupon(coupon.CouponId); err != nil {
			response.FailWithMessage("获取失败", c)
		} else {
			response.OkWithData(gin.H{"detail": result}, c)
		}
	} else {
		uid := c.MustGet("id").(uint)
		if err, result := couponService.GetMyCoupons(uid); err != nil {
			response.FailWithMessage("获取失败", c)
		} else {
			response.OkWithDetailed(gin.H{"couponList": result}, "获取成功", c)
		}
	}

}

//充值余额送优惠券活动
func (couponApi *CouponApi) GetPayCouponList(c *gin.Context) {
	if err, result := couponService.GetPayCouponList(); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"list": result}, "获取成功", c)
	}
}

//邀请用户赠送优惠券
func (couponApi *CouponApi) GetGiveCoupon(c *gin.Context) {
	uid := c.MustGet("id").(uint)
	if err, result := couponService.GetGiveCoupon(uid); err != nil {
		response.FailWithMessage("获取失败", c)

	} else {
		response.OkWithDetailed(gin.H{"list": result}, "获取成功", c)
	}
}
