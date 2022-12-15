package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
)

type Login struct {
	Code     string `json:"code"`     //小程序Code
	Avatar   string `json:"avatar"`   //头像
	Nickname string `json:"nickname"` //昵称
}

type GetPhone struct {
	Code          string `json:"code"`
	EncryptedData string `json:"encryptedData"`
	Iv            string `json:"iv"`
}

type UpdatePayPassWord struct {
	Password string `json:"password"`
}

type SendEmail struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

//充值余额
type RechargeBalance struct {
	Balance float64 `json:"balance" binding:"required"`
}

//用户钱包记录
type UserMoneyRecord struct {
	minapp.Record
	request.PageInfo
}

//余额提现
type BalanceWithdraw struct {
	Money float64 `json:"money"`
}

//验证支付密码
type VerifyUserPassword struct {
	Password string `json:"password"`
}

//订单重新支付
type AgainPayOrder struct {
	OrderNum string  `json:"orderNum"`
	Money    float64 `json:"money"`
}
