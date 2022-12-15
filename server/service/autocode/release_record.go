package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ReleaseRecordService struct {
}

// CreateReleaseRecord 创建ReleaseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (releaseRecordService *ReleaseRecordService) CreateReleaseRecord(releaseRecord []autocode.ReleaseRecord) (err error) {
	err = global.GVA_DB.Create(&releaseRecord).Error
	return err
}

// DeleteReleaseRecord 删除ReleaseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (releaseRecordService *ReleaseRecordService) DeleteReleaseRecord(releaseRecord autocode.ReleaseRecord) (err error) {
	err = global.GVA_DB.Delete(&releaseRecord).Error
	return err
}

// DeleteReleaseRecordByIds 批量删除ReleaseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (releaseRecordService *ReleaseRecordService) DeleteReleaseRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.ReleaseRecord{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateReleaseRecord 更新ReleaseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (releaseRecordService *ReleaseRecordService) UpdateReleaseRecord(releaseRecord autocode.ReleaseRecord) (err error) {
	err = global.GVA_DB.Save(&releaseRecord).Error
	return err
}

// GetReleaseRecord 根据id获取ReleaseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (releaseRecordService *ReleaseRecordService) GetReleaseRecord(id uint) (err error, releaseRecord autocode.ReleaseRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&releaseRecord).Error
	return
}

// GetReleaseRecordInfoList 分页获取ReleaseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (releaseRecordService *ReleaseRecordService) GetReleaseRecordInfoList(info autoCodeReq.ReleaseRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.ReleaseRecord{})

	if info.Uid != 0 {
		db = db.Where("uid = ?", info.Uid)
	}
	var releaseRecords []autocode.ReleaseRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Preload("UserInfo").Offset(offset).Find(&releaseRecords).Error
	return err, releaseRecords, total
}
