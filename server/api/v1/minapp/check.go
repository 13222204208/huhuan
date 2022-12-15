package minapp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

type CheckApi struct{}

//司机认证
func (a *CheckApi) CreateCheckDriver(c *gin.Context) {
	var d autocode.CheckDriver
	_ = c.ShouldBindJSON(&d)
	if err := utils.Verify(d, utils.CheckDriverVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := c.MustGet("id").(uint)
	d.Uid = uid
	if err := checkDriverService.CreateCheckDriver(d); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("提交成功", c)
	}

}

//商家认证
func (a *CheckApi) CreateCheckMerchant(c *gin.Context) {
	var d autocode.CheckMerchant
	_ = c.ShouldBindJSON(&d)
	if err := utils.Verify(d, utils.CheckMerchantVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := c.MustGet("id").(uint)
	d.Uid = uid
	if err := checkMerchantService.CreateCheckMerchant(d); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("提交成功", c)
	}
}

//身份认证
func (a *CheckApi) CreateCheckIdentity(c *gin.Context) {
	var i autocode.CheckIdentity
	_ = c.ShouldBindJSON(&i)
	if err := utils.Verify(i, utils.CheckIdentifyVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := c.MustGet("id").(uint)

	i.Uid = uid

	if err := checkIdentityService.CreateCheckIdentity(i); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("提交成功", c)
	}
}
