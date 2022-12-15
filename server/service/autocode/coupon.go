package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CouponService struct {
}

// CreateCoupon 创建Coupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (couponService *CouponService) CreateCoupon(coupon autocode.Coupon) (err error) {
	err = global.GVA_DB.Create(&coupon).Error
	return err
}

// DeleteCoupon 删除Coupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (couponService *CouponService) DeleteCoupon(coupon autocode.Coupon) (err error) {
	err = global.GVA_DB.Delete(&coupon).Error
	return err
}

// DeleteCouponByIds 批量删除Coupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (couponService *CouponService) DeleteCouponByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Coupon{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCoupon 更新Coupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (couponService *CouponService) UpdateCoupon(coupon autocode.Coupon) (err error) {
	err = global.GVA_DB.Save(&coupon).Error
	return err
}

// GetCoupon 根据id获取Coupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (couponService *CouponService) GetCoupon(id uint) (err error, coupon autocode.Coupon) {
	err = global.GVA_DB.Where("id = ?", id).First(&coupon).Error
	return
}

// GetCouponInfoList 分页获取Coupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (couponService *CouponService) GetCouponInfoList(info autoCodeReq.CouponSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.Coupon{})
	var coupons []autocode.Coupon
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&coupons).Error
	return err, coupons, total
}

func (couponService *CouponService) GetAllCouponsTitle() (err error, list interface{}) {
	var coupon []autocode.Coupon
	err = global.GVA_DB.Find(&coupon).Error
	return err, coupon
}
