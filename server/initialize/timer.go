package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func Timer() {
	if global.GVA_CONFIG.Timer.Start {
		for i := range global.GVA_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.GVA_Timer.AddTaskByFunc("ClearDB", global.GVA_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.GVA_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(global.GVA_CONFIG.Timer.Detail[i])
		}
	}
}

func Task() {
	// 新建一个定时任务对象
	// 根据cron表达式进行时间调度，cron可以精确到秒，大部分表达式格式也是从秒开始。
	//crontab := cron.New()  默认从分开始进行时间调度
	crontab := cron.New(cron.WithSeconds()) //精确到秒
	//定义定时器调用的任务函数
	task := func() {
		nowTime := utils.TodayTime()
		var h autocode.HelpWork
		global.GVA_DB.Where("is_found = ? AND status = ? AND stop_time < ?", 0, 1, nowTime).First(&h)
		if h.Uid > 0 {
			pastDueOrder(h.Uid, h.Price)
			global.GVA_DB.Model(&h).Where("uid = ?", h.Uid).Updates(map[string]interface{}{"status": 5, "is_found": 1})
		}

		var f autocode.HelpFetch
		global.GVA_DB.Where("is_found = ? AND status = ? AND stop_time < ?", 0, 1, nowTime).First(&f)
		if f.Uid > 0 {
			pastDueOrder(f.Uid, f.Price)
			global.GVA_DB.Model(&f).Where("uid = ?", f.Uid).Updates(map[string]interface{}{"status": 5, "is_found": 1})
		}

		var r autocode.HelpRepair
		global.GVA_DB.Where("is_found = ? AND status = ? AND stop_time < ?", 0, 1, nowTime).First(&r)
		if r.Uid > 0 {
			pastDueOrder(r.Uid, r.Price)
			global.GVA_DB.Model(&r).Where("uid = ?", r.Uid).Updates(map[string]interface{}{"status": 5, "is_found": 1})
		}
	}
	//定时任务
	spec := "0 */5 * * * ?" //cron表达式，每5分钟一次
	// 添加定时任务,
	crontab.AddFunc(spec, task)
	// 启动定时器
	crontab.Start()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	//select {} //阻塞主线程停止
}

func pastDueOrder(uid uint, money float64) {
	var r minapp.Record
	r.Uid = uid
	r.Money = money
	r.Type = 3
	r.Status = 1
	r.Remark = "过期订单退款到余额"
	err := global.GVA_DB.Save(&r).Error
	fmt.Println("过期订单退款到余额状态", err)
	if err == nil {
		AddUserBalance(uid, money)
	}
}

//增加用户余额
func AddUserBalance(uid uint, money float64) {
	var user autocode.MinappUser
	global.GVA_DB.Model(&user).Where("id", uid).UpdateColumn("balance", gorm.Expr("balance + ?", money))
}
