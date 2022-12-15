package minapp

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userPublicRouter := Router.Group("user")
	userPrivateRouter := Router.Group("user").Use(middleware.JWTAuthMiddleware())

	exaFileUploadAndDownloadApi := v1.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi

	userApi := v1.ApiGroupApp.MinAppApiGroup.UserApi
	{
		userPublicRouter.POST("login", userApi.Login)
		userPublicRouter.GET("agreement", userApi.Agreement)
		//图片上传
		userPublicRouter.POST("upload", exaFileUploadAndDownloadApi.UploadFile)
	}

	{
		userPrivateRouter.GET("info", userApi.GetUserInfo)

		//修改或保存密码
		userPrivateRouter.POST("password", userApi.UpdateUserPayPassWord)

		//发送验证码到邮箱
		userPrivateRouter.POST("email", userApi.SendCodeToMail)

		//获取用户信用等级
		userPrivateRouter.GET("credit", userApi.GetUserCreditRating)

		//用户余额充值
		userPrivateRouter.POST("recharge", userApi.RechargeBalance)

		//用户提现
		userPrivateRouter.POST("withdraw", userApi.BalanceWithdraw)

		//验证用户支付密码
		userPrivateRouter.POST("verify-password", userApi.UserPasswordVerify)

		//订单重新支付
		userPrivateRouter.POST("again-pay", userApi.AgainPayOrder)

		//获取我的订单中心数据
		userPrivateRouter.GET("order-center", userApi.MyOrderCenter)

		//修改用户地址
		userPrivateRouter.PUT("area", userApi.UpdateUserArea)

		//获取用户钱包记录
		userPrivateRouter.GET("money-record", userApi.GetUserMoneyRecords)

		//获取小程序Url地址
		userPrivateRouter.GET("url", userApi.GetMinAppUrl)

		//创建邀请用户记录
		userPrivateRouter.POST("invite", userApi.CreateUserInviteRecord)

		userPrivateRouter.POST("edit", userApi.EditUser)

		//获取用户邀请记录
		userPrivateRouter.GET("invite-record", userApi.GetUserInviteRecord)

		//获取用户的手机号
		userPrivateRouter.POST("phone", userApi.GetPhone)
	}
}
