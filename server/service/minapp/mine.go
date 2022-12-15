package minapp

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/silenceper/wechat/v2/pay/order"
	"gorm.io/gorm"
	"log"
)

type MineService struct{}

var TodayMinutes = utils.TodayMinutes()

//获取我发布的团购接龙订单
func (s *MineService) MyHelpGrouponInfoList(info autoCodeReq.HelpGrouponSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpGroupon{})

	//判断已过期的订单
	if info.Status == 5 {
		fmt.Println("日期", TodayMinutes)
		db = db.Where("Time < ?", TodayMinutes)
		db = db.Where("status = ?", 0)
	} else {
		if info.Status >= 0 {
			db = db.Where("status = ?", info.Status)
		}
	}
	//判断获取的是否是我的接单
	if info.Way == "take" {
		db = db.Where("take_id = ?", info.Uid)
	} else {
		db = db.Where("uid = ?", info.Uid)
	}

	if info.Way == "history" {
		db = db.Where("status = ?", 5).Or("status = ?", 6)
	}

	var helpGroupon []autocode.HelpGroupon
	err = db.Count(&total).Error

	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("OnlineMessage", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", 0)

	}).Find(&helpGroupon).Error
	return err, helpGroupon, total
}

//获取我发布的国内捎带 订单
func (s *MineService) MyHelpIncidentallyInfoList(info autoCodeReq.HelpIncidentallySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpIncidentally{})

	//判断已过期的订单
	if info.Status == 5 {
		fmt.Println("日期", TodayMinutes)
		db = db.Where("Time < ?", TodayMinutes)
		db = db.Where("status = ?", 0)
	} else {
		if info.Status >= 0 {
			db = db.Where("status = ?", info.Status)
		}
	}
	//判断获取的是否是我的接单
	if info.Way == "take" {
		db = db.Where("take_id = ?", info.Uid)
	} else {
		db = db.Where("uid = ?", info.Uid)
	}

	var helpIncidentally []autocode.HelpIncidentally
	err = db.Count(&total).Error

	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("OnlineMessage", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", 0)

	}).Find(&helpIncidentally).Error
	return err, helpIncidentally, total
}

//获取我发布的帮我做 订单
func (s *MineService) MyHelpRepairInfoList(info autoCodeReq.HelpRepairSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpRepair{})

	//判断已过期的订单
	if info.Status == 5 {
		fmt.Println("日期", TodayMinutes)
		db = db.Where("Time < ?", TodayMinutes)
		db = db.Where("status = ?", 0)
	} else {
		if info.Status >= 0 {
			db = db.Where("status = ?", info.Status)
		}
	}
	//判断获取的是否是我的接单
	if info.Way == "take" {
		db = db.Where("take_id = ?", info.Uid)
	} else {
		db = db.Where("uid = ?", info.Uid)
	}
	var helpRepair []autocode.HelpRepair
	err = db.Count(&total).Error

	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("OnlineMessage", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", 0)

	}).Find(&helpRepair).Error
	return err, helpRepair, total
}

//获取我发布的帮我做 订单
func (s *MineService) MyHelpWorkInfoList(info autoCodeReq.HelpWorkSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpWork{})

	//判断已过期的订单
	if info.Status == 5 {
		fmt.Println("日期", TodayMinutes)
		db = db.Where("Time < ?", TodayMinutes)
		db = db.Where("status = ?", 0)
	} else {
		if info.Status >= 0 {
			db = db.Where("status = ?", info.Status)
		}
	}
	//判断获取的是否是我的接单
	if info.Way == "take" {
		db = db.Where("take_id = ?", info.Uid)
	} else {
		db = db.Where("uid = ?", info.Uid)
	}
	var helpWork []autocode.HelpWork
	err = db.Count(&total).Error

	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("OnlineMessage", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", 0)

	}).Find(&helpWork).Error
	return err, helpWork, total
}

//获取我发布的帮我买订单
func (s *MineService) MyHelpBuyInfoList(info autoCodeReq.HelpBuySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpBuy{})

	//判断已过期的订单
	if info.Status == 5 {
		fmt.Println("日期", TodayMinutes)
		db = db.Where("Time < ?", TodayMinutes)
		db = db.Where("status = ?", 0)
	} else {
		if info.Status >= 0 {
			db = db.Where("status = ?", info.Status)
		}
	}
	//判断获取的是否是我的接单
	if info.Way == "take" {
		db = db.Where("take_id = ?", info.Uid)
	} else {
		db = db.Where("uid = ?", info.Uid)
	}

	var HelpBuy []autocode.HelpBuy
	err = db.Count(&total).Error

	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("OnlineMessage", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", 0)

	}).Find(&HelpBuy).Error
	return err, HelpBuy, total
}

