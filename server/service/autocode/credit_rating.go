package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type CreditRatingService struct {
}

// CreateCreditRating 创建CreditRating记录
// Author [piexlmax](https://github.com/piexlmax)
func (creditRatingService *CreditRatingService) CreateCreditRating(creditRating autocode.CreditRating) (err error) {
	err = global.GVA_DB.Create(&creditRating).Error
	return err
}

// DeleteCreditRating 删除CreditRating记录
// Author [piexlmax](https://github.com/piexlmax)
func (creditRatingService *CreditRatingService)DeleteCreditRating(creditRating autocode.CreditRating) (err error) {
	err = global.GVA_DB.Delete(&creditRating).Error
	return err
}

// DeleteCreditRatingByIds 批量删除CreditRating记录
// Author [piexlmax](https://github.com/piexlmax)
func (creditRatingService *CreditRatingService)DeleteCreditRatingByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CreditRating{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCreditRating 更新CreditRating记录
// Author [piexlmax](https://github.com/piexlmax)
func (creditRatingService *CreditRatingService)UpdateCreditRating(creditRating autocode.CreditRating) (err error) {
	err = global.GVA_DB.Save(&creditRating).Error
	return err
}

// GetCreditRating 根据id获取CreditRating记录
// Author [piexlmax](https://github.com/piexlmax)
func (creditRatingService *CreditRatingService)GetCreditRating(id uint) (err error, creditRating autocode.CreditRating) {
	err = global.GVA_DB.Where("id = ?", id).First(&creditRating).Error
	return
}

// GetCreditRatingInfoList 分页获取CreditRating记录
// Author [piexlmax](https://github.com/piexlmax)
func (creditRatingService *CreditRatingService)GetCreditRatingInfoList(info autoCodeReq.CreditRatingSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.CreditRating{})
    var creditRatings []autocode.CreditRating
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&creditRatings).Error
	return err, creditRatings, total
}
