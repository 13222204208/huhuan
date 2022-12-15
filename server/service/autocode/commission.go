package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type CommissionService struct {
}

// CreateCommission 创建Commission记录
// Author [piexlmax](https://github.com/piexlmax)
func (commissionService *CommissionService) CreateCommission(commission autocode.Commission) (err error) {
	err = global.GVA_DB.Create(&commission).Error
	return err
}

// DeleteCommission 删除Commission记录
// Author [piexlmax](https://github.com/piexlmax)
func (commissionService *CommissionService)DeleteCommission(commission autocode.Commission) (err error) {
	err = global.GVA_DB.Delete(&commission).Error
	return err
}

// DeleteCommissionByIds 批量删除Commission记录
// Author [piexlmax](https://github.com/piexlmax)
func (commissionService *CommissionService)DeleteCommissionByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Commission{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCommission 更新Commission记录
// Author [piexlmax](https://github.com/piexlmax)
func (commissionService *CommissionService)UpdateCommission(commission autocode.Commission) (err error) {
	err = global.GVA_DB.Save(&commission).Error
	return err
}

// GetCommission 根据id获取Commission记录
// Author [piexlmax](https://github.com/piexlmax)
func (commissionService *CommissionService)GetCommission(id uint) (err error, commission autocode.Commission) {
	err = global.GVA_DB.Where("id = ?", id).First(&commission).Error
	return
}

// GetCommissionInfoList 分页获取Commission记录
// Author [piexlmax](https://github.com/piexlmax)
func (commissionService *CommissionService)GetCommissionInfoList(info autoCodeReq.CommissionSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Commission{})
    var commissions []autocode.Commission
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&commissions).Error
	return err, commissions, total
}
