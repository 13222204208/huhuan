package minapp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	minAppModel "github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	payConfig "github.com/silenceper/wechat/v2/pay/config"
	"github.com/silenceper/wechat/v2/pay/order"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type UserService struct{}

//查询后台微信小程序 的设置
func (userService *UserService) GetMinAppSet() (err error, minAppSet autocode.MinappSet) {
	err = global.GVA_DB.First(&minAppSet).Error
	return err, minAppSet
}

func (userService *UserService) GetPhone(code, encryptedData, iv string) (error, string) {
	_, MinAppSet := userService.GetMinAppSet()
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     MinAppSet.Appid,
		AppSecret: MinAppSet.App_secret,
		Cache:     memory,
	}
	mini := wc.GetMiniProgram(cfg)
	a := mini.GetEncryptor()

	c := mini.GetAuth()
	s, errs := c.Code2Session(code)

	sessionKey := ""
	if errs != nil {
		fmt.Println("获取session失败")
	} else {
		sessionKey = s.SessionKey
	}
	r, err := a.Decrypt(sessionKey, encryptedData, iv)
	if err != nil {
		return errors.New("获取手机号失败"), ""
	} else {
		return err, r.PhoneNumber
	}

}

func (userService *UserService) GetOpenid(code string) (err error, openid string) {
	_, MinAppSet := userService.GetMinAppSet()
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     MinAppSet.Appid,
		AppSecret: MinAppSet.App_secret,
		Cache:     memory,
	}
	mini := wc.GetMiniProgram(cfg)
	a := mini.GetAuth()
	r, errs := a.Code2Session(code)

	if errs != nil {
		return
	} else {
		return errs, r.OpenID
	}
}

//小程序用户登录及保存用户
func (userService *UserService) Login(openid, avatar, nickname string) (err error, minAppUser autocode.MinappUser) {
	var MinAppUser autocode.MinappUser
	if !errors.Is(global.GVA_DB.Where("openid = ?", openid).First(&MinAppUser).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		global.GVA_DB.Model(&minAppUser).Updates(map[string]interface{}{"avatar": avatar, "nickname": nickname})
		return errors.New("用户已注册"), MinAppUser
	}

	MinAppUser.Openid = openid
	MinAppUser.Avatar = avatar
	MinAppUser.Nickname = nickname
	MinAppUser.UserNum = utils.UserNum()
	err = global.GVA_DB.Create(&MinAppUser).Error
	if err == nil {
		SaveNewUser(MinAppUser.UserNum) //保存每天新增的用户到集合中
	}
	return err, MinAppUser
}

func (userService *UserService) GetUserInfo(uid uint) (err error, user autocode.MinappUser) {
	var reqUser autocode.MinappUser
	SaveActiveUser(uid) //保存活跃用户id
	err = global.GVA_DB.First(&reqUser, "id = ?", uid).Error

	//判断用户的优惠券数量
	if err == nil {
		var count int64
		var coupon autocode.ReleaseRecord
		global.GVA_DB.Model(&coupon).Where("uid = ? AND status != 1 ", uid).Count(&count)
		reqUser.CouponTotal = count
	}
	return err, reqUser
}

//获取用户协议
func (userService *UserService) GetUserAgreement() (err error, agreement autocode.Agreement) {
	var reqAgreement autocode.Agreement
	err = global.GVA_DB.Last(&reqAgreement).Error
	return err, reqAgreement
}

//修改用户支付密码
func (userService *UserService) UpdateUserPayPassWord(uid uint, password string) (err error) {
	var user autocode.MinappUser
	//password = utils.MD5V([]byte(password)) //加密 密码
	password = utils.HashAndSalt([]byte(password))
	err = global.GVA_DB.Model(&user).Where("id = ?", uid).Update("password", password).Error
	return err
}

var ground = context.Background()

//发送邮件
func (userService *UserService) SendEmail(email string) (err error) {
	subject := "呼唤邻居邮箱验证"
	code := utils.SixNumber()
	body := "您的数字验证码为： " + code
	err = utils.Email(email, subject, body)

	//保存邮箱和发送的验证码
	global.GVA_REDIS.Set(ground, email, code, time.Minute*5)

	return err
}

//取出缓存中的邮箱及验证码并验证
func (userService *UserService) GetEmailCode(mail string) (err error, code string) {
	code, err = global.GVA_REDIS.Get(ground, mail).Result()
	return err, code
}

//保存用户token到redis
func (userService *UserService) SetRedisMinAppJWT(token string, userNum string) (err error) {
	//timer := time.Duration(global.GVA_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.GVA_REDIS.Set(context.Background(), userNum, token, time.Hour*24*29).Err()
	return err
}

