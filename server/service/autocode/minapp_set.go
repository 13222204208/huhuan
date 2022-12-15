package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type MinappSetService struct {
}

// CreateMinappSet 创建MinappSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (minappSetService *MinappSetService) CreateMinappSet(minappSet autocode.MinappSet) (err error) {
	err = global.GVA_DB.Create(&minappSet).Error
	return err
}

// DeleteMinappSet 删除MinappSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (minappSetService *MinappSetService)DeleteMinappSet(minappSet autocode.MinappSet) (err error) {
	err = global.GVA_DB.Delete(&minappSet).Error
	return err
}

// DeleteMinappSetByIds 批量删除MinappSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (minappSetService *MinappSetService)DeleteMinappSetByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.MinappSet{},"id in ?",ids.Ids).Error
	return err
}

// UpdateMinappSet 更新MinappSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (minappSetService *MinappSetService)UpdateMinappSet(minappSet autocode.MinappSet) (err error) {
	err = global.GVA_DB.Save(&minappSet).Error
	return err
}

// GetMinappSet 根据id获取MinappSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (minappSetService *MinappSetService)GetMinappSet(id uint) (err error, minappSet autocode.MinappSet) {
	err = global.GVA_DB.Where("id = ?", id).First(&minappSet).Error
	return
}

// GetMinappSetInfoList 分页获取MinappSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (minappSetService *MinappSetService)GetMinappSetInfoList(info autoCodeReq.MinappSetSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.MinappSet{})
    var minappSets []autocode.MinappSet
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&minappSets).Error
	return err, minappSets, total
}
