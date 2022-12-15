package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

type HelpCarpoolService struct {
}

// CreateHelpCarpool 创建HelpCarpool记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpCarpoolService *HelpCarpoolService) CreateHelpCarpool(helpCarpool autocode.HelpCarpool) (err error) {
	helpCarpool.CreationDate = utils.Today()
	err = global.GVA_DB.Create(&helpCarpool).Error

	//保存发布订单到订单中心
	if err != nil {
		global.GVA_LOG.Error("发布订单错误", zap.Error(err))
	} else {
		CreateOrderCenter(helpCarpool.Uid, "create", helpCarpool.OrderNum, helpCarpool.Price)
	}

	return err
}

// DeleteHelpCarpool 删除HelpCarpool记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpCarpoolService *HelpCarpoolService) DeleteHelpCarpool(helpCarpool autocode.HelpCarpool) (err error) {
	err = global.GVA_DB.Delete(&helpCarpool).Error
	return err
}

// DeleteHelpCarpoolByIds 批量删除HelpCarpool记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpCarpoolService *HelpCarpoolService) DeleteHelpCarpoolByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.HelpCarpool{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHelpCarpool 更新HelpCarpool记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpCarpoolService *HelpCarpoolService) UpdateHelpCarpool(helpCarpool autocode.HelpCarpool) (err error) {
	err = global.GVA_DB.Save(&helpCarpool).Error
	return err
}

// GetHelpCarpool 根据id获取HelpCarpool记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpCarpoolService *HelpCarpoolService) GetHelpCarpool(id uint) (err error, helpCarpool autocode.HelpCarpool) {
	err = global.GVA_DB.Where("id = ?", id).Preload("UserInfo").First(&helpCarpool).Error
	return
}

// GetHelpCarpoolInfoList 分页获取HelpCarpool记录
// Author [piexlmax](https://github.com/piexlmax)
func (helpCarpoolService *HelpCarpoolService) GetHelpCarpoolInfoList(info autoCodeReq.HelpCarpoolSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.HelpCarpool{})
	var helpCarpools []autocode.HelpCarpool
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderNum != "" {
		db = db.Where("order_num LIKE ?", "%"+info.OrderNum+"%")
	}
	if info.ContactPhone != "" {
		db = db.Where("contact_phone = ?", info.ContactPhone)
	}
	if info.TakeName != "" {
		db = db.Where("take_name LIKE ?", "%"+info.TakeName+"%")
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
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.StartAddress != "" {
		db = db.Where("start_address LIKE ?", "%"+info.StartAddress+"%")
	}
	if info.ArriveAddress != "" {
		db = db.Where("arrive_address LIKE ?", "%"+info.ArriveAddress+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("CarUsedNumInfo").Preload("UserInfo").Preload("OnlineMessage").Order("id DESC").Find(&helpCarpools).Error
	return err, helpCarpools, total
}
