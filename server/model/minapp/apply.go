package minapp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
)

//团购接龙报名占位信息
type ApplyHelpGroupon struct {
	global.GVA_MODEL
	Uid         uint                 `json:"uid" form:"uid" gorm:"column:uid;comment:用户id"`
	OrderId     uint                 `json:"orderId" form:"orderId" gorm:"column:order_id;comment:占位的商品订单ID"`
	OrderNum    string               `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:订单号"`
	Total       int                  `json:"total" form:"total" gorm:"column:total; comment:商品数量; default:1"`
	Time        string               `json:"time" form:"time" gorm:"column:time;comment:占位时间;"`
	Status      uint8                `json:"status" form:"status" gorm:"column:status;comment:状态;default:1"`
	HelpGroupon autocode.HelpGroupon `gorm:"foreignKey:OrderId"`
	UserInfo    autocode.MinappUser  `gorm:"foreignKey:Uid"`
}