//获取我发布的帮我取订单
func (s *MineService) MyHelpFetchInfoList(info autoCodeReq.HelpFetchSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpFetch{})
	fmt.Println("帮我取页数", offset, limit)
	//判断已过期的订单
	if info.Status == 5 {
		fmt.Println("日期", TodayMinutes)
		db = db.Where("Time < ?", TodayMinutes)
		db = db.Where("status = ?", 0)
	} else {
		if info.Status >= 0 {
			db = db.Where("status = ?", info.Status)
		}
	}
	//判断获取的是否是我的接单
	if info.Way == "take" {
		db = db.Where("take_id = ?", info.Uid)
	} else {
		db = db.Where("uid = ?", info.Uid)
	}

	var helpFetch []autocode.HelpFetch
	err = db.Count(&total).Error

	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("OnlineMessage", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", 0)

	}).Order("id DESC").Find(&helpFetch).Error
	return err, helpFetch, total
}

//获取我预约的二手闲置订单
func (s *MineService) MyTakeHelpSecond(info autoCodeReq.HelpSecondSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	var take []autocode.TakeHelpSecond
	db := global.GVA_DB.Model(&take)
	db.Where("uid", info.Uid)
	err = db.Count(&total).Error

	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("HelpSecond").Order("id DESC").Find(&take).Error
	return err, take, total
}

type Result struct {
	Id int `json:"id"`
}

//获取我收藏的二手闲置订单
func (s *MineService) MyCollectHelpSecond(info autoCodeReq.HelpSecondSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	var collect []minapp.Collect
	db := global.GVA_DB.Model(&collect)
	db.Where("uid = ?", info.Uid)

	err = db.Count(&total).Error

	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("HelpSecond.UserInfo").Order("id DESC").Find(&collect).Error
	return err, collect, total
}

//获取我发布的二手闲置
func (s *MineService) MyHelpSecondInfoList(info autoCodeReq.HelpSecondSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpSecond{})
	var helpSeconds []autocode.HelpSecond
	db = db.Where("uid = ?", info.Uid)

	if info.History == 1 {
		db = db.Where("status = ?", 5).Or("status = ?", 6)
	} else {
		db = db.Where("status < ?", 4)
	}
	err = db.Count(&total).Error

	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("OnlineMessage", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", 0)

	}).Order("id DESC").Find(&helpSeconds).Error
	return err, helpSeconds, total
}

//获取我预约的组团拼车
func (s *MineService) MyHelpCarpoolTakeList(info autoCodeReq.HelpCarpoolPageInfo) (err error, list interface{}, total int64) {
	fmt.Println("分页信息", info.Page, info.PageSize)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	var car []autocode.CarUsedNum
	db := global.GVA_DB.Model(&car)
	db = db.Where("uid = ?", info.Uid)

	err = db.Count(&total).Error

	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("HelpCarpool").Order("id DESC").Find(&car).Error
	return err, car, total
}

//获取我发布的组团拼车
func (s *MineService) MyHelpCarpoolList(info autoCodeReq.HelpCarpoolPageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	fmt.Println("分页信息2", info.Page, info.PageSize)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpCarpool{})
	var helpCarpools []autocode.HelpCarpool
	db = db.Where("uid = ?", info.Uid)
	if info.History == 1 {
		db = db.Where("status = ?", 5)
		db = db.Or("status = ?", 6)
	} else {
		db = db.Where("status > ?", 0)
	}
	err = db.Count(&total).Error

	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id DESC").Find(&helpCarpools).Error
	return err, helpCarpools, total
}

//取消我发布的组团拼车订单
func (s *MineService) CancelMyHelpCarpool(orderId uint, uid uint) error {
	var carpool autocode.HelpCarpool
	err := global.GVA_DB.Model(&carpool).Where("id = ? And uid = ?", orderId, uid).Update("status", 6).Error
	return err
}

//编辑我发布的组团拼车订单
func (s *MineService) UpdateHelpCarpool(carpool autocode.HelpCarpool) error {
	err := global.GVA_DB.Model(&carpool).Updates(carpool).Error
	return err
}

