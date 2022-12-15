package minapp

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	"gorm.io/gorm"
)

type SignService struct{}

func (s *SignService) CreateSignInRecord(record minapp.SignInRecord) (err error) {
	//global.GVA_DB.First(&record)
	fmt.Println(record.Uid, record.SignInTime)
	if !errors.Is(global.GVA_DB.Where("uid = ? AND sign_in_time = ?", record.Uid, record.SignInTime).First(&record).Error, gorm.ErrRecordNotFound) {
		return errors.New("今天已签到")
	}

	var minAppUser autocode.MinappUser
	//给用户添加积分
	global.GVA_DB.Model(&minAppUser).Where("id = ?", record.Uid).UpdateColumn("integral", gorm.Expr("integral + ?", record.Integral))
	//签到天数
	global.GVA_DB.Model(&minAppUser).Where("id = ?", record.Uid).UpdateColumn("sign_in_day", gorm.Expr("sign_in_day + ?", 1))
	err = global.GVA_DB.Create(&record).Error
	return err
}

//获取签到页面，所有的积分兑换优惠券
func (s *SignService) GetIntegralCoupon() (err error, list interface{}) {
	var integral []autocode.Coupon
	err = global.GVA_DB.Where("integral > ?", 0).Find(&integral).Error
	return err, integral
}

//兑换优惠券时减少用户积分
func (s *SignService) UpdateUserIntegral(uid uint, integral uint) {
	var minAppUser autocode.MinappUser
	global.GVA_DB.Model(&minAppUser).Where("id = ?", uid).UpdateColumn("integral", gorm.Expr("integral - ?", integral))
}

//获取我的签到记录信息
func (s *SignService) GetMySignInRecords(uid uint) (err error, list interface{}) {
	var record []minapp.SignInRecord
	err = global.GVA_DB.Where("uid = ?", uid).Find(&record).Error
	return err, record
}
