package minapp

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"gorm.io/gorm"
)

type CouponService struct{}

func (s *CouponService) GetMyCoupons(uid uint) (err error, releaseRecord []autocode.ReleaseRecord) {
	err = global.GVA_DB.Where("uid = ?", uid).Find(&releaseRecord).Error
	return
}

//获取充值余额赠送的优惠券列表
func (s *CouponService) GetPayCouponList() (err error, list interface{}) {
	var coupon []autocode.PayCoupon
	err = global.GVA_DB.Find(&coupon).Error
	return err, coupon
}

//邀请用户赠送优惠券
func (s *CouponService) GetGiveCoupon(uid uint) (err error, coupon autocode.PayCoupon) {
	err = global.GVA_DB.Where("type = ?", 2).First(&coupon).Error

	var i autocode.InviteRecord
	var count int64
	global.GVA_DB.Model(&i).Where("invite = ?", uid).Count(&count)
	if float64(count) < coupon.Way {
		coupon.Way = coupon.Way - float64(count)
	}
	return err, coupon
}

//充值余额满足条件时赠送优惠券
func (s *CouponService) TopUpGiveCoupon(uid uint, money float64) (err error) {
	var coupon autocode.PayCoupon
	if errors.Is(global.GVA_DB.Where("money < ?", money).First(&coupon).Error, gorm.ErrRecordNotFound) {
		return errors.New("不满足赠送条件")
	} else {
		var give autocode.ReleaseRecord
		give.Uid = uid
		give.CouponId = coupon.CouponId
		give.Title = coupon.Title
		give.Start = coupon.Start
		give.Over = coupon.Over
		give.Condition = coupon.Condition
		give.Way = coupon.Way
		give.Total = coupon.Total
		give.Remark = "充值余额赠送的优惠券"
		err = global.GVA_DB.Save(&give).Error
		if err != nil {
			return errors.New("赠送优惠失败")
		}
		return err
	}
}