//编辑我发布的二手闲置订单
func (s *MineService) UpdateHelpSecond(second autocode.HelpSecond) error {
	err := global.GVA_DB.Model(&second).Updates(second).Error
	return err
}

//接单我想要二手闲置
func (s *MineService) TakeHelpSecond(uid, orderId uint) error {
	var take autocode.TakeHelpSecond
	take.Uid = uid
	take.OrderId = orderId

	//判断是否自己的订单，如果是则不能接单
	var helpSecond autocode.HelpSecond
	global.GVA_DB.First(&helpSecond, orderId)
	if helpSecond.Uid == uid {
		return errors.New("这是您发布的订单")
	}

	if !errors.Is(global.GVA_DB.Where("uid = ? AND order_id = ?", uid, orderId).First(&take).Error, gorm.ErrRecordNotFound) {
		return errors.New("此商品你已经想要")
	}
	take.CreationDate = utils.Today()
	err := global.GVA_DB.Save(&take).Error
	return err
}

//接单组团拼车
func (s *MineService) TakeHelpCarpool(orderId, uid uint) error {
	var carUsedNum autocode.CarUsedNum

	//查询拼车是否已经坐满
	var carPool autocode.HelpCarpool
	err := global.GVA_DB.Where("id = ?", orderId).First(&carPool).Error
	if err != nil {
		return errors.New("查询失败")
	} else {
		//判断 是否是自己发布的订单
		if carPool.Uid == uid {
			return errors.New("此订单是您发布的订单")
		}
		//乘满 则更改订单的状态
		global.GVA_DB.Model(&carPool).Where("id = ?", orderId).Update("status", 2)

		if carPool.UseNum <= carPool.CarUsedNum {
			return errors.New("车已乘满")
		}
	}

	if !errors.Is(global.GVA_DB.Where("order_id = ? AND uid = ?", orderId, uid).First(&carUsedNum).Error, gorm.ErrRecordNotFound) {
		return errors.New("你已接单")
	} else {
		var userinfo autocode.MinappUser
		global.GVA_DB.Select("avatar").Where("id", uid).First(&userinfo)
		carUsedNum.OrderId = orderId
		carUsedNum.Uid = uid
		carUsedNum.Avatar = userinfo.Avatar
		err := global.GVA_DB.Save(&carUsedNum).Error
		if err == nil {
			err = global.GVA_DB.Model(&carPool).Where("id = ?", orderId).UpdateColumn("car_used_num", gorm.Expr("car_used_num + ?", 1)).Error

			//预约up
			CreateOrderCenter(uid, "up", carPool.OrderNum, carPool.Price)

		}
		return err
	}
}

//取消编辑布的帮我取订单
func (s *MineService) UpdateHelpFetch(fetch autocode.HelpFetch) error {
	if fetch.Status == 2 {
		var total int64
		global.GVA_DB.Model(&fetch).Where("id = ? AND status >= ?", fetch.ID, 2).Count(&total)
		if total > 0 {
			return errors.New("无法接单")
		} else {
			if err, result := GetFirstUserInfo(fetch.TakeId); err != nil {
				log.Println(err)
			} else {

				fetch.TakeName = result.Name
				fetch.TakePhone = result.Phone

				//查询帮我取订单信息
				var h autocode.HelpFetch
				global.GVA_DB.Where("id = ?", fetch.ID).First(&h)
				fmt.Println("帮我取", h.Uid, h.TakeId)
				//是否是自己发布的订单
				if h.Uid == fetch.TakeId {
					return errors.New("这是您发布的订单")
				}

				//保存我的订单到订单中心
				CreateOrderCenter(fetch.TakeId, "take", h.OrderNum, h.Price)

			}
		}
	}
	//查询订单信息
	var h autocode.HelpFetch
	global.GVA_DB.Where("id = ?", fetch.ID).First(&h)
	//如果状态为4 则订单完成 则把金额发放给接单人
	if fetch.Status == 4 {

		//修改订单中心状态
		ChangeOrderCenterStatus(h.OrderNum, "take")
		//增加接单人的用户余额
		AddUserBalance(h.TakeId, h.Price)
		//保存钱包记录
		var record minapp.Record
		record.Uid = h.TakeId
		record.Type = 4 //接单收益
		record.OrderNum = h.OrderNum
		record.Money = h.Price
		record.Remark = "接单收益"
		SaveMoneyRecord(record)
	}

	if fetch.Status == 6 {
		cancelOrderBalance(h.Uid, h.Price)
	}
	err := global.GVA_DB.Model(&fetch).Updates(fetch).Error
	return err
}

