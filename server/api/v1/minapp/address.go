package minapp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	minAppRequest "github.com/flipped-aurora/gin-vue-admin/server/model/minapp/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AddressApi struct{}

func (a *AddressApi) CreateAddress(c *gin.Context) {
	var add minAppRequest.CreateAddress
	_ = c.ShouldBindJSON(&add)
	uid := c.MustGet("id").(uint)

	if err := utils.Verify(add, utils.MinAppAddressVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	address := &minapp.Address{
		Name:     add.Name,
		Phone:    add.Phone,
		Country:  add.Country,
		Province: add.Province,
		City:     add.City,
		Area:     add.Area,
		Detail:   add.Detail,
		Uid:      uid,
		PostCode: add.PostCode,
	}

	if err := addressService.Create(*address); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}

//修改地址
func (a *AddressApi) UpdateAddress(c *gin.Context) {
	var u minAppRequest.UpdateAddress
	_ = c.ShouldBindJSON(&u)
	uid := c.MustGet("id").(uint)
	if u.ID < 1 {
		response.FailWithMessage("地址ID不能为空", c)
	}
	address := &minapp.Address{
		Name:     u.Name,
		Phone:    u.Phone,
		Country:  u.Country,
		Province: u.Province,
		City:     u.City,
		Area:     u.Area,
		Detail:   u.Detail,
		PostCode: u.PostCode,
	}

	if err := addressService.Update(*address, uid, u.ID); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithMessage("成功", c)
	}
}

func (a *AddressApi) DeleteAddress(c *gin.Context) {
	var u minAppRequest.DeleteAddress
	_ = c.ShouldBindJSON(&u)
	uid := c.MustGet("id").(uint)
	if u.ID < 1 {
		response.FailWithMessage("地址ID不能为空", c)
	}

	if err := addressService.Delete(u.ID, uid); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithMessage("成功", c)
	}
}

//获取用户的地址
func (a *AddressApi) ListAddress(c *gin.Context) {
	uid := c.MustGet("id").(uint)
	if err, result := addressService.MyAddress(uid); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithDetailed(gin.H{"addressList": result}, "成功", c)
	}
}
