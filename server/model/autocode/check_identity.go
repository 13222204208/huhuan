// 自动生成模板CheckIdentity
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CheckIdentity 结构体
// 如果含有time.Time 请自行import time包
type CheckIdentity struct {
	global.GVA_MODEL
	Name        string     `json:"name" form:"name" gorm:"column:name;comment:姓名;"`
	Phone       string     `json:"phone" form:"phone" gorm:"column:phone;comment:联系方式;"`
	Selfie      string     `json:"selfie" form:"selfie" gorm:"column:selfie;comment:自拍照;"`
	Certificate string     `json:"certificate" form:"certificate" gorm:"column:certificate;comment:证件照 包括 营业执照;"`
	Mail        string     `json:"mail" form:"mail" gorm:"column:mail;comment:邮箱;"`
	Time        string     `json:"time" form:"time" gorm:"column:time;comment:申请日期;"`
	Uid         uint       `json:"uid" form:"uid" gorm:"column:uid;comment:申请人ID;"`
	Gender      string     `json:"gender" form:"gender" gorm:"column:gender;comment:性别;"`
	Card        string     `json:"card" form:"card" gorm:"column:card;comment:身份证"`
	Status      int        `json:"status" form:"status" gorm:"column:status;comment:申请的状态;size:2;default:1"`
	MinAppUser  MinappUser `json:"minAppUser" form:"minAppUser" gorm:"foreignKey:Uid"`
}
