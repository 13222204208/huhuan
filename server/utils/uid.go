package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

func Today() string {
	startTimestamp := time.Now().Unix()
	//获得时间戳
	startTimeStr := time.Unix(startTimestamp, 0).Format("2006-01-02")
	return startTimeStr
}

func TodayTime() string {
	startTimestamp := time.Now()

	return startTimestamp.Format("2006-01-02 15:04:05")
}

// GetBeforeTime 获取n天前的秒时间戳、日期时间戳
// _day为负则代表取前几天，为正则代表取后几天，0则为今天
func GetBeforeTime(_day int) (int64, string) {
	// 时区
	//timeZone, _ := time.LoadLocation(ServerInfo["timezone"])
	timeZone := time.FixedZone("CST", 8*3600) // 东八区

	// 前n天
	nowTime := time.Now().In(timeZone)
	beforeTime := nowTime.AddDate(0, 0, _day)

	// 时间转换格式
	beforeTimeS := beforeTime.Unix()                             // 秒时间戳
	beforeDate := time.Unix(beforeTimeS, 0).Format("2006-01-02") // 固定格式的日期时间戳

	return beforeTimeS, beforeDate
}

//不带秒显示的时间
func TodayMinutes() string {
	startTimestamp := time.Now()

	return startTimestamp.Format("2006-01-02 15:04")
}

func UserNum() string {

	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	startTimestamp := time.Now().Unix()
	//获得时间戳
	startTimeStr := time.Unix(startTimestamp, 0).Format("20060102") //把时间戳转换成时间,并格式化为年月日
	nowTime := sup(hour, 2) + sup(minute, 2) + sup(second, 2)
	randNum := randInt(100, 999)

	code := startTimeStr + nowTime + strconv.Itoa(randNum)
	code = code[2:]
	return code
}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

//对长度不足n的数字前面补0
func sup(i int, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}

func SixNumber() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return code
}

//保留两位小数，一般做为显示金额
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

//float64转int
func FloatToInt(f float64) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%1.0f", f))
	return i
}

//float64转string
func FloatToString(f float64) string {
	var s string
	i := FloatToInt(f)
	s = strconv.Itoa(i)
	return s
}

//日期转时间戳
func DateToTime(timeStr string) int64 {
	var LOC, _ = time.LoadLocation("Asia/Shanghai")
	//要转换成时间日期的格式模板（go诞生时间，模板必须是这个时间）
	temp := "2006-01-02"
	tim, _ := time.ParseInLocation(temp, timeStr, LOC)
	c := tim.Unix()
	return c
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func GetBetweenDates(startDate, endDate string) []string {
	var d []string
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(startDate) {
		timeFormatTpl = timeFormatTpl[0:len(startDate)]
	}
	date, err := time.Parse(timeFormatTpl, startDate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, endDate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

// 加密密码
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {

	}
	return string(hash)
}

// 验证密码
//byteHash 加密过的密码
//plainPwd 需要验证的密码
func ValidatePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd) //加密过的代码
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

//生成房间号
func CreateRoomName(id1, id2 int) string {
	if id1 > id2 {
		temp := id2
		id2 = id1
		id1 = temp
	}
	one := strconv.Itoa(id1)
	two := strconv.Itoa(id2)
	one = ZeroFillByStr(one, 9, true)
	two = ZeroFillByStr(two, 9, true)
	str := "R" + one + two
	return str
}

//  ZeroFillByStr
//  @Description: 字符串补零
//  @param str :需要操作的字符串
//  @param resultLen 结果字符串的长度
//  @param reverse true 为前置补零，false 为后置补零
//  @return string
//
func ZeroFillByStr(str string, resultLen int, reverse bool) string {
	if len(str) > resultLen || resultLen <= 0 {
		return str
	}
	if reverse {
		return fmt.Sprintf("%0*s", resultLen, str) //不足前置补零
	}
	result := str
	for i := 0; i < resultLen-len(str); i++ {
		result += "0"
	}
	return result
}