//获取用户保存到redis的token
func (userService *UserService) GetRedisMinAppJWT(userNum string) (err error, redisToken string) {
	redisToken, err = global.GVA_REDIS.Get(context.Background(), userNum).Result()
	return err, redisToken
}

//获取用户的信用等级
func (userService *UserService) GetUserCreditRating(uid uint) (err error, list interface{}) {
	var user autocode.MinappUser
	var allCredit []autocode.CreditRating
	err = global.GVA_DB.Model(&user).Where("id = ?", uid).Error
	if err == nil {
		//获取微信小程序设置 分数设置
		var set autocode.MinappSet

		if err = global.GVA_DB.Model(&set).Where("id = ?", 1).Error; err == nil {
			cancelNumber := user.CancelNumber * set.CancelNumber
			allOrderNumber := user.AllOrderNumber * set.AllOrderNumber
			takeOrderNumber := user.TakeOrderNumber * set.TakeOrderNumber
			complaintNumber := user.ComplaintNumber * set.ComplaintNumber

			allNumber := cancelNumber + allOrderNumber + takeOrderNumber + complaintNumber
			var credit autocode.CreditRating
			err = global.GVA_DB.Find(&allCredit).Error
			//判断信用等级，如果都不符合条件则为第一项谨慎沟通
			if errors.Is(global.GVA_DB.Model(&credit).Where("minGrade <= ? || maxGrade >=", allNumber, allNumber).Error, gorm.ErrRecordNotFound) {
				allCredit[0].Checked = 1

				return err, allCredit
			} else {
				var i uint
				if credit.ID > 0 {
					i = credit.ID - 1
				} else {
					i = 0
				}

				allCredit[i].Checked = 1
			}
		}
	}
	return err, allCredit
}

//支付配置
func (userService *UserService) PayConfig() {
	_, r := userService.GetMinAppSet()

	pc := &payConfig.Config{
		AppID:     r.Appid,
		MchID:     r.Mch_id,
		Key:       r.Pay_secret,
		NotifyURL: r.NotifyUrl,
	}
	global.GVA_PAY_CONFIG = pc
}

//用户余额充值
func (userService *UserService) RechargeBalance(uid uint, openid string, createIP string, money float64) (err error, config order.Config) {
	userService.PayConfig()

	var params order.Params
	//balance := money
	//加元转换成人民币
	//exchangeRate := GetMinAppSet()
	//money = money * exchangeRate.ExchangeRate //汇率相乘
	money = utils.Decimal(money) * 100
	totalFee := utils.FloatToString(money)

	tradeNo := "BAL" + utils.UserNum()
	params.TotalFee = totalFee
	params.CreateIP = createIP
	params.Body = "充值余额"
	params.OutTradeNo = tradeNo
	params.OpenID = openid
	params.TradeType = "JSAPI"
	params.NotifyURL = global.GVA_PAY_CONFIG.NotifyURL

	var sendOrder order.Order
	sendOrder = order.Order{
		global.GVA_PAY_CONFIG,
	}

	config, err = sendOrder.BridgeConfig(&params)
	fmt.Println("充值返回信息", config, err)
	if err == nil {
		//余额充值记录
		var record minAppModel.Record
		record.Uid = uid
		record.Type = 1
		record.OrderNum = tradeNo
		record.Money = money //加元
		record.Remark = "充值余额"
		global.GVA_DB.Save(&record)
	}

	return err, config
}

//后台充值用户余额
func (userService *UserService) BackgroundTopUpBalance(uid uint, balance float64, way uint) (err error) {
	var user autocode.MinappUser
	if way == 1 {
		err = global.GVA_DB.Model(&user).Where("id", uid).UpdateColumn("balance", gorm.Expr("balance + ?", balance)).Error
	} else {
		err = global.GVA_DB.Model(&user).Where("id", uid).UpdateColumn("balance", gorm.Expr("balance - ?", balance)).Error
	}

	//err := global.GVA_DB.Model(&user).Where("id", uid).Update("balance", balance).Error
	return err
}

//验证支付密码是否成功
func (userService *UserService) VerifyUserPassword(uid uint, password string) error {
	err, r := userService.GetUserInfo(uid)
	if err != nil {
		return errors.New("查询用户失败")
	} else {
		existPassword := r.Password
		ok := utils.ValidatePasswords(existPassword, []byte(password))

		if ok == true {
			return nil
		} else {
			return errors.New("密码错误")
		}
	}
}

//保存信息到订单中心
func CreateOrderCenter(uid uint, way, orderNum string, money float64) {
	var o autocode.OrderCenter
	o.Uid = uid
	o.Way = way
	o.OrderNum = orderNum
	o.Money = money
	err := global.GVA_DB.Save(&o).Error
	if err != nil {
		global.GVA_LOG.Error("保存到订单中心失败", zap.Error(err))
	}
}

