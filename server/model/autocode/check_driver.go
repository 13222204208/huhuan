// 自动生成模板CheckDriver
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CheckDriver 结构体
// 如果含有time.Time 请自行import time包
type CheckDriver struct {
	global.GVA_MODEL
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:姓名;"`
	Phone      string     `json:"phone" form:"phone" gorm:"column:phone;comment:联系方式;"`
	Gender     string     `json:"gender" form:"gender" gorm:"column:gender;comment:性别"`
	CarType    string     `json:"carType" form:"carType" gorm:"column:car_type;comment:车型"`
	Licence    string     `json:"licence" form:"licence" gorm:"column:licence;comment:驾照;"`
	CarImage   string     `json:"carImage" form:"carImage" gorm:"column:car_image;comment:车辆图片"`
	Charter    string     `json:"charter" form:"charter" gorm:"column:charter;comment:营业执照;"`
	Mail       string     `json:"mail" form:"mail" gorm:"column:mail;comment:邮箱;"`
	Time       string     `json:"time" form:"time" gorm:"column:time;comment:申请日期;"`
	Uid        uint       `json:"uid" form:"uid" gorm:"column:uid;comment:申请人ID;"`
	Status     int        `json:"status" form:"status" gorm:"column:status;default:1;comment:申请的状态;size:2;"`
	MinAppUser MinappUser `json:"minAppUser" form:"minAppUser" gorm:"foreignKey:Uid"`
}
