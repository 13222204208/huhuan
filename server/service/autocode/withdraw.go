package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type WithdrawService struct {
}

// CreateWithdraw 创建Withdraw记录
// Author [piexlmax](https://github.com/piexlmax)
func (withdrawService *WithdrawService) CreateWithdraw(withdraw autocode.Withdraw) (err error) {
	err = global.GVA_DB.Create(&withdraw).Error
	return err
}

// DeleteWithdraw 删除Withdraw记录
// Author [piexlmax](https://github.com/piexlmax)
func (withdrawService *WithdrawService) DeleteWithdraw(withdraw autocode.Withdraw) (err error) {
	err = global.GVA_DB.Delete(&withdraw).Error
	return err
}

// DeleteWithdrawByIds 批量删除Withdraw记录
// Author [piexlmax](https://github.com/piexlmax)
func (withdrawService *WithdrawService) DeleteWithdrawByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Withdraw{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateWithdraw 更新Withdraw记录
// Author [piexlmax](https://github.com/piexlmax)
func (withdrawService *WithdrawService) UpdateWithdraw(withdraw autocode.Withdraw) (err error) {
	//err = global.GVA_DB.Save(&withdraw).Error
	//return err
	//更改提现记录的状态
	err = global.GVA_DB.Model(&withdraw).Where("id = ?", withdraw.ID).Update("status", withdraw.Status).Error
	return err
}

// GetWithdraw 根据id获取Withdraw记录
// Author [piexlmax](https://github.com/piexlmax)
func (withdrawService *WithdrawService) GetWithdraw(id uint) (err error, withdraw autocode.Withdraw) {
	err = global.GVA_DB.Where("id = ?", id).First(&withdraw).Error
	return
}

// GetWithdrawInfoList 分页获取Withdraw记录
// Author [piexlmax](https://github.com/piexlmax)
func (withdrawService *WithdrawService) GetWithdrawInfoList(info autoCodeReq.WithdrawSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.Withdraw{})
	var withdraws []autocode.Withdraw
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.UserNum != "" {
		db = db.Where("user_num LIKE ?", "%"+info.UserNum+"%")
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.ApplyTime != "" {
		db = db.Where("apply_time LIKE ?", "%"+info.ApplyTime+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("MinAppUser").Order("id DESC").Find(&withdraws).Error
	return err, withdraws, total
}
