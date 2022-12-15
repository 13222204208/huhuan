// 自动生成模板MinappSet
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MinappSet 结构体
// 如果含有time.Time 请自行import time包
type MinappSet struct {
	global.GVA_MODEL
	Appid           string  `json:"appid" form:"appid" gorm:"column:appid;comment:小程序appid;"`
	App_secret      string  `json:"app_secret" form:"app_secret" gorm:"column:app_secret;comment:小程序秘钥;"`
	Mch_id          string  `json:"mch_id" form:"mch_id" gorm:"column:mch_id;comment:微信商户号;"`
	Pay_secret      string  `json:"pay_secret" form:"pay_secret" gorm:"column:pay_secret;comment:微信支付秘钥;"`
	ExchangeRate    float64 `json:"exchangeRate" form:"exchangeRate" gorm:"column:exchange_rate;comment:加元与人民币的汇率"`
	NotifyUrl       string  `json:"notifyUrl" form:"notifyUrl" gorm:"column:notify_url;comment:支付回调地址"`
	Integral        uint    `json:"integral" form:"integral" gorm:"column:integral;comment:签到每天送的积分"`
	CancelNumber    uint    `json:"cancelNumber" form:"cancelNumber" gorm:"column:cancel_number;default:0;comment:取消订单次数"`
	AllOrderNumber  uint    `json:"allOrderNumber" form:"allOrderNumber" gorm:"column:all_order_number;default:0;comment:订单量包括发单和接单，二手闲置和团购接龙除外"`
	TakeOrderNumber uint    `json:"takeOrderNumber" form:"takeOrderNumber" gorm:"column:take_order_number;default:0;comment:接单量 二手闲置和团购接龙除外"`
	ComplaintNumber uint    `json:"complaintNumber" form:"complaintNumber" gorm:"column:complaint_number;default:0;comment:被投诉的订单量"`
}
