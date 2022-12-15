package minapp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	minAppRequest "github.com/flipped-aurora/gin-vue-admin/server/model/minapp/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

type SignApi struct{}

//更改签到提醒状态
func (s *SignApi) ChangeSignInState(c *gin.Context) {
	var state minAppRequest.ChangeSignInState
	_ = c.ShouldBindJSON(&state)

	var stateNum uint8
	if state.State == "on" {
		stateNum = 1
	} else {
		stateNum = 0
	}

	uid := c.MustGet("id").(uint)

	if err := minAppUserService.ChangeSignInState(uid, stateNum); err != nil {
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithMessage("成功", c)
	}
}

//用户签到
func (s *SignApi) CreateUserSignIn(c *gin.Context) {
	var r minapp.SignInRecord

	if err, result := minAppSettingService.GetMinappSet(1); err != nil {
		response.FailWithMessage("获取每日签到积分失败", c)
	} else {
		uid := c.MustGet("id").(uint)
		r.Uid = uid
		r.SignInTime = utils.Today()
		r.Integral = result.Integral
		if err := userSignInService.CreateSignInRecord(r); err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithMessage("成功", c)
		}
	}
}

//获取 可用积分兑换的优惠券
func (s *SignApi) GetIntegralCouponList(c *gin.Context) {
	if err, result := userSignInService.GetIntegralCoupon(); err != nil {
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithData(gin.H{"list": result}, c)
	}
}

//使用积分兑换优惠券
func (s *SignApi) ExchangeCoupon(c *gin.Context) {
	var coupon minAppRequest.ExchangeCoupon
	_ = c.ShouldBindJSON(&coupon)
	uid := c.MustGet("id").(uint)
	//查询用户积分
	if err, result := minAppUserService.GetMinappUser(uid); err != nil {
		response.FailWithMessage("找不到用户", c)
		return
	} else {
		//查询优惠券信息
		if err, r := AutoCodeCouponService.GetCoupon(coupon.CouponId); err != nil {
			response.FailWithMessage("失败", c)
			return
		} else {
			//判断积分是否足够
			if result.Integral < r.Integral {
				response.FailWithMessage("积分不足", c)
				return
			} else {
				//添加优惠券信息到 发放的优惠券信息中
				var record autocode.ReleaseRecord
				record.Uid = uid
				record.CouponId = coupon.CouponId
				record.Title = r.Title
				record.Condition = r.Condition
				record.Start = r.Start
				record.Over = r.Over
				record.Way = r.Way
				var recordData []autocode.ReleaseRecord
				recordData = append(recordData, record)
				if err := releaseRecordService.CreateReleaseRecord(recordData); err != nil {

					response.FailWithMessage("兑换失败", c)
				} else {
					//兑换优惠券后减少用户积分
					userSignInService.UpdateUserIntegral(uid, r.Integral)
					response.OkWithMessage("成功", c)
				}
			}
		}
	}
}

//我的签到记录
func (s *SignApi) GetMySignInRecords(c *gin.Context) {
	uid := c.MustGet("id").(uint)
	if err, result := userSignInService.GetMySignInRecords(uid); err != nil {
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithData(gin.H{"list": result}, c)
	}
}