//修改订单中心的状态
func (userService *UserService) UpdateOrderCenterStatus(orderNum string, way string) (err error) {
	var center autocode.OrderCenter
	err = global.GVA_DB.Model(&center).Where("order_num = ? AND way = ?", orderNum, way).Update("status", 1).Error
	return err
}

//获取我的订单中心数据
type CenterData struct {
	Earnings map[string]float64 `json:"earnings"`
	Create   map[string]float64 `json:"create"`
	Take     map[string]float64 `json:"take"`
	Up       map[string]float64 `json:"up"`
}

func (userService *UserService) GetMyOrderCenter(uid uint) (center CenterData) {

	center = CenterData{
		Earnings: make(map[string]float64, 2),
		Create:   make(map[string]float64, 2),
		Take:     make(map[string]float64, 2),
		Up:       make(map[string]float64, 2),
	}

	_, weekDay := utils.GetBeforeTime(-7)
	fmt.Println("一周前日期", weekDay)
	var record minAppModel.Record
	//一周的接单收益
	var weekEarnings float64
	global.GVA_DB.Model(&record).Where("uid = ? AND type = ?", uid, 4).Where("creation_date >= ?", weekDay).Scan(&weekEarnings)
	center.Earnings["weekEarnings"] = weekEarnings
	//所有的收益
	var allEarnings float64
	global.GVA_DB.Model(&record).Where("uid = ? AND type = ?", uid, 4).Scan(&weekEarnings)
	center.Earnings["allEarnings"] = allEarnings

	var orderCenter autocode.OrderCenter
	//我的总发布量
	var createCount int64
	global.GVA_DB.Model(&orderCenter).Where("way", "create").Where("uid", uid).Count(&createCount)
	fmt.Println("订单数量", createCount, uid)
	center.Create["total"] = float64(createCount)

	//我的节省金额
	var createMoney float64
	global.GVA_DB.Model(&orderCenter).Select("sum(money) as createMoney").Where("way", "create").Where("uid", uid).Scan(&createMoney)
	center.Create["money"] = createMoney
	//我的总接单量
	var takeCount int64
	global.GVA_DB.Model(&orderCenter).Where("way", "take").Where("uid", uid).Count(&takeCount)
	center.Take["total"] = float64(takeCount)
	//我的总赚取金额
	var takeMoney float64
	global.GVA_DB.Model(&orderCenter).Select("sum(money) as takeMoney").Where("way", "take").Where("uid", uid).Scan(&takeMoney)
	center.Take["money"] = takeMoney
	//我的预约量
	var upCount int64
	global.GVA_DB.Model(&orderCenter).Where("way", "up").Where("uid", uid).Count(&upCount)
	center.Up["total"] = float64(upCount)
	//我的私信留言
	var message minAppModel.SecondMessage
	var messageCount int64
	global.GVA_DB.Model(&message).Where("uid = ?", uid).Count(&messageCount)
	center.Up["num"] = float64(messageCount)
	return center
}

//修改用记所在地址
func (userService *UserService) UpdateUserArea(uid uint, area string) (err error) {
	var user autocode.MinappUser
	err = global.GVA_DB.Model(&user).Where("id", uid).Update("area", area).Error
	return err
}

func (userService *UserService) EditUser(user autocode.MinappUser, uid uint) error {
	err := global.GVA_DB.Model(&user).Where("id", uid).Updates(&user).Error
	return err
}

//获取今天的汇率
func GetRate() float64 {
	//获取缓存中的汇率信息
	val, err := global.GVA_REDIS.Get(context.Background(), "rate").Result()
	fmt.Println("缓存内容", val, err)
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	if val != "" {
		f, _ := strconv.ParseFloat(val, 64)

		return f
	} else {
		//panic(err.Error())
		params := url.Values{}                                                     //参数集合
		reqUrl, err := url.Parse("https://api.exchangerate-api.com/v4/latest/CAD") //请求地址
		if err != nil {
			panic(err.Error())
		}
		//params.Set("age", "10") //设置参数
		//params.Set("id", "100")
		reqUrl.RawQuery = params.Encode()      //组合url
		resp, err := http.Get(reqUrl.String()) //发起get请求

		if err != nil {
			panic(err.Error())
		}
		defer resp.Body.Close()                //关闭请求
		body, err := ioutil.ReadAll(resp.Body) //解析请求信息
		if err != nil {
			panic(err.Error())
		}
		var results map[string]interface{}   //请求结果集
		err = json.Unmarshal(body, &results) //转换为map
		if err != nil {
			panic(err.Error())
		}

		cny := results["rates"]
		aa := cny.(map[string]interface{})

		f := aa["CNY"].(float64)
		fmt.Println("结果", f)
		global.GVA_REDIS.Set(context.Background(), "rate", f, 24*time.Hour)
		return f
	}
}

