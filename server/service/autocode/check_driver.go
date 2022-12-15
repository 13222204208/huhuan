package autocode

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type CheckDriverService struct {
}

// CreateCheckDriver 创建CheckDriver记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkDriverService *CheckDriverService) CreateCheckDriver(checkDriver autocode.CheckDriver) (err error) {
	if !errors.Is(global.GVA_DB.Where("uid = ?", checkDriver.Uid).First(&checkDriver).Error, gorm.ErrRecordNotFound) {
		return errors.New("已经提交了身份验证")
	}
	checkDriver.Time = utils.Today()
	err = global.GVA_DB.Create(&checkDriver).Error

	//修改用户是否提交了申请
	var user autocode.MinappUser
	global.GVA_DB.Model(&user).Where("id", checkDriver.Uid).UpdateColumn("check_driver", 2)

	return err
}

// DeleteCheckDriver 删除CheckDriver记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkDriverService *CheckDriverService) DeleteCheckDriver(checkDriver autocode.CheckDriver) (err error) {
	err = global.GVA_DB.Delete(&checkDriver).Error
	return err
}

// DeleteCheckDriverByIds 批量删除CheckDriver记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkDriverService *CheckDriverService) DeleteCheckDriverByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CheckDriver{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCheckDriver 更新CheckDriver记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkDriverService *CheckDriverService) UpdateCheckDriver(checkDriver autocode.CheckDriver) (err error) {
	err = global.GVA_DB.Save(&checkDriver).Error

	//如果审核的状态为 1 代表未审核，2 代表审核通过，3 审核拒绝
	if err != nil {
		return err
	} else {
		var checkState int
		if checkDriver.Status == 2 {
			checkState = 1
		}
		var user autocode.MinappUser
		fmt.Println("用户id", checkDriver.Uid, checkState)
		err = global.GVA_DB.Model(&user).Where("id", checkDriver.Uid).Update("check_driver", checkState).Error
		return err
	}

}

// GetCheckDriver 根据id获取CheckDriver记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkDriverService *CheckDriverService) GetCheckDriver(id uint) (err error, checkDriver autocode.CheckDriver) {
	err = global.GVA_DB.Where("id = ?", id).First(&checkDriver).Error
	return
}

// GetCheckDriverInfoList 分页获取CheckDriver记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkDriverService *CheckDriverService) GetCheckDriverInfoList(info autoCodeReq.CheckDriverSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.CheckDriver{})
	var checkDrivers []autocode.CheckDriver
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}

	if info.Time != "" {
		db = db.Where("time = ?", info.Time)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("MinAppUser").Find(&checkDrivers).Error
	return err, checkDrivers, total
}
