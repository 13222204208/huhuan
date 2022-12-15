package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HelpMeService struct {
}

// CreateHelpMe 创建HelpMe记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpMeService *HelpMeService) CreateHelpMe(helpMe autocode.HelpMe) (err error) {
	err = global.GVA_DB.Create(&helpMe).Error
	return err
}

// DeleteHelpMe 删除HelpMe记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpMeService *HelpMeService) DeleteHelpMe(helpMe autocode.HelpMe) (err error) {
	err = global.GVA_DB.Delete(&helpMe).Error
	return err
}

// DeleteHelpMeByIds 批量删除HelpMe记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpMeService *HelpMeService) DeleteHelpMeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.HelpMe{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHelpMe 更新HelpMe记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpMeService *HelpMeService) UpdateHelpMe(helpMe autocode.HelpMe) (err error) {
	err = global.GVA_DB.Save(&helpMe).Error
	return err
}

// GetHelpMe 根据id获取HelpMe记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpMeService *HelpMeService) GetHelpMe(id uint) (err error, helpMe autocode.HelpMe) {
	err = global.GVA_DB.Where("id = ?", id).First(&helpMe).Error
	return
}

// GetHelpMeInfoList 分页获取HelpMe记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpMeService *HelpMeService) GetHelpMeInfoList(info autoCodeReq.HelpMeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpMe{})
	var helpMes []autocode.HelpMe
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&helpMes).Error
	return err, helpMes, total
}

func (helpMeService *HelpMeService) GetHelpMeTitle() (err error, list interface{}) {
	var helpMes []autocode.HelpMe
	err = global.GVA_DB.Where("id > 4").Find(&helpMes).Error
	return err, helpMes
}
