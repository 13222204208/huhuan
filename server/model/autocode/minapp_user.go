// 自动生成模板MinappUser
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 如果含有time.Time 请自行import time包
type MinappUser struct {
	global.GVA_MODEL
	Openid          string  `json:"openid" form:"openid" gorm:"column:openid;comment:微信openid;"`
	UserNum         string  `json:"userNum" form:"userNum" gorm:"unique;column:user_num;comment:用户编号;"`
	Avatar          string  `json:"avatar" form:"avatar" gorm:"column:avatar;comment:微信小程序头像;"`
	Nickname        string  `json:"nickname" form:"nickname" gorm:"column:nickname;comment:小程序昵称;"`
	Password        string  `json:"password" form:"password" gorm:"column:password;comment:用户支付密码;size:255"`
	Name            string  `json:"name" form:"name" gorm:"column:name;comment:用户姓名"`
	Gender          string  `json:"gender" form:"gender" gorm:"column:gender;comment:用户姓别"`
	Mail            string  `json:"mail" form:"mail" gorm:"column:mail"`
	Phone           string  `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:11;"`
	Area            string  `json:"area" form:"area" gorm:"column:area;comment:地区;"`
	Number          string  `json:"number" form:"number" gorm:"column:number;comment:证件号;"`
	Balance         float64 `json:"balance" form:"balance" gorm:"column:balance;comment:余额"`
	Integral        uint    `json:"integral" form:"integral" gorm:"column:integral;comment:用户积分;default:0"`
	SignInDay       int     `json:"signInDay" form:"signInDay" gorm:"column:sign_in_day;comment:签到天数;default:0"`
	SignInState     uint8   `json:"signInState" form:"signInState" gorm:"column:sign_in_state;comment:签到开关;default:0"`
	CancelNumber    uint    `json:"cancelNumber" form:"cancelNumber" gorm:"column:cancel_number;default:0;comment:取消订单次数"`
	AllOrderNumber  uint    `json:"allOrderNumber" form:"allOrderNumber" gorm:"column:all_order_number;default:0;comment:订单量包括发单和接单，二手闲置和团购接龙除外"`
	TakeOrderNumber uint    `json:"takeOrderNumber" form:"takeOrderNumber" gorm:"column:take_order_number;default:0;comment:接单量 二手闲置和团购接龙除外"`
	ComplaintNumber uint    `json:"complaintNumber" form:"complaintNumber" gorm:"column:complaint_number;default:0;comment:被投诉的订单量"`
	Verify          int     `json:"verify" form:"verify" gorm:"column:verify;comment:接单必须身份为验证;default:0"`
	CheckDriver     int     `json:"checkDriver" form:"checkDriver" gorm:"column:check_driver;comment:司机认证的状态，为1时认证通过;default:0"`
	CheckMerchant   int     `json:"checkMerchant" form:"checkMerchant" gorm:"check_merchant;comment:商家认证的状态，为1 时认证通过;default:0"`
	Status          *int    `json:"status" form:"status" gorm:"column:status;comment:0 禁用 1启用;size:1;"`
	CouponTotal     int64   `json:"couponTotal" gorm:"-"`
}
