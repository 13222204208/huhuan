package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type HelpRepairService struct {
}

// CreateHelpRepair 创建HelpRepair记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpRepairService *HelpRepairService) CreateHelpRepair(helpRepair autocode.HelpRepair) (err error) {
	err = global.GVA_DB.Create(&helpRepair).Error
	return err
}

// DeleteHelpRepair 删除HelpRepair记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpRepairService *HelpRepairService) DeleteHelpRepair(helpRepair autocode.HelpRepair) (err error) {
	err = global.GVA_DB.Delete(&helpRepair).Error
	return err
}

// DeleteHelpRepairByIds 批量删除HelpRepair记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpRepairService *HelpRepairService) DeleteHelpRepairByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.HelpRepair{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHelpRepair 更新HelpRepair记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpRepairService *HelpRepairService) UpdateHelpRepair(helpRepair autocode.HelpRepair) (err error) {
	err = global.GVA_DB.Save(&helpRepair).Error
	return err
}

// GetHelpRepair 根据id获取HelpRepair记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpRepairService *HelpRepairService) GetHelpRepair(id uint) (err error, helpRepair autocode.HelpRepair) {
	err = global.GVA_DB.Where("id = ?", id).Preload("UserInfo").First(&helpRepair).Error
	return
}

// GetHelpRepairInfoList 分页获取HelpRepair记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpRepairService *HelpRepairService) GetHelpRepairInfoList(info autoCodeReq.HelpRepairSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpRepair{})
	var helpRepairs []autocode.HelpRepair
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
		db = db.Where("uid LIKE ?", "%"+info.UserNum+"%")
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

	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
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
	err = db.Limit(limit).Offset(offset).Order("id DESC").Preload("UserInfo").Preload("CouponInfo").Preload("OnlineMessage").Find(&helpRepairs).Error
	return err, helpRepairs, total
}
