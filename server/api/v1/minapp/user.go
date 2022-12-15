package minapp

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	minAppRequest "github.com/flipped-aurora/gin-vue-admin/server/model/minapp/request"
	minAppResponse "github.com/flipped-aurora/gin-vue-admin/server/model/minapp/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"time"
)

type UserApi struct{}

func (a *UserApi) GetPhone(c *gin.Context) {
	var p minAppRequest.GetPhone
	_ = c.ShouldBindJSON(&p)

	if err, phone := userService.GetPhone(p.Code, p.EncryptedData, p.Iv); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithDetailed(gin.H{"phone": phone}, "获取成功", c)
	}
}

func (a *UserApi) Login(c *gin.Context) {
	var l minAppRequest.Login
	_ = c.ShouldBindJSON(&l)

	errs := utils.Verify(l, utils.MinAppLoginVerify)
	if errs != nil {
		response.FailWithMessage(errs.Error(), c)
		return
	}

	err, openid := userService.GetOpenid(l.Code)
	if err != nil {
		response.FailWithMessage("获取openid失败", c)
		return
	} else {
		_, userinfo := userService.Login(openid, l.Avatar, l.Nickname)

		//token, err := genToken(userinfo.ID, userinfo.Openid, userinfo.UserNum)
		//if err != nil {
		//	global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		//	response.FailWithMessage("获取token失败", c)
		//	return
		//} else {
		//	response.OkWithDetailed(minAppResponse.LoginResponse{
		//		User:  userinfo,
		//		Token: token,
		//	}, "成功", c)
		//	return
		//}

		if err, tokenStr := userService.GetRedisMinAppJWT(userinfo.UserNum); err == redis.Nil || tokenStr == "" {
			token, err := genToken(userinfo.ID, userinfo.Openid, userinfo.UserNum)
			if err != nil {
				global.GVA_LOG.Error("获取token失败!", zap.Error(err))
				response.FailWithMessage("获取token失败", c)
				return
			} else {
				if err := userService.SetRedisMinAppJWT(token, userinfo.UserNum); err != nil {
					response.FailWithMessage("保存token到redis失败", c)
					return
				}

				response.OkWithDetailed(minAppResponse.LoginResponse{
					User:  userinfo,
					Token: token,
				}, "成功", c)
				return
			}

		} else if err != nil {
			panic(err)
		} else {
			response.OkWithDetailed(minAppResponse.LoginResponse{
				User:  userinfo,
				Token: tokenStr,
			}, "返回成功", c)
		}

	}
}

// GenToken 生成JWT
func genToken(id uint, openid string, userNum string) (string, error) {
	// 创建一个我们自己的声明
	c := minAppRequest.MyClaims{
		id, // 自定义字段
		openid,
		userNum,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(minAppRequest.TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-project",                                             // 签发人
		},
	}
	fmt.Println("token内容", c)
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(minAppRequest.MySecret)
}

func (a *UserApi) GetUserInfo(c *gin.Context) {
	uid := c.MustGet("id").(uint)
	if err, ReqUser := userService.GetUserInfo(uid); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
	}
}

func (a *UserApi) Agreement(c *gin.Context) {
	if err, Agreement := userService.GetUserAgreement(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"agreementInfo": Agreement}, "获取成功", c)
	}
}

func (a *UserApi) UpdateUserPayPassWord(c *gin.Context) {
	var p minAppRequest.UpdatePayPassWord
	_ = c.ShouldBindJSON(&p)

	password := p.Password

	if len(password) != 6 {
		response.FailWithMessage("请填写六位支付密码", c)
		return
	}

	uid := c.MustGet("id").(uint)
	if err := userService.UpdateUserPayPassWord(uid, password); err != nil {
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithMessage("成功", c)
	}
}

//发送验证码到邮箱
func (a *UserApi) SendCodeToMail(c *gin.Context) {
	var e minAppRequest.SendEmail
	_ = c.ShouldBindJSON(&e)
	if len(e.Email) < 6 {
		response.FailWithMessage("请填写正确的邮箱", c)
		return
	}

	if len(e.Code) == 6 {
		emailCode := e.Code
		if err, code := userService.GetEmailCode(e.Email); err == redis.Nil {
			response.FailWithMessage("失败", c)
		} else {
			if code == emailCode {
				response.OkWithData("验证通过", c)
				return
			} else {
				response.OkWithData("验证失败", c)
			}
		}
	}

	if err, _ := userService.GetEmailCode(e.Email); err == redis.Nil {
		if err := userService.SendEmail(e.Email); err != nil {
			response.FailWithMessage("发送失败", c)
		} else {
			response.OkWithData("发送成功", c)
		}
	} else if err != nil {
		panic(err)
	} else {
		response.FailWithMessage("验证码已发送", c)
	}

}

//获取用户的信用等级
func (a *UserApi) GetUserCreditRating(c *gin.Context) {
	uid := c.MustGet("id").(uint)
	if err, result := userService.GetUserCreditRating(uid); err != nil {
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithDetailed(gin.H{"list": result}, "成功", c)
	}
}

