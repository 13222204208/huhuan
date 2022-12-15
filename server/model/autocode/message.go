package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SecondOnlineMessage struct {
	global.GVA_MODEL
	Uid         uint       `json:"uid" form:"uid" gorm:"column:uid;default:0;comment:留言的用户ID"`
	Tid         uint       `json:"tid" form:"tid" gorm:"column:tid;default:0;comment:消息接收人的ID"`
	OrderNum    string     `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:订单号"`
	Room        string     `json:"room" form:"room" gorm:"column:room;comment:在线聊天的房间id "`
	Contents    string     `json:"contents" form:"contents" gorm:"column:contents;comment:留言的内容;type:text;"`
	MessageTime string     `json:"messageTime" form:"messageTime" gorm:"column:message_time;comment:留言的时间"`
	Status      uint       `json:"status" form:"status" gorm:"column:status;default:0;comment:是否已读"`
	MinAppUser  MinappUser `gorm:"foreignKey:Uid"`
}
