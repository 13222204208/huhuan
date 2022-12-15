package minapp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
)

type AddressService struct{}

func (s *AddressService) Create(address minapp.Address) (err error) {
	err = global.GVA_DB.Create(&address).Error
	return err
}

func (s *AddressService) Update(a minapp.Address, uid uint, id uint) (err error) {
	var address minapp.Address
	err = global.GVA_DB.Where("id = ? and uid = ?", id, uid).First(&address).Updates(&a).Error
	return err
}

func (s *AddressService) Delete(id uint, uid uint) error {
	var address minapp.Address
	err := global.GVA_DB.Where("id = ? and uid = ? ", id, uid).Delete(&address).Error
	return err
}

func (s *AddressService) MyAddress(uid uint) (err error, list interface{}) {
	var address []minapp.Address
	err = global.GVA_DB.Where("uid = ? ", uid).Find(&address).Error
	return err, address
}
