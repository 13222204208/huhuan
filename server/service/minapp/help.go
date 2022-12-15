package minapp

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HelpService struct{}

//创建帮我修订单
func (h *HelpService) CreateHelpRepair(helpRepair *autocode.HelpRepair) error {
	helpRepair.CreationDate = utils.Today()
	err := global.GVA_DB.Create(&helpRepair).Error

	//保存发布订单到订单中心
	if err != nil {
		global.GVA_LOG.Error("发布订单错误", zap.Error(err))
	} else {
		CreateOrderCenter(helpRepair.Uid, "create", helpRepair.OrderNum, helpRepair.CouponAmount)
	}

	return err
}

//创建帮我做订单
func (h *HelpService) CreateHelpWork(helpWork *autocode.HelpWork) error {
	helpWork.CreationDate = utils.Today()
	err := global.GVA_DB.Create(&helpWork).Error

	//保存发布订单到订单中心
	if err != nil {
		global.GVA_LOG.Error("发布订单错误", zap.Error(err))
	} else {
		CreateOrderCenter(helpWork.Uid, "create", helpWork.OrderNum, helpWork.CouponAmount)
	}

	return err
}

//创建帮我取订单
func (h *HelpService) CreateHelpFetch(helpFetch *autocode.HelpFetch) error {
	helpFetch.CreationDate = utils.Today()
	err := global.GVA_DB.Create(&helpFetch).Error

	//保存发布订单到订单中心
	if err != nil {
		global.GVA_LOG.Error("发布订单错误", zap.Error(err))
	} else {
		CreateOrderCenter(helpFetch.Uid, "create", helpFetch.OrderNum, helpFetch.CouponAmount)
	}

	return err
}

//创建帮我买订单
func (h *HelpService) CreateHelpBuy(helpBuy *autocode.HelpBuy) error {
	helpBuy.CreationDate = utils.Today()
	err := global.GVA_DB.Create(&helpBuy).Error

	//保存发布订单到订单中心
	if err != nil {
		global.GVA_LOG.Error("发布订单错误", zap.Error(err))
	} else {
		//发布订单时优惠的金额
		CreateOrderCenter(helpBuy.Uid, "create", helpBuy.OrderNum, helpBuy.CouponAmount)
	}

	return err
}

//二手闲置增加收藏
func (h *HelpService) CreateCollect(uid uint, orderId uint) (err error) {
	var collect minapp.Collect

	if !errors.Is(global.GVA_DB.Where("uid = ? AND order_id = ?", uid, orderId).First(&collect).Error, gorm.ErrRecordNotFound) {
		return errors.New("此商品你已收藏")
	}

	collect.Uid = uid
	collect.OrderId = orderId
	err = global.GVA_DB.Save(&collect).Error
	return err
}

//二手闲置取消收藏
func (h *HelpService) CancelCollect(uid uint, orderId uint) (err error) {
	var collect minapp.Collect
	err = global.GVA_DB.Where("order_id = ? AND uid = ?", orderId, uid).Delete(&collect).Error

	return err
}

//二手闲置的留言
func (h *HelpService) CreateSecondMessage(message minapp.SecondMessage) (err error) {
	err = global.GVA_DB.Save(&message).Error
	return err
}

//获取二手闲置订单的所有留言
func (h *HelpService) GetSecondOrderMessage(orderNum string) (err error, list interface{}) {
	var message []minapp.SecondMessage
	err = global.GVA_DB.Where("order_num=?", orderNum).Preload("MinAppUser").Find(&message).Error
	return err, message
}

//团购接龙报名占位
func (h *HelpService) ApplyHelpGroupon(uid uint, orderId uint, total int) (err error) {
	var apply minapp.ApplyHelpGroupon
	fmt.Println(uid, orderId, total)
	if !errors.Is(global.GVA_DB.Where("uid = ? AND order_id = ?", uid, orderId).First(&apply).Error, gorm.ErrRecordNotFound) {
		return errors.New("此商品你已经占位")
	}

	//判断此商品团购是否已经结束
	var groupon autocode.HelpGroupon
	if !errors.Is(global.GVA_DB.Where("id = ? AND start_time > ? AND over_time < ? ", orderId, utils.TodayTime(), utils.TodayTime()).First(&groupon).Error, gorm.ErrRecordNotFound) {
		global.GVA_DB.Model(&groupon).Where("id = ?", orderId).Update("status", 5)

		return errors.New("此商品团购已结束")
	} else {
		global.GVA_DB.Where("id = ?", orderId).First(&groupon)

		if groupon.Uid == uid {
			return errors.New("是您发布的订单")
		}

		if groupon.GoodsTotal <= groupon.GrouponTotal {
			global.GVA_DB.Model(&groupon).Where("id = ?", orderId).Update("status", 2)
			return errors.New("此团购已结束")
		}
		apply.OrderNum = groupon.OrderNum

		t1 := groupon.GrouponTotal + total
		if t1 <= groupon.GoodsTotal {
			global.GVA_DB.Model(&groupon).Where("id = ?", orderId).UpdateColumn("groupon_total", gorm.Expr("groupon_total + ?", total))
		} else {

			return errors.New("剩余团购套数不足")
		}
	}
	apply.OrderId = orderId
	apply.Uid = uid
	apply.Total = total
	apply.Time = utils.TodayTime()
	err = global.GVA_DB.Save(&apply).Error

	if err == nil {
		//团购接龙预约
		CreateOrderCenter(uid, "up", groupon.OrderNum, 0)

	}

	return err
}

//团购接龙更新 查看的总次数，及跟团人数
func (h *HelpService) UpdateGrouponFollow(orderId uint) {
	var groupon autocode.HelpGroupon
	global.GVA_DB.Model(&groupon).Where("id = ?", orderId).Update("follow", gorm.Expr("follow + ?", 1))

}

func (h *HelpService) UpdateGrouponLook(orderId uint) {
	var groupon autocode.HelpGroupon
	global.GVA_DB.Model(&groupon).Where("id = ?", orderId).Update("look", gorm.Expr("look + ?", 1))

}

//获取团购接龙订单占位
func (h *HelpService) TakeGrouponInfo(oid uint) {
	var apply minapp.ApplyHelpGroupon
	global.GVA_DB.Where("order_id = ?").Preload("Userinfo").Find(&apply)
}
