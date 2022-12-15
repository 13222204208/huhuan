package minapp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
)

type SecondMessage struct {
	global.GVA_MODEL
	Uid         uint                `json:"uid" form:"uid" gorm:"column:uid;default:0;comment:留言的用户ID"`
	OrderNum    string              `json:"orderNum" form:"orderNum" gorm:"column:order_num"`
	Pid         uint                `json:"pid" form:"pid" gorm:"column:pid;default:0;comment:留言的父ID"`
	Tid         uint                `json:"tid" form:"tid" gorm:"column:tid;default:0;comment:接收留言的用户Id"`
	Contents    string              `json:"contents" form:"contents" gorm:"column:contents;comment:留言的内容"`
	MessageTime string              `json:"messageTime" form:"messageTime" gorm:"column:message_time;comment:留言的时间"`
	Status      int                 `json:"status" form:"status" gorm:"column:status;default:0;comment:留言的状态"`
	MinAppUser  autocode.MinappUser `gorm:"foreignKey:Uid"`
}
