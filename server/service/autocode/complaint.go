package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ComplaintService struct {
}

// CreateComplaint 创建Complaint记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintService *ComplaintService) CreateComplaint(complaint autocode.Complaint) (err error) {
	err = global.GVA_DB.Create(&complaint).Error
	return err
}

// DeleteComplaint 删除Complaint记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintService *ComplaintService) DeleteComplaint(complaint autocode.Complaint) (err error) {
	err = global.GVA_DB.Delete(&complaint).Error
	return err
}

// DeleteComplaintByIds 批量删除Complaint记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintService *ComplaintService) DeleteComplaintByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Complaint{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateComplaint 更新Complaint记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintService *ComplaintService) UpdateComplaint(complaint autocode.Complaint) (err error) {
	err = global.GVA_DB.Save(&complaint).Error
	return err
}

// GetComplaint 根据id获取Complaint记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintService *ComplaintService) GetComplaint(id uint) (err error, complaint autocode.Complaint) {
	err = global.GVA_DB.Where("id = ?", id).First(&complaint).Error
	return
}

// GetComplaintInfoList 分页获取Complaint记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintService *ComplaintService) GetComplaintInfoList(info autoCodeReq.ComplaintSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.Complaint{})
	var complaints []autocode.Complaint
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.OrderNum != "" {
		db = db.Where("order_num LIKE ?", "%"+info.OrderNum+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Preload("ComplainantUserInfo").Preload("ByComplainantUserInfo").Offset(offset).Find(&complaints).Error
	return err, complaints, total
}
