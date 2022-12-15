package autocode

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type CheckMerchantService struct {
}

// CreateCheckMerchant 创建CheckMerchant记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkMerchantService *CheckMerchantService) CreateCheckMerchant(checkMerchant autocode.CheckMerchant) (err error) {
	if !errors.Is(global.GVA_DB.Where("uid", checkMerchant.Uid).First(&checkMerchant).Error, gorm.ErrRecordNotFound) {
		return errors.New("已经提交了身份验证")
	}

	checkMerchant.Time = utils.Today()
	err = global.GVA_DB.Create(&checkMerchant).Error

	//修改用户是否提交了申请
	var user autocode.MinappUser
	global.GVA_DB.Model(&user).Where("id", checkMerchant.Uid).UpdateColumn("check_merchant", 2)

	return err
}

// DeleteCheckMerchant 删除CheckMerchant记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkMerchantService *CheckMerchantService) DeleteCheckMerchant(checkMerchant autocode.CheckMerchant) (err error) {
	err = global.GVA_DB.Delete(&checkMerchant).Error
	return err
}

// DeleteCheckMerchantByIds 批量删除CheckMerchant记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkMerchantService *CheckMerchantService) DeleteCheckMerchantByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CheckMerchant{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCheckMerchant 更新CheckMerchant记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkMerchantService *CheckMerchantService) UpdateCheckMerchant(checkMerchant autocode.CheckMerchant) (err error) {
	err = global.GVA_DB.Save(&checkMerchant).Error
	//如果审核的状态为 1 代表未审核，2 代表审核通过，3 审核拒绝
	if err != nil {
		return err
	} else {
		var checkState int
		if checkMerchant.Status == 2 {
			checkState = 1
		}
		var user autocode.MinappUser

		err = global.GVA_DB.Model(&user).Where("id", checkMerchant.Uid).Update("check_merchant", checkState).Error
		return err
	}
}

// GetCheckMerchant 根据id获取CheckMerchant记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkMerchantService *CheckMerchantService) GetCheckMerchant(id uint) (err error, checkMerchant autocode.CheckMerchant) {
	err = global.GVA_DB.Where("id = ?", id).First(&checkMerchant).Error
	return
}

// GetCheckMerchantInfoList 分页获取CheckMerchant记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkMerchantService *CheckMerchantService) GetCheckMerchantInfoList(info autoCodeReq.CheckMerchantSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.CheckMerchant{})
	var checkMerchants []autocode.CheckMerchant
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}

	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}

	if info.Phone != "" {
		db = db.Where("phone = ?", info.Phone)
	}

	if info.Time != "" {
		db = db.Where("time = ?", info.Time)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("MinAppUser").Find(&checkMerchants).Error
	return err, checkMerchants, total
}
