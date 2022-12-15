package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AreaService struct {
}

// CreateArea 创建Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService) CreateArea(area autocode.Area) (err error) {
	err = global.GVA_DB.Create(&area).Error
	return err
}

// DeleteArea 删除Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService) DeleteArea(area autocode.Area) (err error) {
	err = global.GVA_DB.Delete(&area).Error
	return err
}

// DeleteAreaByIds 批量删除Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService) DeleteAreaByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Area{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateArea 更新Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService) UpdateArea(area autocode.Area) (err error) {
	err = global.GVA_DB.Save(&area).Error
	return err
}

// GetArea 根据id获取Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService) GetArea(id uint) (err error, area autocode.Area) {
	err = global.GVA_DB.Where("id = ?", id).First(&area).Error
	return
}

// GetAreaInfoList 分页获取Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService) GetAreaInfoList(info autoCodeReq.AreaSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.Area{})
	var areas []autocode.Area
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&areas).Error
	return err, areas, total
}
