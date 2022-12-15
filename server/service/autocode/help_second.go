package autocode

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HelpSecondService struct {
}

// CreateHelpSecond 创建HelpSecond记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpSecondService *HelpSecondService) CreateHelpSecond(helpSecond autocode.HelpSecond) (err error) {
	err = global.GVA_DB.Create(&helpSecond).Error

	//保存发布订单到订单中心
	if err != nil {
		global.GVA_LOG.Error("发布订单错误", zap.Error(err))
	} else {
		CreateOrderCenter(helpSecond.Uid, "create", helpSecond.OrderNum, helpSecond.Price)
	}

	return err
}

// DeleteHelpSecond 删除HelpSecond记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpSecondService *HelpSecondService) DeleteHelpSecond(helpSecond autocode.HelpSecond) (err error) {
	err = global.GVA_DB.Delete(&helpSecond).Error
	return err
}

// DeleteHelpSecondByIds 批量删除HelpSecond记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpSecondService *HelpSecondService) DeleteHelpSecondByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.HelpSecond{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHelpSecond 更新HelpSecond记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpSecondService *HelpSecondService) UpdateHelpSecond(helpSecond autocode.HelpSecond) (err error) {
	err = global.GVA_DB.Save(&helpSecond).Error
	return err
}

// GetHelpSecond 根据id获取HelpSecond记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpSecondService *HelpSecondService) GetHelpSecond(id uint) (err error, helpSecond autocode.HelpSecond) {
	err = global.GVA_DB.Where("id = ?", id).Preload("UserInfo").First(&helpSecond).Error

	return err, helpSecond
}

func (helpSecondService *HelpSecondService) GetHelpSecondStatus(uid, orderId uint) (err error, helpSecond autocode.HelpSecond) {
	err = global.GVA_DB.Where("id = ?", orderId).Preload("UserInfo").First(&helpSecond).Error

	var collect minapp.Collect
	if errors.Is(global.GVA_DB.Where("order_id = ? And uid = ?", orderId, uid).First(&collect).Error, gorm.ErrRecordNotFound) {
		helpSecond.Collect = 0
	} else {
		helpSecond.Collect = 1
	}

	var take autocode.TakeHelpSecond
	if errors.Is(global.GVA_DB.Where("order_id = ? And uid = ?", orderId, uid).First(&take).Error, gorm.ErrRecordNotFound) {
		helpSecond.Take = 0
	} else {
		helpSecond.Take = 1
	}

	return err, helpSecond
}

// GetHelpSecondInfoList 分页获取HelpSecond记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpSecondService *HelpSecondService) GetHelpSecondInfoList(info autoCodeReq.HelpSecondSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpSecond{})
	var helpSeconds []autocode.HelpSecond
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderNum != "" {
		db = db.Where("order_num LIKE ?", "%"+info.OrderNum+"%")
	}
	if info.GoodsName != "" {
		db = db.Where("goods_name LIKE ?", "%"+info.GoodsName+"%")
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
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

	if info.Category != "" {
		db = db.Where("category LIKE ?", "%"+info.Category+"%")
	}

	if info.Uid > 0 {
		db = db.Where("uid = ?", info.Uid)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("UserInfo").Preload("OnlineMessage").Order("id DESC").Find(&helpSeconds).Error
	return err, helpSeconds, total
}
