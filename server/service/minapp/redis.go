package minapp

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

type RedisService struct{}

var today = utils.Today()
var ctx = context.Background()

//今日新增的用户
func SaveNewUser(userNum string) {
	k := "user:num:" + today
	err := global.GVA_REDIS.SAdd(ctx, k, userNum).Err()
	if err != nil {
		global.GVA_LOG.Error("保存新增用户失败", zap.Error(err))
	}
}

//今日活跃的用户
func SaveActiveUser(uid uint) {
	k := "user:id:" + today
	err := global.GVA_REDIS.SAdd(ctx, k, uid).Err()
	if err != nil {
		global.GVA_LOG.Error("保存活跃用户失败", zap.Error(err))
	}
}

type everydayUserCount struct {
	Date  string `json:"date"`
	Total int64  `json:"total"`
}

type userStatistics struct {
	List   []everydayUserCount `json:"list"`
	Active int64               `json:"active"`
	All    int64               `json:"all"`
}

//获取时间范围内的用户数据统计
func (s *RedisService) GetUserCount(start, end string) (err error, data interface{}) {
	//获取当前时间戳
	timeUnix := time.Now().Unix()
	//判断如果超出当前时期则改变结束时间为当天
	endUnix := utils.DateToTime(end)

	if endUnix > timeUnix {
		end = utils.Today()
	}

	rangeDate := utils.GetBetweenDates(start, end)
	var everyday everydayUserCount
	var list []everydayUserCount

	//添加每日新增的用户到数组
	for _, v := range rangeDate {
		num, _ := global.GVA_REDIS.SCard(ctx, "user:num:"+v).Result()
		everyday.Date = v
		everyday.Total = num
		list = append(list, everyday)
	}

	var users userStatistics
	users.List = list
	//获取今日活跃的用户数量
	if r, err := global.GVA_REDIS.SCard(ctx, "user:id:"+today).Result(); err != nil {
		return errors.New("获取今日活跃用户错误"), users
	} else {
		users.Active = r
	}

	var count int64
	//查询所有用户的总数
	var minUser autocode.MinappUser
	if err := global.GVA_DB.Model(&minUser).Count(&count).Error; err != nil {
		return errors.New("获取用户总数错误"), users
	} else {
		users.All = count
	}

	return err, users
}

func (s *RedisService) GetMonthlyUsers() (err error, data interface{}) {
	t := time.Now()
	year := strconv.Itoa(t.Year())
	rangeDate := utils.GetBetweenDates(year+"-01-01", utils.Today())

	month := [12]int64{}
	//添加每日新增的用户到数组
	for _, v := range rangeDate {
		num, _ := global.GVA_REDIS.SCard(ctx, "user:id:"+v).Result()
		arr := strings.Split(v, "-")
		m := arr[1]
		switch m {
		case "01":
			month[0] += num
		case "02":
			month[1] += num
		case "03":
			month[2] += num
		case "04":
			month[3] += num
		case "05":
			month[4] += num
		case "06":
			month[5] += num
		case "07":
			month[6] += num
		case "08":
			month[7] += num
		case "09":
			month[8] += num
		case "10":
			month[9] += num
		case "11":
			month[10] += num
		case "12":
			month[11] += num

		}
	}

	return err, month
}