//获取getAccessToken
func (userService *UserService) GetAccessToken() string {
	val, err := global.GVA_REDIS.Get(context.Background(), "access_token").Result()
	if err != nil {
		fmt.Println(err)
	}

	if val != "" {
		return val
	}

	params := url.Values{}                                              //参数集合
	reqUrl, err := url.Parse("https://api.weixin.qq.com/cgi-bin/token") //请求地址
	if err != nil {
		panic(err.Error())
	}
	minAppSet := GetMinAppSet()
	params.Set("grant_type", "client_credential") //设置参数
	params.Set("appid", minAppSet.Appid)
	params.Set("secret", minAppSet.App_secret)
	reqUrl.RawQuery = params.Encode()      //组合url
	resp, err := http.Get(reqUrl.String()) //发起get请求

	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()                //关闭请求
	body, err := ioutil.ReadAll(resp.Body) //解析请求信息
	if err != nil {
		panic(err.Error())
	}
	var results map[string]interface{}   //请求结果集
	err = json.Unmarshal(body, &results) //转换为map
	if err != nil {
		panic(err.Error())
	}
	t := results["access_token"].(string)
	global.GVA_REDIS.Set(context.Background(), "access_token", t, time.Hour)
	return t
}

type Data struct {
	JumpWxa Params `json:"jump_wxa"`
}

type Params struct {
	Query string `json:"query"`
}

// post请求
func (userService *UserService) PostRequest(urls string, uid uint) string {
	id := strconv.Itoa(int(uid))
	val, err := global.GVA_REDIS.Get(context.Background(), "open_link"+id).Result()
	if err != nil {
		fmt.Println(err)
	}
	if val != "" {
		return val
	}
	//jsonStr := []byte(`{"jump_wxa":{"query":"id=3"}}`)
	q := "id=" + id
	jsonStr := Data{
		JumpWxa: Params{
			Query: q,
		},
	}
	//data := url.Values{"jump_wxa": {"query:{id=}" + id}}
	//body := strings.NewReader("")
	s, _ := json.Marshal(jsonStr)
	str := bytes.NewBuffer(s)
	fmt.Println("请求的数据", str)
	urlStr := urls
	fmt.Println("请求的url", urlStr)
	//resp, err := http.Post(urlStr, "application/x-www-form-urlencoded", body)
	resp, err := http.Post(urlStr, "application/json", str)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyC, _ := ioutil.ReadAll(resp.Body)

	var jsonMap map[string]interface{}
	err = json.Unmarshal(bodyC, &jsonMap)
	if err != nil {
		fmt.Println(err)
	}
	code := jsonMap["errcode"].(float64)
	if code == 0 {
		openlink := jsonMap["openlink"].(string)
		global.GVA_REDIS.Set(context.Background(), "open_link"+id, openlink, 24*29*time.Hour)
		return openlink
	}
	fmt.Println("结果", jsonMap)
	return "获取链接失败"
}

//用户邀请的记录
func (userService *UserService) CreateInviteRecord(uid, invite uint) error {
	var record autocode.InviteRecord
	if !errors.Is(global.GVA_DB.Where("uid = ?", uid).First(&record).Error, gorm.ErrRecordNotFound) {
		return errors.New("你已经被邀请过")
	}

	record.Uid = uid
	record.Invite = invite
	record.CreationDate = utils.TodayTime()
	err := global.GVA_DB.Save(&record).Error

	var count int64
	global.GVA_DB.Model(&record).Where("invite = ?", invite).Count(&count)
	//判断邀请用户是否满足赠送条件
	if count > 0 {
		var coupon autocode.PayCoupon
		global.GVA_DB.Where("type = ?", 2).Find(&coupon)
		people := utils.FloatToInt(coupon.Money)
		if int(count)%people == 0 {
			var give autocode.ReleaseRecord
			give.Uid = invite
			give.CouponId = coupon.CouponId
			give.Title = coupon.Title
			give.Start = coupon.Start
			give.Over = coupon.Over
			give.Condition = coupon.Condition
			give.Way = coupon.Way
			give.Total = coupon.Total
			give.Remark = "邀请用户赠送的优惠券"
			err = global.GVA_DB.Save(&give).Error
			if err != nil {
				return errors.New("赠送优惠失败")
			}
			return err
		}
	}

	return err
}

//用户邀请记录
func (userService *UserService) UserInviteRecord(info autoCodeReq.InviteRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&autocode.InviteRecord{})
	var record []autocode.InviteRecord
	db.Where("invite = ?", info.Invite)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("UserInfo").Order("id DESC").Find(&record).Error
	return err, record, total
}
