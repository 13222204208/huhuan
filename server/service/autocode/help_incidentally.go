package autocode

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

type HelpIncidentallyService struct {
}

// CreateHelpIncidentally 创建HelpIncidentally记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpIncidentallyService *HelpIncidentallyService) CreateHelpIncidentally(helpIncidentally autocode.HelpIncidentally) (err error) {
	err = global.GVA_DB.Create(&helpIncidentally).Error

	//保存发布订单到订单中心
	if err != nil {
		global.GVA_LOG.Error("发布订单错误", zap.Error(err))
	} else {
		CreateOrderCenter(helpIncidentally.Uid, "create", helpIncidentally.OrderNum, helpIncidentally.Price)
	}

	return err
}

// DeleteHelpIncidentally 删除HelpIncidentally记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpIncidentallyService *HelpIncidentallyService) DeleteHelpIncidentally(helpIncidentally autocode.HelpIncidentally) (err error) {
	err = global.GVA_DB.Delete(&helpIncidentally).Error
	return err
}

// DeleteHelpIncidentallyByIds 批量删除HelpIncidentally记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpIncidentallyService *HelpIncidentallyService) DeleteHelpIncidentallyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.HelpIncidentally{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHelpIncidentally 更新HelpIncidentally记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpIncidentallyService *HelpIncidentallyService) UpdateHelpIncidentally(helpIncidentally autocode.HelpIncidentally) (err error) {
	err = global.GVA_DB.Save(&helpIncidentally).Error
	return err
}

// GetHelpIncidentally 根据id获取HelpIncidentally记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpIncidentallyService *HelpIncidentallyService) GetHelpIncidentally(id uint) (err error, helpIncidentally autocode.HelpIncidentally) {
	err = global.GVA_DB.Where("id = ?", id).Preload("UserInfo").First(&helpIncidentally).Error
	return
}

// GetHelpIncidentallyInfoList 分页获取HelpIncidentally记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpIncidentallyService *HelpIncidentallyService) GetHelpIncidentallyInfoList(info autoCodeReq.HelpIncidentallySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	fmt.Println(info, "测试")
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpIncidentally{})
	var helpIncidentallys []autocode.HelpIncidentally
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderNum != "" {
		db = db.Where("order_num LIKE ?", "%"+info.OrderNum+"%")
	}
	if info.Linkman != "" {
		db = db.Where("linkman LIKE ?", "%"+info.Linkman+"%")
	}
	if info.ContactPhone != "" {
		db = db.Where("contact_phone = ?", info.ContactPhone)
	}
	if info.TakeName != "" {
		db = db.Where("take_name LIKE ?", "%"+info.TakeName+"%")
	}

	if info.City != "" {
		db = db.Where("city = ?", info.City)
	}

	if info.UserNum != "" {
		db = db.Where("user_num LIKE ?", "%"+info.UserNum+"%")
	}
	if info.Status > 0 {
		if info.Status == 5 {
			db = db.Where("time < ?", utils.Today())
		} else {
			db = db.Where("status = ?", info.Status)
		}
	}
	if info.Area != "" {
		db = db.Where("area LIKE ?", "%"+info.Area+"%")
	}
	if info.Type > 0 {
		db = db.Where("type = ?", info.Type)
	}
	if info.StartCity != "" {
		db = db.Where("start_city LIKE ?", "%"+info.StartCity+"%")
	}
	if info.ArriveCity != "" {
		db = db.Where("arrive_city LIKE ?", "%"+info.ArriveCity+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("UserInfo").Preload("OnlineMessage").Order("id DESC").Find(&helpIncidentallys).Error
	return err, helpIncidentallys, total
}

//保存信息到订单中心
func CreateOrderCenter(uid uint, way, orderNum string, money float64) {
	var o autocode.OrderCenter
	o.Uid = uid
	o.Way = way
	o.OrderNum = orderNum
	o.Money = money
	err := global.GVA_DB.Save(&o).Error
	if err != nil {
		global.GVA_LOG.Error("保存到订单中心失败", zap.Error(err))
	}
}
