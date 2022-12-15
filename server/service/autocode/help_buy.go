package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type HelpBuyService struct {
}

// CreateHelpBuy 创建HelpBuy记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpBuyService *HelpBuyService) CreateHelpBuy(helpBuy autocode.HelpBuy) (err error) {
	err = global.GVA_DB.Create(&helpBuy).Error
	return err
}

// DeleteHelpBuy 删除HelpBuy记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpBuyService *HelpBuyService) DeleteHelpBuy(helpBuy autocode.HelpBuy) (err error) {
	err = global.GVA_DB.Delete(&helpBuy).Error
	return err
}

// DeleteHelpBuyByIds 批量删除HelpBuy记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpBuyService *HelpBuyService) DeleteHelpBuyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.HelpBuy{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHelpBuy 更新HelpBuy记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpBuyService *HelpBuyService) UpdateHelpBuy(helpBuy autocode.HelpBuy) (err error) {
	err = global.GVA_DB.Save(&helpBuy).Error
	return err
}

// GetHelpBuy 根据id获取HelpBuy记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpBuyService *HelpBuyService) GetHelpBuy(id uint) (err error, helpBuy autocode.HelpBuy) {
	err = global.GVA_DB.Where("id = ?", id).Preload("UserInfo").First(&helpBuy).Error
	return
}

// GetHelpBuyInfoList 分页获取HelpBuy记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpBuyService *HelpBuyService) GetHelpBuyInfoList(info autoCodeReq.HelpBuySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpBuy{})
	var helpBuys []autocode.HelpBuy
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderNum != "" {
		db = db.Where("order_num LIKE ?", "%"+info.OrderNum+"%")
	}
	if info.Linkman != "" {
		db = db.Where("linkman LIKE ?", "%"+info.Linkman+"%")
	}
	if info.ContactPhone != "" {
		db = db.Where("contact_phone = ?", info.ContactPhone)
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
	if info.Status > 0 {
		if info.Status == 5 {
			db = db.Where("time < ?", utils.TodayTime())
		} else {
			db = db.Where("status = ?", info.Status)
		}
	}
	if info.ReceiveAddress != "" {
		db = db.Where("receive_address LIKE ?", "%"+info.ReceiveAddress+"%")
	}

	if info.Area != "" {
		db = db.Where("area LIKE ?", "%"+info.Area+"%")
	}
	if info.CreationDate != "" {
		db = db.Where("creation_date = ?", info.CreationDate)
	}

	if info.Uid > 0 {
		db = db.Where("uid = ?", info.Uid)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("UserInfo").Preload("CouponInfo").Preload("OnlineMessage").Order("id DESC").Find(&helpBuys).Error
	return err, helpBuys, total
}
