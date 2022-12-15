package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MinappUserService struct {
}

// CreateMinappUser 创建MinappUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (minappUserService *MinappUserService) CreateMinappUser(minappUser autocode.MinappUser) (err error) {
	err = global.GVA_DB.Create(&minappUser).Error
	return err
}

// DeleteMinappUser 删除MinappUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (minappUserService *MinappUserService) DeleteMinappUser(minappUser autocode.MinappUser) (err error) {
	err = global.GVA_DB.Delete(&minappUser).Error
	return err
}

// DeleteMinappUserByIds 批量删除MinappUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (minappUserService *MinappUserService) DeleteMinappUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.MinappUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMinappUser 更新MinappUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (minappUserService *MinappUserService) UpdateMinappUser(minappUser autocode.MinappUser) (err error) {
	err = global.GVA_DB.Save(&minappUser).Error
	return err
}

// GetMinappUser 根据id获取MinappUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (minappUserService *MinappUserService) GetMinappUser(id uint) (err error, minappUser autocode.MinappUser) {
	err = global.GVA_DB.Where("id = ?", id).First(&minappUser).Error
	return
}

// GetMinappUserInfoList 分页获取MinappUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (minappUserService *MinappUserService) GetMinappUserInfoList(info autoCodeReq.MinappUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.MinappUser{})
	var minappUsers []autocode.MinappUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+info.Nickname+"%")
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.Area != "" {
		db = db.Where("area = ?", info.Area)
	}
	if info.Number != "" {
		db = db.Where("number LIKE ?", "%"+info.Number+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&minappUsers).Error
	return err, minappUsers, total
}

//更改签到提醒状态
func (minappUserService *MinappUserService) ChangeSignInState(uid uint, state uint8) (err error) {
	db := global.GVA_DB.Model(&autocode.MinappUser{})

	err = db.Where("id = ?", uid).Update("sign_in_state", state).Error
	return err
}

type HelpList struct {
	Uid   uint    `json:"uid"`
	Area  string  `json:"area"`
	Total int     `json:"total"`
	Money float64 `json:"money"`
}

type AllUserTop struct {
	BuyList    []HelpList `json:"buyList"`
	WorkList   []HelpList `json:"workList"`
	RepairList []HelpList `json:"repairList"`
	FetchList  []HelpList `json:"fetchList"`
}

//用户排行 ，用户发布订单的排行
func (minappUserService *MinappUserService) UserTop() (err error, list interface{}) {
	var userTop AllUserTop
	var helpList []HelpList
	//帮我买
	err = global.GVA_DB.Raw("SELECT uid,area,COUNT(id) total,SUM(price) money FROM help_buys GROUP BY uid ORDER BY total DESC").Scan(&helpList).Error
	userTop.BuyList = helpList

	//帮我做
	err = global.GVA_DB.Raw("SELECT uid,area,COUNT(id) total,SUM(price) money FROM help_works GROUP BY uid ORDER BY total DESC").Scan(&helpList).Error
	userTop.WorkList = helpList

	//帮我修
	err = global.GVA_DB.Raw("SELECT uid,area,COUNT(id) total,SUM(price) money FROM help_repairs GROUP BY uid ORDER BY total DESC").Scan(&helpList).Error
	userTop.RepairList = helpList

	//帮我取
	err = global.GVA_DB.Raw("SELECT uid,area,COUNT(id) total,SUM(price) money FROM help_fetches GROUP BY uid ORDER BY total DESC").Scan(&helpList).Error
	userTop.FetchList = helpList
	return err, userTop
}
