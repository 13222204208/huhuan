package minapp

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	payConfig "github.com/silenceper/wechat/v2/pay/config"
	"github.com/silenceper/wechat/v2/pay/order"
	"gorm.io/gorm"
)

type PayService struct{}

//获取加元与人民币的币的汇率
func GetMinAppSet() (minAppSet autocode.MinappSet) {

	err := global.GVA_DB.First(&minAppSet).Error
	if err != nil {
		minAppSet.ExchangeRate = 5.0672
	}
	return minAppSet
}
func GetUserInfo(uid uint) (err error, user autocode.MinappUser) {
	var reqUser autocode.MinappUser
	err = global.GVA_DB.First(&reqUser, "id = ?", uid).Error
	return err, reqUser
}

//余额支付的公共方法
//订单号，支付金额，支付的用户ID
func (s *PayService) BalancePayMethod(orderNum string, money float64, uid uint) (err error) {
	//加元转换成人民币
	//exchangeRate := GetMinAppSet()
	//money = money * exchangeRate.ExchangeRate //汇率相乘
	money = utils.Decimal(money)
	if err, userinfo := GetUserInfo(uid); err != nil {
		return errors.New("找不到用户")
	} else {
		balance := userinfo.Balance
		fmt.Println("余额", balance)
		if balance < money {
			return errors.New("余额不足")
		} else {
			if err := reduceBalance(uid, money); err != nil {
				return errors.New("扣除余额失败")
			} else {
				//保存使用余额支付订单的流水记录
				BalancePayRecord(uid, orderNum, money)
				return nil
			}
		}
	}
}

//扣除用户余额
func reduceBalance(uid uint, money float64) (err error) {
	var user autocode.MinappUser
	err = global.GVA_DB.Model(&user).Where("id=?", uid).UpdateColumn("balance", gorm.Expr("balance - ?", money)).Error
	return err
}

//保存余额支付订单的流水记录
func BalancePayRecord(uid uint, orderNum string, money float64) {
	var record minapp.Record
	record.Uid = uid
	record.OrderNum = orderNum
	record.Money = money
	record.Type = 3
	record.Status = 1
	record.Remark = "用余额支付的订单号为" + orderNum
	global.GVA_DB.Save(&record)

}

//支付配置
func PayConfig() {
	r := GetMinAppSet()

	pc := &payConfig.Config{
		AppID:     r.Appid,
		MchID:     r.Mch_id,
		Key:       r.Pay_secret,
		NotifyURL: r.NotifyUrl,
	}
	global.GVA_PAY_CONFIG = pc
}

//微信支付的公共方法
func (s *PayService) UserPay(openid string, createIP string, money float64, orderNum string) (err error, config order.Config) {
	PayConfig()
	var params order.Params
	//加元转换成人民币
	//exchangeRate := GetMinAppSet()
	//money = money * GetRate() //exchangeRate.ExchangeRate //汇率相乘
	money = utils.Decimal(money) * 100
	totalFee := utils.FloatToString(money)

	params.TotalFee = totalFee
	params.CreateIP = createIP
	params.Body = "微信支付"
	params.OutTradeNo = orderNum
	params.OpenID = openid
	params.TradeType = "JSAPI"
	params.NotifyURL = global.GVA_PAY_CONFIG.NotifyURL

	var sendOrder order.Order
	sendOrder = order.Order{
		global.GVA_PAY_CONFIG,
	}

	config, err = sendOrder.BridgeConfig(&params)
	return err, config
}

//重新支付订单
func (s *PayService) AgainPayOrder(openid string, createIP string, money float64, orderNum string) (err error, config order.Config) {
	PayConfig()
	var params order.Params
	//加元转换成人民币
	//exchangeRate := GetMinAppSet()
	//money = money * GetRate() //exchangeRate.ExchangeRate //汇率相乘
	money = utils.Decimal(money) * 100
	totalFee := utils.FloatToString(money)

	params.TotalFee = totalFee
	params.CreateIP = createIP
	params.Body = "微信支付"
	params.OutTradeNo = orderNum
	params.OpenID = openid
	params.TradeType = "JSAPI"
	params.NotifyURL = global.GVA_PAY_CONFIG.NotifyURL

	var sendOrder order.Order
	sendOrder = order.Order{
		global.GVA_PAY_CONFIG,
	}

	config, err = sendOrder.BridgeConfig(&params)
	return err, config
}

//优惠券使用的公共方法
//uid 用户id  cid 优惠券的id
//discounts优惠的金额
func (s *PayService) UserCoupon(uid uint, cid int, money float64) (err error, discounts float64) {
	var releaseRecord autocode.ReleaseRecord
	today := utils.Today()
	if errors.Is(global.GVA_DB.Where("id = ? AND uid = ? AND status = 0 ", cid, uid).Where(" start <= ? AND over >= ?", today, today).First(&releaseRecord).Error, gorm.ErrRecordNotFound) {
		return errors.New("此优惠券无法使用"), 0
	}
	if money < releaseRecord.Condition || money < releaseRecord.Way {
		return errors.New("不满足使用条件"), 0
	}
	return err, releaseRecord.Way
}

//输入赏金时触发，满足条件时 自动选择优惠券
func (s *PayService) AutoSelectCoupon(uid uint, money float64) (err error, r autocode.ReleaseRecord) {
	today := utils.Today()
	if errors.Is(global.GVA_DB.Where("uid = ? AND status = 0 AND  condition <= ?", uid, money).Where(" start <= ? AND over >= ?", today, today).First(&r).Error, gorm.ErrRecordNotFound) {
		return errors.New("没有符合条件的优惠券"), r
	} else {
		return err, r
	}
}