//取消编辑布的帮我买订单
func (s *MineService) UpdateHelpBuy(buy autocode.HelpBuy) error {
	if buy.Status == 2 {
		var total int64
		global.GVA_DB.Model(&buy).Where("id = ? AND status >= 2", buy.ID).Count(&total)
		if total > 0 {
			return errors.New("无法接单")
		} else {
			if err, result := GetFirstUserInfo(buy.TakeId); err != nil {
				log.Println(err)
			} else {

				buy.TakeName = result.Name
				buy.TakePhone = result.Phone

				//查询帮我买的订单
				var h autocode.HelpBuy
				global.GVA_DB.Where("id = ?", buy.ID).First(&h)

				//是否是自己发布的订单
				if h.Uid == buy.TakeId {
					return errors.New("这是您发布的订单")
				}
				//保存我的订单到订单中心
				CreateOrderCenter(buy.TakeId, "take", h.OrderNum, h.Price)
			}
		}
	}
	//如果状态为4 则订单完成 则把金额发放给接单人
	if buy.Status == 4 {
		//查询订单信息
		var h autocode.HelpBuy
		global.GVA_DB.Where("id = ?", buy.ID).First(&h)
		//修改订单中心状态
		ChangeOrderCenterStatus(h.OrderNum, "take")
		//增加接单人的用户余额
		AddUserBalance(h.TakeId, h.Price)
		//保存钱包记录
		var record minapp.Record
		record.Uid = h.TakeId
		record.Type = 4 //接单收益
		record.OrderNum = h.OrderNum
		record.Money = h.Price
		record.Remark = "接单收益"
		SaveMoneyRecord(record)
	}
	err := global.GVA_DB.Model(&buy).Updates(buy).Error
	return err
}

//取消编辑布的帮我修订单
func (s *MineService) UpdateHelpRepair(repair autocode.HelpRepair) error {
	if repair.Status == 2 {
		var total int64
		global.GVA_DB.Model(&repair).Where("id = ? AND status >= 2", repair.ID).Count(&total)
		if total > 0 {
			return errors.New("无法接单")
		} else {
			if err, result := GetFirstUserInfo(repair.TakeId); err != nil {
				log.Println(err)
			} else {

				repair.TakeName = result.Name
				repair.TakePhone = result.Phone

				//查询订单信息
				var h autocode.HelpRepair
				global.GVA_DB.Where("id = ?", repair.ID).First(&h)

				if h.Uid == repair.TakeId {
					return errors.New("这是您发布的订单")
				}

				//保存我的订单到订单中心
				CreateOrderCenter(repair.TakeId, "take", h.OrderNum, h.Price)
			}
		}
	}
	//查询订单信息
	var h autocode.HelpRepair
	global.GVA_DB.Where("id = ?", repair.ID).First(&h)
	//如果状态为4 则订单完成 则把金额发放给接单人
	if repair.Status == 4 {

		//修改订单中心状态
		ChangeOrderCenterStatus(h.OrderNum, "take")
		//增加接单人的用户余额
		AddUserBalance(h.TakeId, h.Price)
		//保存钱包记录
		var record minapp.Record
		record.Uid = h.TakeId
		record.Type = 4 //接单收益
		record.OrderNum = h.OrderNum
		record.Money = h.Price
		record.Remark = "接单收益"
		SaveMoneyRecord(record)
	}

	if repair.Status == 6 {
		cancelOrderBalance(h.Uid, h.Price)
	}

	err := global.GVA_DB.Model(&repair).Updates(repair).Error
	return err
}

