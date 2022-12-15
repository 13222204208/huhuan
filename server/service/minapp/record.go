package minapp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	minAppRequest "github.com/flipped-aurora/gin-vue-admin/server/model/minapp/request"
)

type RecordService struct{}

func (s *RecordService) CreateUserTopUpRecord(record minapp.Record) error {
	err := global.GVA_DB.Save(&record).Error
	return err
}

//获取用户钱包记录
func (s *RecordService) GetUserMoneyRecord(info minAppRequest.UserMoneyRecord) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	var record []minapp.Record
	db := global.GVA_DB.Model(&minapp.Record{})
	if info.Type > 0 {
		db = db.Where("type = ?", info.Type)
	}

	db = db.Where("uid", info.Uid)
	//db = db.Where("status", 1)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&record).Error
	return err, record, total
}
