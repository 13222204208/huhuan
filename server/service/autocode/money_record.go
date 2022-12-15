package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
)

type MoneyRecordService struct {
}

// CreateMoneyRecord 创建MoneyRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyRecordService *MoneyRecordService) CreateMoneyRecord(moneyRecord autocode.MoneyRecord) (err error) {
	err = global.GVA_DB.Create(&moneyRecord).Error
	return err
}

// DeleteMoneyRecord 删除MoneyRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyRecordService *MoneyRecordService) DeleteMoneyRecord(moneyRecord autocode.MoneyRecord) (err error) {
	err = global.GVA_DB.Delete(&moneyRecord).Error
	return err
}

// DeleteMoneyRecordByIds 批量删除MoneyRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyRecordService *MoneyRecordService) DeleteMoneyRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.MoneyRecord{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMoneyRecord 更新MoneyRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyRecordService *MoneyRecordService) UpdateMoneyRecord(moneyRecord autocode.MoneyRecord) (err error) {
	err = global.GVA_DB.Save(&moneyRecord).Error
	return err
}

// GetMoneyRecord 根据id获取MoneyRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyRecordService *MoneyRecordService) GetMoneyRecord(id uint) (err error, moneyRecord autocode.MoneyRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&moneyRecord).Error
	return
}

// GetMoneyRecordInfoList 分页获取MoneyRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyRecordService *MoneyRecordService) GetMoneyRecordInfoList(info autoCodeReq.MoneyRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&minapp.Record{})
	var moneyRecords []minapp.Record
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Order("id DESC").Preload("UserInfo").Offset(offset).Find(&moneyRecords).Error
	return err, moneyRecords, total
}
