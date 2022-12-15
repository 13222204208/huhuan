package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SystemMessageService struct {
}

// CreateSystemMessage 创建SystemMessage记录
// Author [piexlmax](https://github.com/piexlmax)
func (systemMessageService *SystemMessageService) CreateSystemMessage(systemMessage []autocode.SystemMessage) (err error) {
	err = global.GVA_DB.Create(&systemMessage).Error
	return err
}

// DeleteSystemMessage 删除SystemMessage记录
// Author [piexlmax](https://github.com/piexlmax)
func (systemMessageService *SystemMessageService) DeleteSystemMessage(systemMessage autocode.SystemMessage) (err error) {
	err = global.GVA_DB.Delete(&systemMessage).Error
	return err
}

// DeleteSystemMessageByIds 批量删除SystemMessage记录
// Author [piexlmax](https://github.com/piexlmax)
func (systemMessageService *SystemMessageService) DeleteSystemMessageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.SystemMessage{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSystemMessage 更新SystemMessage记录
// Author [piexlmax](https://github.com/piexlmax)
func (systemMessageService *SystemMessageService) UpdateSystemMessage(systemMessage autocode.SystemMessage) (err error) {
	err = global.GVA_DB.Save(&systemMessage).Error
	return err
}

// GetSystemMessage 根据id获取SystemMessage记录
// Author [piexlmax](https://github.com/piexlmax)
func (systemMessageService *SystemMessageService) GetSystemMessage(id uint) (err error, systemMessage autocode.SystemMessage) {
	err = global.GVA_DB.Where("id = ?", id).First(&systemMessage).Error
	return
}

// GetSystemMessageInfoList 分页获取SystemMessage记录
// Author [piexlmax](https://github.com/piexlmax)
func (systemMessageService *SystemMessageService) GetSystemMessageInfoList(info autoCodeReq.SystemMessageSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.SystemMessage{})
	var systemMessages []autocode.SystemMessage
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&systemMessages).Error
	return err, systemMessages, total
}
