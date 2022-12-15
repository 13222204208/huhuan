package minapp

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Address struct {
	global.GVA_MODEL
	Uid      uint   `json:"uid" gorm:"comment:用户id"`
	Name     string `json:"name" gorm:"comment:姓名"`
	Phone    string `json:"phone" gorm:"comment:手机号"`
	Country  string `json:"country" gorm:"comment:国家"`
	Province string `json:"province" gorm:"comment:省"`
	City     string `json:"city" gorm:"comment:城市"`
	Area     string `json:"area" gorm:"comment:区"`
	Detail   string `json:"detail" gorm:"comment:地址详情"`
	PostCode string `json:"postCode" form:"postCode" gorm:"column:post_code;comment:邮编"`
	Default  int    `json:"default" gorm:"comment:是否默认"`
}
