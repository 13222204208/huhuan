package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type HelpWorkService struct {
}

// CreateHelpWork 创建HelpWork记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpWorkService *HelpWorkService) CreateHelpWork(helpWork autocode.HelpWork) (err error) {
	err = global.GVA_DB.Create(&helpWork).Error
	return err
}

// DeleteHelpWork 删除HelpWork记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpWorkService *HelpWorkService) DeleteHelpWork(helpWork autocode.HelpWork) (err error) {
	err = global.GVA_DB.Delete(&helpWork).Error
	return err
}

// DeleteHelpWorkByIds 批量删除HelpWork记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpWorkService *HelpWorkService) DeleteHelpWorkByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.HelpWork{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHelpWork 更新HelpWork记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpWorkService *HelpWorkService) UpdateHelpWork(helpWork autocode.HelpWork) (err error) {
	err = global.GVA_DB.Save(&helpWork).Error
	return err
}

// GetHelpWork 根据id获取HelpWork记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpWorkService *HelpWorkService) GetHelpWork(id uint) (err error, helpWork autocode.HelpWork) {
	err = global.GVA_DB.Where("id = ?", id).Preload("UserInfo").First(&helpWork).Error
	return
}

// GetHelpWorkInfoList 分页获取HelpWork记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpWorkService *HelpWorkService) GetHelpWorkInfoList(info autoCodeReq.HelpWorkSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpWork{})
	var helpWorks []autocode.HelpWork
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
	if info.UserNum != "" {
		db = db.Where("uid LIKE ?", "%"+info.UserNum+"%")
	}

	if info.City != "" {
		db = db.Where("city = ?", info.City)
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
	err = db.Limit(limit).Offset(offset).Preload("UserInfo").Preload("CouponInfo").Preload("OnlineMessage").Order("id DESC").Find(&helpWorks).Error
	return err, helpWorks, total
}
