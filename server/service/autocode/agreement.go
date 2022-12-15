package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type AgreementService struct {
}

// CreateAgreement 创建Agreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (agreementService *AgreementService) CreateAgreement(agreement autocode.Agreement) (err error) {
	err = global.GVA_DB.Create(&agreement).Error
	return err
}

// DeleteAgreement 删除Agreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (agreementService *AgreementService)DeleteAgreement(agreement autocode.Agreement) (err error) {
	err = global.GVA_DB.Delete(&agreement).Error
	return err
}

// DeleteAgreementByIds 批量删除Agreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (agreementService *AgreementService)DeleteAgreementByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Agreement{},"id in ?",ids.Ids).Error
	return err
}

// UpdateAgreement 更新Agreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (agreementService *AgreementService)UpdateAgreement(agreement autocode.Agreement) (err error) {
	err = global.GVA_DB.Save(&agreement).Error
	return err
}

// GetAgreement 根据id获取Agreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (agreementService *AgreementService)GetAgreement(id uint) (err error, agreement autocode.Agreement) {
	err = global.GVA_DB.Where("id = ?", id).First(&agreement).Error
	return
}

// GetAgreementInfoList 分页获取Agreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (agreementService *AgreementService)GetAgreementInfoList(info autoCodeReq.AgreementSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Agreement{})
    var agreements []autocode.Agreement
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Status != nil {
        db = db.Where("status = ?",info.Status)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&agreements).Error
	return err, agreements, total
}