//取消编辑国内捎带订单
func (s *MineService) UpdateHelpIncidentally(incidentally autocode.HelpIncidentally) error {
	if incidentally.Status == 2 {
		var i autocode.HelpIncidentally
		var total int64
		global.GVA_DB.Model(&i).Where("id = ? AND status >= 2", incidentally.ID).Count(&total)

		if total > 0 {
			return errors.New("无法接单")
		} else {
			if err, result := GetFirstUserInfo(incidentally.TakeId); err != nil {
				log.Println(err)
			} else {

				incidentally.TakeName = result.Name
				incidentally.TakePhone = result.Phone

				var h autocode.HelpIncidentally
				global.GVA_DB.Where("id = ?", incidentally.ID).First(&h)
				if h.Uid == incidentally.TakeId {
					return errors.New("这是您发布的订单")
				}
				//预约up
				CreateOrderCenter(incidentally.TakeId, "up", incidentally.OrderNum, incidentally.Price)

			}
		}
	}
	err := global.GVA_DB.Model(&incidentally).Updates(incidentally).Error
	return err
}

//取消编辑帮我做订单
func (s *MineService) UpdateHelpWork(work autocode.HelpWork) error {
	if work.Status == 2 {
		var total int64
		global.GVA_DB.Model(&work).Where("id = ? AND status >= 2", work.ID).Count(&total)
		if total > 0 {
			return errors.New("无法接单")
		} else {
			if err, result := GetFirstUserInfo(work.TakeId); err != nil {
				log.Println(err)
			} else {

				work.TakeName = result.Name
				work.TakePhone = result.Phone

				//查询订单信息
				var h autocode.HelpWork
				global.GVA_DB.Where("id = ?", work.ID).First(&h)

				//是否是自己发布的订单
				if h.Uid == work.TakeId {
					return errors.New("这是您发布的订单")
				}

				//保存我的订单到订单中心
				CreateOrderCenter(work.TakeId, "take", h.OrderNum, h.Price)
			}
		}
	}

	//查询订单信息
	var h autocode.HelpWork
	global.GVA_DB.Where("id = ?", work.ID).First(&h)

	//如果状态为4 则订单完成 则把金额发放给接单人
	if work.Status == 4 {

		//修改订单中心状态
		ChangeOrderCenterStatus(h.OrderNum, "take")
		//增加接单人的用户余额
		AddUserBalance(h.TakeId, h.Price)
		//保存钱包记录
		var record minapp.Record
		record.Uid = h.TakeId
		record.Type = 4 //接单收益
		record.OrderNum = h.OrderNum
		record.Money = h.Price
		record.Remark = "接单收益"
		SaveMoneyRecord(record)
	}

	if work.Status == 6 {
		cancelOrderBalance(h.Uid, h.Price)
	}

	err := global.GVA_DB.Model(&work).Updates(work).Error
	return err
}

//乘客发布的组拼车，加价
func (s *MineService) AddPriceHelpCarpool(orderId uint, price float64) error {
	var carpool autocode.HelpCarpool
	//在原本价格上自增
	err := global.GVA_DB.Model(&carpool).Where("id=?", orderId).UpdateColumn("price", gorm.Expr("price + ?", price)).Error
	return err
}

//帮我修订单，加价
func (s *MineService) AddPriceHelpRepair(orderId uint, price float64) error {
	var repair autocode.HelpRepair
	//在原本价格上自增
	err := global.GVA_DB.Model(&repair).Where("id=?", orderId).UpdateColumn("price", gorm.Expr("price + ?", price)).Error
	return err
}

//我发布的帮我取订单 加价
func (s *MineService) AddPriceHelpFetch(orderId uint, price float64) error {
	var fetch autocode.HelpFetch
	err := global.GVA_DB.Model(&fetch).Where("id=?", orderId).UpdateColumn("price", gorm.Expr("price + ?", price)).Error
	return err
}

//我发布的国内捎带 加价
func (s *MineService) AddPriceHelpIncidentally(orderId uint, price float64) error {
	var h autocode.HelpIncidentally
	err := global.GVA_DB.Model(&h).Where("id=?", orderId).UpdateColumn("price", gorm.Expr("price + ?", price)).Error
	return err
}

//我发布的帮我做订单 加价
func (s *MineService) AddPriceHelpWork(orderId uint, price float64) error {
	var work autocode.HelpWork
	err := global.GVA_DB.Model(&work).Where("id=?", orderId).UpdateColumn("price", gorm.Expr("price + ?", price)).Error
	return err
}

//我发布的帮我取订单 加价
func (s *MineService) AddPriceHelpBuy(orderId uint, price float64) error {
	var buy autocode.HelpBuy
	err := global.GVA_DB.Model(&buy).Where("id=?", orderId).UpdateColumn("price", gorm.Expr("price + ?", price)).Error
	return err
}

