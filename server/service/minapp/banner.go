package minapp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
)

type BannerService struct{}

func (s *BannerService) GetBannerList(city string, class int) (err error, list interface{}) {
	var banner []autocode.Banner
	err = global.GVA_DB.Where("city = ? and class = ?", city, class).Find(&banner).Error
	return err, banner
}
