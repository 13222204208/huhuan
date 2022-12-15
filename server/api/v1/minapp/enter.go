package minapp

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	UserApi
	HelpApi
	BannerApi
	AddressApi
	ComplaintApi
	CouponApi
	CheckApi
	SignApi
	NotifyApi
	MessageApi
	MineApi
}

var (
	userService      = service.ServiceGroupApp.MinAppServiceGroup.UserService
	helpService      = service.ServiceGroupApp.MinAppServiceGroup.HelpService
	bannerService    = service.ServiceGroupApp.MinAppServiceGroup.BannerService
	addressService   = service.ServiceGroupApp.MinAppServiceGroup.AddressService
	complaintService = service.ServiceGroupApp.MinAppServiceGroup.ComplaintService
	couponService    = service.ServiceGroupApp.MinAppServiceGroup.CouponService

	helpRepairService = service.ServiceGroupApp.AutoCodeServiceGroup.HelpRepairService
	helpWorkService   = service.ServiceGroupApp.AutoCodeServiceGroup.HelpWorkService
	helpFetchService  = service.ServiceGroupApp.AutoCodeServiceGroup.HelpFetchService
	helpBuyService    = service.ServiceGroupApp.AutoCodeServiceGroup.HelpBuyService

	//微信支付公共方法
	payService = service.ServiceGroupApp.MinAppServiceGroup.PayService
	//我发布的二手闲置订单
	mineService = service.ServiceGroupApp.MinAppServiceGroup.MineService

	//国内捎带服务
	helpIncidentallyService = service.ServiceGroupApp.AutoCodeServiceGroup.HelpIncidentallyService

	//组团拼车
	helpCarpoolService = service.ServiceGroupApp.AutoCodeServiceGroup.HelpCarpoolService

	//团购接龙
	helpGrouponService = service.ServiceGroupApp.AutoCodeServiceGroup.HelpGrouponService

	//二手闲置
	helpSecondService = service.ServiceGroupApp.AutoCodeServiceGroup.HelpSecondService

	//获取二手闲置分类
	helpSecondCategoryService = service.ServiceGroupApp.AutoCodeServiceGroup.CategoryService

	//司机认证
	checkDriverService = service.ServiceGroupApp.AutoCodeServiceGroup.CheckDriverService

	//商家认证
	checkMerchantService = service.ServiceGroupApp.AutoCodeServiceGroup.CheckMerchantService

	//身份认证
	checkIdentityService = service.ServiceGroupApp.AutoCodeServiceGroup.CheckIdentityService

	//微信用户修改
	minAppUserService = service.ServiceGroupApp.AutoCodeServiceGroup.MinappUserService

	//获取微信小程序设置
	minAppSettingService = service.ServiceGroupApp.AutoCodeServiceGroup.MinappSetService

	//用户签到
	userSignInService = service.ServiceGroupApp.MinAppServiceGroup.SignService

	//优惠券
	AutoCodeCouponService = service.ServiceGroupApp.AutoCodeServiceGroup.CouponService

	//发放的优惠券记录
	releaseRecordService = service.ServiceGroupApp.AutoCodeServiceGroup.ReleaseRecordService

	//用户提现
	withdrawService = service.ServiceGroupApp.AutoCodeServiceGroup.WithdrawService

	//在线聊天
	messageService = service.ServiceGroupApp.MinAppServiceGroup.MessageService

	//用户钱包记录
	recordService = service.ServiceGroupApp.MinAppServiceGroup.RecordService
)
