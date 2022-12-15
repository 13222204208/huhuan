package minapp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
)

type SecondRoom struct {
	global.GVA_MODEL
	Buy           uint                           `json:"buy" form:"buy" gorm:"column:buy;comment:买家的用户id"`
	Sell          uint                           `json:"sell" form:"sell" gorm:"column:sell;comment:卖家用户的id"`
	RoomName      string                         `json:"roomName" form:"roomName" gorm:"column:room_name;comment:房间名称"`
	OrderId       uint                           `json:"orderId" form:"orderId" gorm:"column:order_id;default:0;comment:闲置订单id"`
	OrderNum      string                         `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:订单号"`
	BuyUser       autocode.MinappUser            `gorm:"foreignKey:Buy"`
	SellUser      autocode.MinappUser            `gorm:"foreignKey:Sell"`
	OnlineMessage []autocode.SecondOnlineMessage `gorm:"foreignKey:Room"`
}
