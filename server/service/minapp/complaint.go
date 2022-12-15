package minapp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
)

type ComplaintService struct{}

//用户订单投诉
func (s *ComplaintService) GetAllComplaintLabels() (err error, list interface{}) {
	var labels []autocode.ComplaintLabel
	err = global.GVA_DB.Find(&labels).Error
	return err, labels
}

func (s *ComplaintService) CreateComplaintOrder(complaint autocode.Complaint) (err error) {
	err = global.GVA_DB.Save(&complaint).Error
	return err
}
