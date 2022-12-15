package minapp

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type SignInRecord struct {
	global.GVA_MODEL
	Uid        uint   `json:"uid" gorm:"comment:用户id"`
	SignInTime string `json:"signInTime" gorm:"column:sign_in_time;comment:签到时间"`
	Integral   uint   `json:"integral" gorm:"column:integral;comment:获得的积分"`
}
