package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type HelpFetchService struct {
}

// CreateHelpFetch 创建HelpFetch记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpFetchService *HelpFetchService) CreateHelpFetch(helpFetch autocode.HelpFetch) (err error) {
	err = global.GVA_DB.Create(&helpFetch).Error
	return err
}

// DeleteHelpFetch 删除HelpFetch记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpFetchService *HelpFetchService) DeleteHelpFetch(helpFetch autocode.HelpFetch) (err error) {
	err = global.GVA_DB.Delete(&helpFetch).Error
	return err
}

// DeleteHelpFetchByIds 批量删除HelpFetch记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpFetchService *HelpFetchService) DeleteHelpFetchByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.HelpFetch{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHelpFetch 更新HelpFetch记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpFetchService *HelpFetchService) UpdateHelpFetch(helpFetch autocode.HelpFetch) (err error) {
	err = global.GVA_DB.Save(&helpFetch).Error
	return err
}

// GetHelpFetch 根据id获取HelpFetch记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpFetchService *HelpFetchService) GetHelpFetch(id uint) (err error, helpFetch autocode.HelpFetch) {
	err = global.GVA_DB.Where("id = ?", id).Preload("UserInfo").First(&helpFetch).Error
	return
}

// GetHelpFetchInfoList 分页获取HelpFetch记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpFetchService *HelpFetchService) GetHelpFetchInfoList(info autoCodeReq.HelpFetchSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpFetch{})
	var helpFetchs []autocode.HelpFetch
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

	if info.City != "" {
		db = db.Where("city = ?", info.City)
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
	err = db.Limit(limit).Offset(offset).Preload("UserInfo").Preload("CouponInfo").Preload("OnlineMessage").Order("id DESC").Find(&helpFetchs).Error
	return err, helpFetchs, total
}
