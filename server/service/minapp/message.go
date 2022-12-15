package minapp

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type MessageService struct {
}

//保存二手闲置私信聊天的内容
func (s *MessageService) SaveSecondOnlineMessage(uid, tid uint, room string, contents string, orderNum string) (err error) {
	var online autocode.SecondOnlineMessage
	online.Tid = tid
	online.Uid = uid
	online.Room = room
	online.OrderNum = orderNum
	online.Contents = contents
	online.MessageTime = utils.TodayTime()
	err = global.GVA_DB.Save(&online).Error
	return err
}

//创建二手闲置私聊天的房间
func (s *MessageService) SaveSecondRoom(room minapp.SecondRoom) (err error, roomId uint) {
	room.RoomName = utils.CreateRoomName(int(room.Sell), int(room.Buy))
	if !errors.Is(global.GVA_DB.Where("room_name = ? ", room.RoomName).First(&room).Error, gorm.ErrRecordNotFound) {
		return errors.New("房间已创建"), room.ID
	} else {
		r := global.GVA_DB.Create(&room)
		return r.Error, room.ID
	}
}

//获取二手闲置私信房间信息及内容
func (s *MessageService) GetSecondRoomMessage(roomId uint) (err error, roomInfo interface{}, message interface{}) {
	var room minapp.SecondRoom
	//查询房间的信息
	global.GVA_DB.Preload("OnlineMessage").First(&room, roomId)
	//查询房间的聊天信息
	var roomMessage []autocode.SecondOnlineMessage
	err = global.GVA_DB.Where("room = ?", roomId).Find(&roomMessage).Error
	return err, room, roomMessage
}

//修改二手闲置私房间信息为已读
func (s *MessageService) UpdateSecondRoomMessage(roomId uint, uid uint) error {
	var roomMessage autocode.SecondOnlineMessage
	fmt.Println("房间号", roomId)
	err := global.GVA_DB.Model(&roomMessage).Where("room = ? ", roomId).Not("uid = ?", uid).Update("status", 1).Error
	return err
}

//私信列表
func (s *MessageService) GetOnlineMessage(uid uint) (err error, list interface{}) {
	var room []minapp.SecondRoom
	err = global.GVA_DB.Where("buy = ?", uid).Or("sell = ?", uid).Preload("OnlineMessage", "status = ?", 0).Preload("BuyUser").Preload("SellUser").Order("id DESC").Find(&room).Error
	return err, room
}

//获取房间信息内容
func (s *MessageService) GetRoomOnlineMessage(room uint) (err error, list interface{}) {
	var message []autocode.SecondOnlineMessage
	err = global.GVA_DB.Where("room = ?", room).Preload("MinAppUser").Find(&message).Error
	return err, message
}

//获取系统消息
func (s *MessageService) GetSystemMessage(uid uint) (err error, list interface{}) {
	var m []autocode.SystemMessage
	err = global.GVA_DB.Where("uid = ?", uid).Order("id DESC").Find(&m).Error
	return err, m
}

func (s *MessageService) GetSystemMessageDetail(id uint) (err error, m autocode.SystemMessage) {
	err = global.GVA_DB.Where("id = ?", id).First(&m).Error
	global.GVA_DB.Model(&m).Where("id = ?", id).Update("status", 1)
	return err, m
}

//获取留言列表
func (s *MessageService) GetLeaveMessage(uid uint) (err error, list interface{}) {
	var m []minapp.SecondMessage
	err = global.GVA_DB.Where("tid = ?", uid).Preload("MinAppUser").Find(&m).Error
	return err, m
}

func (s *MessageService) GetLeaveMessageDetail(id uint) (err error, m minapp.SecondMessage) {
	err = global.GVA_DB.Where("id = ?", id).Preload("MinAppUser").First(&m).Error
	global.GVA_DB.Model(&m).Where("id = ?", id).Update("status", 1)

	return err, m
}

//获取所有消息的未读地数量
func (s *MessageService) GetUnreadMessageCount(uid string) (total int64) {
	var countOne int64
	var sm autocode.SystemMessage
	global.GVA_DB.Model(&sm).Where("uid=? AND status=?", uid, 0).Count(&countOne)

	var countTwo int64
	var hm minapp.SecondMessage
	global.GVA_DB.Model(&hm).Where("tid=? AND status=?", uid, 0).Not("uid=?", uid).Count(&countTwo)

	var countThree int64
	var som autocode.SecondOnlineMessage
	global.GVA_DB.Model(&som).Where("tid=? AND status=?", uid, 0).Not("uid=?", uid).Count(&countThree)
	return countOne + countTwo + countThree
}

//删除自己的留言消息
func (s *MessageService) DeleteMessage(id string, uid uint) (err error) {
	var message minapp.SecondMessage

	if errors.Is(global.GVA_DB.Where("id = ? AND uid = ?", id, uid).First(&message).Error, gorm.ErrRecordNotFound) {
		return errors.New("这不是你的留言")
	}

	err = global.GVA_DB.Where("id = ? AND uid = ?", id, uid).Or("pid = ?", id).Delete(&message).Error
	return err
}
