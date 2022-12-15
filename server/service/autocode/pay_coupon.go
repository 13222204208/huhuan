package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type PayCouponService struct {
}

// CreatePayCoupon 创建PayCoupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (payCouponService *PayCouponService) CreatePayCoupon(payCoupon autocode.PayCoupon) (err error) {
	err = global.GVA_DB.Create(&payCoupon).Error
	return err
}

// DeletePayCoupon 删除PayCoupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (payCouponService *PayCouponService)DeletePayCoupon(payCoupon autocode.PayCoupon) (err error) {
	err = global.GVA_DB.Delete(&payCoupon).Error
	return err
}

// DeletePayCouponByIds 批量删除PayCoupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (payCouponService *PayCouponService)DeletePayCouponByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.PayCoupon{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePayCoupon 更新PayCoupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (payCouponService *PayCouponService)UpdatePayCoupon(payCoupon autocode.PayCoupon) (err error) {
	err = global.GVA_DB.Save(&payCoupon).Error
	return err
}

// GetPayCoupon 根据id获取PayCoupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (payCouponService *PayCouponService)GetPayCoupon(id uint) (err error, payCoupon autocode.PayCoupon) {
	err = global.GVA_DB.Where("id = ?", id).First(&payCoupon).Error
	return
}

// GetPayCouponInfoList 分页获取PayCoupon记录
// Author [piexlmax](https://github.com/piexlmax)
func (payCouponService *PayCouponService)GetPayCouponInfoList(info autoCodeReq.PayCouponSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.PayCoupon{})
    var payCoupons []autocode.PayCoupon
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&payCoupons).Error
	return err, payCoupons, total
}
