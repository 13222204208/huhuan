package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BannerService struct {
}

// CreateBanner 创建Banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannerService *BannerService) CreateBanner(banner autocode.Banner) (err error) {
	err = global.GVA_DB.Create(&banner).Error
	return err
}

// DeleteBanner 删除Banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannerService *BannerService) DeleteBanner(banner autocode.Banner) (err error) {
	err = global.GVA_DB.Delete(&banner).Error
	return err
}

// DeleteBannerByIds 批量删除Banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannerService *BannerService) DeleteBannerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Banner{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBanner 更新Banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannerService *BannerService) UpdateBanner(banner autocode.Banner) (err error) {
	err = global.GVA_DB.Save(&banner).Error
	return err
}

// GetBanner 根据id获取Banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannerService *BannerService) GetBanner(id uint) (err error, banner autocode.Banner) {
	err = global.GVA_DB.Where("id = ?", id).First(&banner).Error
	return
}

// GetBannerInfoList 分页获取Banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannerService *BannerService) GetBannerInfoList(info autoCodeReq.BannerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.Banner{})
	var banners []autocode.Banner
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.City != "" {
		db = db.Where("city = ?", info.City)
	}
	if info.Class != 0 {
		db = db.Where("class = ?", info.Class)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&banners).Error
	return err, banners, total
}
