package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

type HelpGrouponService struct {
}

// CreateHelpGroupon 创建HelpGroupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpGrouponService *HelpGrouponService) CreateHelpGroupon(helpGroupon autocode.HelpGroupon) (err error) {
	err = global.GVA_DB.Create(&helpGroupon).Error

	//保存发布订单到订单中心
	if err != nil {
		global.GVA_LOG.Error("发布订单错误", zap.Error(err))
	} else {
		CreateOrderCenter(helpGroupon.Uid, "create", helpGroupon.OrderNum, helpGroupon.Price)
	}

	return err
}

// DeleteHelpGroupon 删除HelpGroupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpGrouponService *HelpGrouponService) DeleteHelpGroupon(helpGroupon autocode.HelpGroupon) (err error) {
	err = global.GVA_DB.Delete(&helpGroupon).Error
	return err
}

// DeleteHelpGrouponByIds 批量删除HelpGroupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpGrouponService *HelpGrouponService) DeleteHelpGrouponByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.HelpGroupon{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHelpGroupon 更新HelpGroupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpGrouponService *HelpGrouponService) UpdateHelpGroupon(helpGroupon autocode.HelpGroupon) (err error) {
	err = global.GVA_DB.Save(&helpGroupon).Error
	return err
}

// GetHelpGroupon 根据id获取HelpGroupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpGrouponService *HelpGrouponService) GetHelpGroupon(id uint) (err error, helpGroupon autocode.HelpGroupon) {
	err = global.GVA_DB.Where("id = ?", id).First(&helpGroupon).Error
	return
}

// GetHelpGrouponInfoList 分页获取HelpGroupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpGrouponService *HelpGrouponService) GetHelpGrouponInfoList(info autoCodeReq.HelpGrouponSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpGroupon{})
	var helpGroupons []autocode.HelpGroupon
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderNum != "" {
		db = db.Where("order_num LIKE ?", "%"+info.OrderNum+"%")
	}
	if info.GoodsName != "" {
		db = db.Where("goods_name LIKE ?", "%"+info.GoodsName+"%")
	}
	if info.TakeName != "" {
		db = db.Where("take_name LIKE ?", "%"+info.TakeName+"%")
	}
	if info.UserNum != "" {
		db = db.Where("user_num LIKE ?", "%"+info.UserNum+"%")
	}
	if info.Status > 0 {
		if info.Status == 5 {
			db = db.Where("time < ?", utils.Today())
		} else {
			db = db.Where("status = ?", info.Status)
		}
	}
	if info.Area != "" {
		db = db.Where("area LIKE ?", "%"+info.Area+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("UserInfo").Find(&helpGroupons).Error
	return err, helpGroupons, total
}

//获取我预约的团购接龙商品
func (helpGrouponService *HelpGrouponService) GetTakeHelpGrouponInfoList(info autoCodeReq.HelpGrouponSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	var apply minapp.ApplyHelpGroupon
	db := global.GVA_DB.Model(&apply)
	// 如果有条件搜索 下方会自动创建搜索语句

	db = db.Where("uid", info.Uid)

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("HelpGroupon").Find(&apply).Error
	return err, apply, total
}

func (helpGrouponService *HelpGrouponService) GetMinAppHelpGrouponInfoList(info autoCodeReq.HelpGrouponSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpGroupon{})
	var helpGroupon []autocode.HelpGroupon
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderNum != "" {
		db = db.Where("order_num LIKE ?", "%"+info.OrderNum+"%")
	}
	if info.GoodsName != "" {
		db = db.Where("goods_name LIKE ?", "%"+info.GoodsName+"%")
	}
	if info.TakeName != "" {
		db = db.Where("take_name LIKE ?", "%"+info.TakeName+"%")
	}

	if info.City != "" {
		db = db.Where("city = ?", info.City)
	}
	if info.UserNum != "" {
		db = db.Where("user_num LIKE ?", "%"+info.UserNum+"%")
	}
	db = db.Where("start_time < ?", utils.TodayTime())
	db = db.Where("over_time > ?", utils.TodayTime())
	if info.Status > 0 {
		if info.Status == 5 {
			db = db.Where("over_time < ?", utils.TodayTime())
		} else {
			db = db.Where("status = ?", info.Status)
		}
	}
	if info.Area != "" {
		db = db.Where("area LIKE ?", "%"+info.Area+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("UserInfo").Preload("OnlineMessage").Order("id DESC").Find(&helpGroupon).Error
	return err, helpGroupon, total
}
