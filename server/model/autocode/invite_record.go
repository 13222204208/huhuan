package autocode

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type InviteRecord struct {
	global.GVA_MODEL
	Uid      uint       `json:"uid" form:"uid" gorm:"column:uid;comment:受到邀请的用户id;"`
	Invite   uint       `json:"invite" form:"invite" gorm:"column:invite;comment:邀请人的用户id"`
	Status   int        `json:"status" form:"status" gorm:"column:status;default:1;comment:邀请的状态"`
	UserInfo MinappUser `gorm:"foreignKey:Uid"`
}
