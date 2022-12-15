package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type ComplaintLabelService struct {
}

// CreateComplaintLabel 创建ComplaintLabel记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintLabelService *ComplaintLabelService) CreateComplaintLabel(complaintLabel autocode.ComplaintLabel) (err error) {
	err = global.GVA_DB.Create(&complaintLabel).Error
	return err
}

// DeleteComplaintLabel 删除ComplaintLabel记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintLabelService *ComplaintLabelService)DeleteComplaintLabel(complaintLabel autocode.ComplaintLabel) (err error) {
	err = global.GVA_DB.Delete(&complaintLabel).Error
	return err
}

// DeleteComplaintLabelByIds 批量删除ComplaintLabel记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintLabelService *ComplaintLabelService)DeleteComplaintLabelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.ComplaintLabel{},"id in ?",ids.Ids).Error
	return err
}

// UpdateComplaintLabel 更新ComplaintLabel记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintLabelService *ComplaintLabelService)UpdateComplaintLabel(complaintLabel autocode.ComplaintLabel) (err error) {
	err = global.GVA_DB.Save(&complaintLabel).Error
	return err
}

// GetComplaintLabel 根据id获取ComplaintLabel记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintLabelService *ComplaintLabelService)GetComplaintLabel(id uint) (err error, complaintLabel autocode.ComplaintLabel) {
	err = global.GVA_DB.Where("id = ?", id).First(&complaintLabel).Error
	return
}

// GetComplaintLabelInfoList 分页获取ComplaintLabel记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintLabelService *ComplaintLabelService)GetComplaintLabelInfoList(info autoCodeReq.ComplaintLabelSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.ComplaintLabel{})
    var complaintLabels []autocode.ComplaintLabel
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&complaintLabels).Error
	return err, complaintLabels, total
}