//用户余额充值
//满足充值额度时，赠送优惠券
func (a *UserApi) RechargeBalance(c *gin.Context) {
	var balance minAppRequest.RechargeBalance
	_ = c.ShouldBindJSON(&balance)
	if balance.Balance > 0 {
		openid := c.MustGet("openid").(string)
		uid := c.MustGet("id").(uint)
		createIP := c.ClientIP()
		totalFee := balance.Balance
		if err, result := userService.RechargeBalance(uid, openid, createIP, totalFee); err != nil {
			response.FailWithMessage("失败", c)
		} else {
			response.OkWithData(gin.H{"params": result}, c)
		}

	} else {
		response.FailWithMessage("额度必须大于0", c)
	}

}

//微信余额提现
func (a *UserApi) BalanceWithdraw(c *gin.Context) {
	var w minAppRequest.BalanceWithdraw
	var withdraw autocode.Withdraw
	_ = c.ShouldBindJSON(&w)
	money := w.Money
	if money < 1 {
		response.FailWithMessage("请填写正确的金额", c)
		return
	}
	uid := c.MustGet("id").(uint)
	if err, result := userService.GetUserInfo(uid); err != nil {
		response.FailWithMessage("获取用户信息失败", c)
		return
	} else {
		if result.Balance < money {
			response.FailWithMessage("你的余额不足", c)
			return
		}
		withdraw.Money = money
		withdraw.Uid = uid
		withdraw.UserNum = result.UserNum
		withdraw.Name = result.Name
		withdraw.Phone = result.Phone
		withdraw.ApplyTime = utils.Today()
		if err := withdrawService.CreateWithdraw(withdraw); err != nil {
			response.FailWithMessage("提交申请失败", c)
		} else {
			response.OkWithMessage("成功", c)
		}
	}
}

//用户支付密码验证
func (a *UserApi) UserPasswordVerify(c *gin.Context) {
	uid := c.MustGet("id").(uint)
	var p minAppRequest.VerifyUserPassword
	_ = c.ShouldBindJSON(&p)
	fmt.Println("用户id:", uid)
	if len(p.Password) != 6 {
		response.FailWithMessage("请填写合法密码", c)
	} else {
		if err := userService.VerifyUserPassword(uid, p.Password); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			response.OkWithMessage("验证成功", c)
			return
		}
	}
}

//订单重新支付
func (a *UserApi) AgainPayOrder(c *gin.Context) {
	var order minAppRequest.AgainPayOrder
	_ = c.ShouldBindJSON(&order)
	orderNum := order.OrderNum
	money := order.Money

	if len(orderNum) < 16 {
		response.FailWithMessage("请填写正确的订单号", c)
		return
	}

	if money <= 0 {
		response.FailWithMessage("金额不能为0", c)
	} else {
		openid := c.MustGet("openid").(string)
		if err, config := payService.AgainPayOrder(openid, c.ClientIP(), money, orderNum); err != nil {
			response.FailWithMessage("参数错误", c)
		} else {
			response.OkWithData(gin.H{
				"config": config,
			}, c)
		}
	}
}

//我的订单中心数据接口
func (a *UserApi) MyOrderCenter(c *gin.Context) {
	uid := c.MustGet("id").(uint)
	result := userService.GetMyOrderCenter(uid)
	response.OkWithData(gin.H{
		"list": result,
	}, c)
}

//修改用户地址
func (a *UserApi) UpdateUserArea(c *gin.Context) {
	var userinfo autocode.MinappUser
	_ = c.ShouldBindJSON(&userinfo)
	uid := c.MustGet("id").(uint)
	area := userinfo.Area
	if len(area) < 2 {
		response.FailWithMessage("请填写修改的地址", c)
		return
	}
	if err := userService.UpdateUserArea(uid, userinfo.Area); err != nil {
		response.FailWithMessage("修改地址错误", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

func (a *UserApi) EditUser(c *gin.Context) {
	var userinfo autocode.MinappUser
	_ = c.ShouldBindJSON(&userinfo)
	uid := c.MustGet("id").(uint)
	fmt.Println("用户信息", userinfo)
	if err := userService.EditUser(userinfo, uid); err != nil {
		response.FailWithMessage("错误", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

//获取用户钱包记录
func (a *UserApi) GetUserMoneyRecords(c *gin.Context) {
	var pageInfo minAppRequest.UserMoneyRecord
	_ = c.ShouldBindQuery(&pageInfo)
	uid := c.MustGet("id").(uint)
	pageInfo.Uid = uid

	if err, list, total := recordService.GetUserMoneyRecord(pageInfo); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

//获取微信小程序链接url地址
func (a *UserApi) GetMinAppUrl(c *gin.Context) {
	uid := c.MustGet("id").(uint)
	r := userService.GetAccessToken()

	link := userService.PostRequest("https://api.weixin.qq.com/wxa/generatescheme?access_token="+r, uid)
	response.OkWithMessage(link, c)
}

//创建用户邀请记录
func (a *UserApi) CreateUserInviteRecord(c *gin.Context) {
	var r autocode.InviteRecord
	_ = c.ShouldBindJSON(&r)
	uid := c.MustGet("id").(uint)
	if r.Invite < 1 {
		response.FailWithMessage("请填写邀请人的id", c)
		return
	}

	if err := userService.CreateInviteRecord(uid, r.Invite); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建记录成功", c)
	}
}

//用户的邀请记录
func (a *UserApi) GetUserInviteRecord(c *gin.Context) {
	var pageInfo autocodeReq.InviteRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	uid := c.MustGet("id").(uint)
	pageInfo.Invite = uid
	if err, list, total := userService.UserInviteRecord(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