//订单打赏
func (s *MineService) RewardOrder(openid string, createIP string, reward minapp.Reward) (err error, config order.Config) {
	var u UserService
	u.PayConfig()
	var params order.Params
	reward.OrderNum = "REW" + utils.UserNum()
	//money, _ := strconv.Atoi(fmt.Sprintf("%1.0f", reward.Money*100))
	totalFee := CadChangeRmb(reward.Money)
	params.TotalFee = totalFee
	params.CreateIP = createIP
	params.Body = "金额打赏"
	params.OutTradeNo = reward.OrderNum
	params.OpenID = openid
	params.TradeType = "JSAPI"
	params.NotifyURL = global.GVA_PAY_CONFIG.NotifyURL

	var sendOrder order.Order
	sendOrder = order.Order{
		global.GVA_PAY_CONFIG,
	}

	err = global.GVA_DB.Save(&reward).Error
	config, err = sendOrder.BridgeConfig(&params)
	fmt.Println(err)
	return err, config
}

//订单加价
func (s *MineService) AddPriceOrder(openid string, createIP string, reward minapp.Reward) (err error, config order.Config) {
	var u UserService
	u.PayConfig()
	var params order.Params
	reward.OrderNum = "ADD" + utils.UserNum()
	//money, _ := strconv.Atoi(fmt.Sprintf("%1.0f", reward.Money*100))
	totalFee := CadChangeRmb(reward.Money)
	params.TotalFee = totalFee
	params.CreateIP = createIP
	params.Body = "订单加价"
	params.OutTradeNo = reward.OrderNum
	params.OpenID = openid
	params.TradeType = "JSAPI"
	params.NotifyURL = global.GVA_PAY_CONFIG.NotifyURL

	var sendOrder order.Order
	sendOrder = order.Order{
		global.GVA_PAY_CONFIG,
	}

	err = global.GVA_DB.Save(&reward).Error
	config, err = sendOrder.BridgeConfig(&params)
	fmt.Println(err)
	return err, config
}

//接单前验证是否填写身份信息
func (s *MineService) FindUserVerifying(uid uint) (err error, verify int) {
	var user autocode.MinappUser
	err = global.GVA_DB.Select("verify").Where("id = ?", uid).Find(&user).Error
	return err, user.Verify
}

func GetFirstUserInfo(uid uint) (err error, user autocode.MinappUser) {
	var reqUser autocode.MinappUser
	err = global.GVA_DB.First(&reqUser, "id = ?", uid).Error

	return err, reqUser
}

func CadChangeRmb(money float64) string {
	//加元转换成人民币
	exchangeRate := GetMinAppSet()
	money = money * exchangeRate.ExchangeRate //汇率相乘
	money = utils.Decimal(money) * 100
	totalFee := utils.FloatToString(money)
	return totalFee
}

//退款到用户余额
func (s *MineService) RefundToUser(uid uint, orderNum string, money float64) {
	money = utils.Decimal(money)
	var refund autocode.Refund
	refund.Uid = uid
	refund.OrderNum = orderNum
	refund.Money = money
	refund.Remark = "退款到余额"
	global.GVA_DB.Save(&refund)
	var user autocode.MinappUser
	global.GVA_DB.Model(&user).Where("id", uid).UpdateColumn("balance", gorm.Expr("balance + ?", money))

}

//修改订单中心的状态
func ChangeOrderCenterStatus(orderNum string, way string) {
	var center autocode.OrderCenter
	global.GVA_DB.Model(&center).Where("order_num = ? AND way = ?", orderNum, way).Update("status", 1)
}

//增加用户余额
func AddUserBalance(uid uint, money float64) {
	var user autocode.MinappUser
	global.GVA_DB.Model(&user).Where("id", uid).UpdateColumn("balance", gorm.Expr("balance + ?", money))
}

//保存钱包记录
func SaveMoneyRecord(record minapp.Record) {
	record.CreationDate = utils.Today()
	record.Status = 1
	global.GVA_DB.Save(&record)
}

//取消订单退款到余额
func cancelOrderBalance(uid uint, money float64) {
	var r minapp.Record
	r.Uid = uid
	r.Money = money
	r.Type = 3
	r.Status = 1
	r.Remark = "取消订单退款到余额"
	err := global.GVA_DB.Save(&r).Error
	fmt.Println("取消订单退款到余额状态", err)
	if err == nil {
		AddUserBalance(uid, money)
	}
}
