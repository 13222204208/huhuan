package minapp

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/pay"
	payConfig "github.com/silenceper/wechat/v2/pay/config"
	"github.com/silenceper/wechat/v2/pay/notify"
	"go.uber.org/zap"
)

type NotifyApi struct{}

func GetMinAppSet() (minAppSet autocode.MinappSet) {

	err := global.GVA_DB.First(&minAppSet).Error
	if err != nil {
		global.GVA_LOG.Error("获取小程序设置失败, err:", zap.Error(err))
	}
	return minAppSet
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

func (a *NotifyApi) PayNotify(c *gin.Context) {
	PayConfig()
	var pn pay.Pay
	pn.GetNotify()
	var r notify.PaidResult
	_ = c.ShouldBindXML(&r)
	var n notify.Notify
	n = notify.Notify{
		global.GVA_PAY_CONFIG,
	}
	fmt.Println("test", *r.ReturnCode)
	ok := n.PaidVerifySign(r)
	fmt.Println("返回信息", ok)
	if ok {
		if *r.ReturnCode == "SUCCESS" {
			tradeNo := *r.OutTradeNo
			s := tradeNo[:3]
			switch s {
			case "BAL":
				var record minapp.Record
				global.GVA_DB.Where("order_num = ?", tradeNo).First(&record)
				if record.Status == 0 {
					global.GVA_DB.Model(&record).Where("order_num = ?", tradeNo).Update("status", 1)
					//充值余额到用户
					if err := userService.BackgroundTopUpBalance(record.Uid, record.Money, 1); err != nil {
						global.GVA_LOG.Error("增加余额失败", zap.Error(err))
					}
					//金额满足时赠送优惠券
					if err := couponService.TopUpGiveCoupon(record.Uid, record.Money); err != nil {
						global.GVA_LOG.Error(err.Error(), zap.Error(err))
					}
				}
			case "BWM":
				var buy autocode.HelpBuy
				//修改帮我买订单状态
				global.GVA_DB.Model(&buy).Where("order_num = ?", tradeNo).Update("status", 1)
				//修改订单中心发布状态
				if err := userService.UpdateOrderCenterStatus(tradeNo, "create"); err != nil {
					global.GVA_LOG.Error("修改订单状态失败", zap.Error(err))

				}
			case "BWQ":
				var fetch autocode.HelpFetch
				//修改帮我订单状态
				global.GVA_DB.Model(&fetch).Where("order_num = ?", tradeNo).Update("status", 1)
				//修改订单中心发布状态
				if err := userService.UpdateOrderCenterStatus(tradeNo, "create"); err != nil {
					global.GVA_LOG.Error("修改订单状态失败", zap.Error(err))

				}
			case "BWZ":
				var work autocode.HelpWork
				//修改帮我订单状态
				global.GVA_DB.Model(&work).Where("order_num = ?", tradeNo).Update("status", 1)
				//修改订单中心发布状态
				if err := userService.UpdateOrderCenterStatus(tradeNo, "create"); err != nil {
					global.GVA_LOG.Error("修改订单状态失败", zap.Error(err))

				}
			case "BWX":
				var repair autocode.HelpRepair
				//修改帮我订单状态
				global.GVA_DB.Model(&repair).Where("order_num = ?", tradeNo).Update("status", 1)
				//修改订单中心发布状态
				if err := userService.UpdateOrderCenterStatus(tradeNo, "create"); err != nil {
					global.GVA_LOG.Error("修改订单状态失败", zap.Error(err))

				}

			case "REW":
				var reward minapp.Reward
				global.GVA_DB.Where("order_num =?", tradeNo).First(&reward)
				if reward.Status == 0 {
					//修改打赏记录状态
					global.GVA_DB.Model(&minapp.Reward{}).Where("order_num = ?", tradeNo).Update("status", 1)

					//充值余额到用户
					if err := userService.BackgroundTopUpBalance(reward.Tid, reward.Money, 1); err != nil {
						global.GVA_LOG.Error("增加余额失败", zap.Error(err))
					}
				}

			}
		}
		var p notify.PaidResp
		p.ReturnCode = "SUCCESS"
		p.ReturnMsg = "OK"
		c.XML(200, p)
		return
	}

}
