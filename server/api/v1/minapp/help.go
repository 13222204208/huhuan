package minapp

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	minAppRequest "github.com/flipped-aurora/gin-vue-admin/server/model/minapp/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type HelpApi struct{}

//发布帮我修
func (a *HelpApi) CreateHelpRepair(c *gin.Context) {
	var h autocode.HelpRepair
	_ = c.ShouldBindJSON(&h)
	if err := utils.Verify(h, utils.MinAppHelpRepair); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := c.MustGet("id").(uint)
	userNum := c.MustGet("userNum").(string)

	//判断使用的支付方式 h.Way 如果为wxPay 则为余额支付 balancePay 为微信支付
	payType := h.Way

	orderNum := "BWX" + utils.UserNum()
	money := h.Price

	//判断是否使用了优惠券
	if h.CouponId > 0 {
		if err, discounts := payService.UserCoupon(uid, h.CouponId, money); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			h.CouponAmount = discounts //优惠券优惠的金额
			money -= discounts         //支付减去金额
		}
	}

	//判断是否是余额支付
	if payType == "balancePay" {
		//余额支付的流程
		if err := payService.BalancePayMethod(orderNum, money, uid); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			h.PayStatus = 1 //表示已经支付成功
			h.Status = 1    //订单待接单
		}
	}
	h.OrderNum = orderNum
	h.Uid = uid
	h.UserNum = userNum

	if err := helpService.CreateHelpRepair(&h); err != nil {
		global.GVA_LOG.Error("失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		//如果是微信支付
		if payType == "wxPay" {
			openid := c.MustGet("openid").(string)
			ip := c.ClientIP()
			if err, config := payService.UserPay(openid, ip, money, orderNum); err != nil {
				response.FailWithMessage("生成参数错误", c)
				return
			} else {
				response.OkWithData(gin.H{"config": config}, c)
				return
			}
		}
		response.OkWithMessage("创建成功", c)

	}

}

//获取帮我修订单
//id 订单的ID
//address 搜索的内容信息
func (a *HelpApi) GetHelpRepair(c *gin.Context) {
	var pageInfo autocodeReq.HelpRepairSearch
	_ = c.ShouldBindQuery(&pageInfo)

	if pageInfo.ID > 0 {
		if err, result := helpRepairService.GetHelpRepair(pageInfo.ID); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"helpRepair": result}, c)
		}
		return
	}

	if err, list, total := helpRepairService.GetHelpRepairInfoList(pageInfo); err != nil {
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

//发布帮我做
func (a *HelpApi) CreateHelpWork(c *gin.Context) {
	var h autocode.HelpWork
	_ = c.ShouldBindJSON(&h)
	if err := utils.Verify(h, utils.MinAppHelpWork); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := c.MustGet("id").(uint)
	userNum := c.MustGet("userNum").(string)

	//判断使用的支付方式 h.Way 如果为wxPay 则为余额支付 balancePay 为微信支付
	payType := h.Way

	orderNum := "BWZ" + utils.UserNum()
	money := h.Price

	//判断是否使用了优惠券
	if h.CouponId > 0 {
		if err, discounts := payService.UserCoupon(uid, h.CouponId, money); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			h.CouponAmount = discounts //优惠券优惠的金额
			money -= discounts         //支付减去金额
		}
	}

	//判断是否是余额支付
	if payType == "balancePay" {
		//余额支付的流程
		if err := payService.BalancePayMethod(orderNum, money, uid); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			h.PayStatus = 1 //表示已经支付成功
			h.Status = 1    //订单待接单
		}
	}
	h.OrderNum = orderNum
	h.Uid = uid
	h.UserNum = userNum

	if err := helpService.CreateHelpWork(&h); err != nil {
		global.GVA_LOG.Error("失败", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {

		//如果是微信支付
		if payType == "wxPay" {
			openid := c.MustGet("openid").(string)
			ip := c.ClientIP()
			if err, config := payService.UserPay(openid, ip, money, orderNum); err != nil {
				response.FailWithMessage("生成参数错误", c)
				return
			} else {
				response.OkWithData(gin.H{"config": config}, c)
				return
			}
		}

		response.OkWithMessage("创建成功", c)

	}

}

//获取帮我做订单
//id 订单的ID
//address 搜索的内容信息
func (a *HelpApi) GetHelpWork(c *gin.Context) {
	var pageInfo autocodeReq.HelpWorkSearch
	_ = c.ShouldBindQuery(&pageInfo)

	if pageInfo.ID > 0 {
		if err, result := helpWorkService.GetHelpWork(pageInfo.ID); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"helpWork": result}, c)
		}
		return
	}

	if err, list, total := helpWorkService.GetHelpWorkInfoList(pageInfo); err != nil {
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

//发布帮我取
func (a *HelpApi) CreateHelpFetch(c *gin.Context) {
	//var h minAppRequest.CreateHelpRepair
	var h autocode.HelpFetch
	_ = c.ShouldBindJSON(&h)
	if err := utils.Verify(h, utils.MinAppHelpFetch); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if h.FetchAddress == h.ReceiveAddress {
		response.FailWithMessage("地址不能相同", c)
		return
	}

	uid := c.MustGet("id").(uint)
	userNum := c.MustGet("userNum").(string)

	//判断使用的支付方式 h.Way 如果为wxPay 则为余额支付 balancePay 为微信支付
	payType := h.Way
	orderNum := "BWQ" + utils.UserNum()
	money := h.Price

	//自动匹配用户可以使用的优惠券
	if payType == "mateCoupon" && money > 0 {
		if err, r := payService.AutoSelectCoupon(uid, money); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			response.OkWithData(gin.H{"coupon": r}, c)
			return
		}
	}

	//判断是否使用了优惠券
	if h.CouponId > 0 {
		if err, discounts := payService.UserCoupon(uid, h.CouponId, money); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			h.CouponAmount = discounts //优惠券优惠的金额
			money -= discounts         //支付减去金额
		}
	}

	//判断是否是余额支付
	if payType == "balancePay" {
		//余额支付的流程
		if err := payService.BalancePayMethod(orderNum, money, uid); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			h.PayStatus = 1 //表示已经支付成功
			h.Status = 1    //订单待接单
		}
	}

	h.OrderNum = orderNum
	h.Uid = uid
	h.UserNum = userNum

	if err := helpService.CreateHelpFetch(&h); err != nil {
		global.GVA_LOG.Error("失败", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		//如果是微信支付
		if payType == "wxPay" {
			openid := c.MustGet("openid").(string)
			ip := c.ClientIP()
			if err, config := payService.UserPay(openid, ip, money, orderNum); err != nil {
				response.FailWithMessage("生成参数错误", c)
				return
			} else {
				response.OkWithData(gin.H{"config": config}, c)
				return
			}
		}

		response.OkWithMessage("创建成功", c)

	}
}

//获取帮我取订单
//id 订单的ID
//address 搜索的内容信息
func (a *HelpApi) GetHelpFetch(c *gin.Context) {
	var pageInfo autocodeReq.HelpFetchSearch
	_ = c.ShouldBindQuery(&pageInfo)

	if pageInfo.ID > 0 {
		if err, result := helpFetchService.GetHelpFetch(pageInfo.ID); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"helpFetch": result}, c)
		}
		return
	}

	if err, list, total := helpFetchService.GetHelpFetchInfoList(pageInfo); err != nil {
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

//根据订单号获取订单详情
func (a *HelpApi) GetOrder(c *gin.Context) {
	var o autocodeReq.Order
	_ = c.ShouldBindQuery(&o)
	orderNum := o.OrderNum
	fmt.Println("求的订单号", orderNum)
	s := orderNum[:3]
	fmt.Println("求的订单号", orderNum, s)
	switch s {
	case "SEC":
		var h autocode.HelpSecond
		global.GVA_DB.Where("order_num = ?", orderNum).First(&h)
		response.OkWithData(gin.H{"info": h}, c)
		return

	case "BWX":
		var h autocode.HelpRepair
		global.GVA_DB.Where("order_num = ?", orderNum).First(&h)
		response.OkWithData(gin.H{"info": h}, c)
		return

	case "GRO":
		var h autocode.HelpGroupon
		global.GVA_DB.Where("order_num = ?", orderNum).First(&h)
		response.OkWithData(gin.H{"info": h}, c)
		return

	case "BWQ":
		var h autocode.HelpFetch
		global.GVA_DB.Where("order_num = ?", orderNum).First(&h)
		response.OkWithData(gin.H{"info": h}, c)
		return

	case "BWM":
		var h autocode.HelpBuy
		global.GVA_DB.Where("order_num = ?", orderNum).First(&h)
		response.OkWithData(gin.H{"info": h}, c)
		return

	case "BWZ":
		var h autocode.HelpWork
		global.GVA_DB.Where("order_num = ?", orderNum).First(&h)
		response.OkWithData(gin.H{"info": h}, c)
		return
	}
}

//发布帮我买
func (a *HelpApi) CreateHelpBuy(c *gin.Context) {
	var h autocode.HelpBuy
	_ = c.ShouldBindJSON(&h)
	if err := utils.Verify(h, utils.MinAppHelpBuy); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if h.BuyAddress == h.ReceiveAddress {
		response.FailWithMessage("地址不能相同", c)
		return
	}
	uid := c.MustGet("id").(uint)
	userNum := c.MustGet("userNum").(string)

	//判断使用的支付方式 h.Way 如果为wxPay 则为余额支付 balancePay 为微信支付
	payType := h.Way
	orderNum := "BWM" + utils.UserNum()
	money := h.Price

	if money < 0.002 {
		response.FailWithMessage("金额太小", c)
		return
	}

	//自动匹配用户可以使用的优惠券
	if payType == "mateCoupon" && money > 0 {
		if err, r := payService.AutoSelectCoupon(uid, money); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			response.OkWithData(gin.H{"coupon": r}, c)
			return
		}
	}

	//判断是否使用了优惠券
	if h.CouponId > 0 {
		if err, discounts := payService.UserCoupon(uid, h.CouponId, money); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			h.CouponAmount = discounts //优惠券优惠的金额
			money -= discounts         //支付减去金额
		}
	}

	//判断是否是余额支付
	if payType == "balancePay" {
		//余额支付的流程
		if err := payService.BalancePayMethod(orderNum, money, uid); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			h.PayStatus = 1 //表示已经支付成功
			h.Status = 1    //订单待接单
		}
	}

	h.OrderNum = orderNum
	h.Uid = uid
	h.UserNum = userNum

	if err := helpService.CreateHelpBuy(&h); err != nil {
		global.GVA_LOG.Error("失败", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		//如果是微信支付
		if payType == "wxPay" {
			openid := c.MustGet("openid").(string)
			ip := c.ClientIP()
			if err, config := payService.UserPay(openid, ip, money, orderNum); err != nil {
				response.FailWithMessage("生成参数错误", c)
				return
			} else {
				response.OkWithData(gin.H{"config": config}, c)
				return
			}
		}

		response.OkWithMessage("创建成功", c)

	}
}

//获取帮我买订单
//id 订单的ID
//address 搜索的内容信息
func (a *HelpApi) GetHelpBuy(c *gin.Context) {
	var pageInfo autocodeReq.HelpBuySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if pageInfo.ID > 0 {
		if err, result := helpBuyService.GetHelpBuy(pageInfo.ID); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"helpFetch": result}, c)
		}
		return
	}

	if err, list, total := helpBuyService.GetHelpBuyInfoList(pageInfo); err != nil {
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

//创建国内捎带订单
func (a *HelpApi) CreateHelpIncidentally(c *gin.Context) {
	var h autocode.HelpIncidentally
	_ = c.ShouldBindJSON(&h)

	//判断如果我是乘机人 type = 1
	typeNum := h.Type
	if typeNum > 2 || typeNum < 1 {
		response.FailWithMessage("type只能是 1 或者 2", c)
		return
	}
	if typeNum == 1 {
		if err := utils.Verify(h, utils.MinAppHelpIncidentallyTypeOne); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	} else {
		if err := utils.Verify(h, utils.MinAppHelpIncidentallyTypeTwo); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}

	if h.FetchAddress == h.ReceiveAddress {
		response.FailWithMessage("地址不能相同", c)
		return
	}
	uid := c.MustGet("id").(uint)
	userNum := c.MustGet("userNum").(string)
	h.Uid = uid
	h.UserNum = userNum
	h.OrderNum = "INC" + utils.UserNum()
	if err := helpIncidentallyService.CreateHelpIncidentally(h); err != nil {
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

//获取国内捎带订单
//id 订单的ID
//address 搜索的内容信息
func (a *HelpApi) GetHelpIncidentally(c *gin.Context) {
	var pageInfo autocodeReq.HelpIncidentallySearch
	_ = c.ShouldBindQuery(&pageInfo)

	if pageInfo.ID > 0 {
		if err, result := helpIncidentallyService.GetHelpIncidentally(pageInfo.ID); err != nil {

			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"result": result}, c)
		}
		return
	}

	if err, list, total := helpIncidentallyService.GetHelpIncidentallyInfoList(pageInfo); err != nil {
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

//创建组团拼车订单
func (a *HelpApi) CreateHelpCarpool(c *gin.Context) {
	var h autocode.HelpCarpool
	_ = c.ShouldBindJSON(&h)

	//判断如果我是乘机人 type = 1
	typeNum := h.Type
	if *typeNum > 2 || *typeNum < 1 {
		response.FailWithMessage("type只能是 1 或者 2", c)
		return
	}
	if *typeNum == 1 {
		if err := utils.Verify(h, utils.MinAppHelpCarpoolTypeOne); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	} else {
		if err := utils.Verify(h, utils.MinAppHelpCarpoolTypeTwo); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}
	if h.StartAddress == h.ArriveAddress {
		response.FailWithMessage("地址不能相同", c)
		return
	}

	uid := c.MustGet("id").(uint)
	userNum := c.MustGet("userNum").(string)
	h.Uid = uid
	h.UserNum = userNum
	h.OrderNum = "CAR" + utils.UserNum()
	if err := helpCarpoolService.CreateHelpCarpool(h); err != nil {
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

//获取组团拼车
//id 订单的ID
//address 搜索的内容信息
func (a *HelpApi) GetHelpCarpool(c *gin.Context) {
	var pageInfo autocodeReq.HelpCarpoolSearch
	_ = c.ShouldBindQuery(&pageInfo)

	if pageInfo.ID > 0 {
		if err, result := helpCarpoolService.GetHelpCarpool(pageInfo.ID); err != nil {

			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"carpool": result}, c)
		}
		return
	}

	if err, list, total := helpCarpoolService.GetHelpCarpoolInfoList(pageInfo); err != nil {
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

//团购接龙，发布商品页
func (a *HelpApi) CreateHelpGroupon(c *gin.Context) {
	var h autocode.HelpGroupon
	_ = c.ShouldBindJSON(&h)

	if err := utils.Verify(h, utils.MinAppHelpGroupon); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := c.MustGet("id").(uint)
	userNum := c.MustGet("userNum").(string)
	h.Uid = uid
	h.UserNum = userNum
	h.OrderNum = "GRO" + utils.UserNum()
	if err := helpGrouponService.CreateHelpGroupon(h); err != nil {
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

//团购接龙，报名占位
func (a *HelpApi) ApplyHelpGroupon(c *gin.Context) {
	var h autocode.HelpGroupon
	_ = c.ShouldBindJSON(&h)

	if h.GoodsTotal < 1 {
		response.FailWithMessage("请填写商品数量", c)
		return
	}

	if h.ID < 1 {
		response.FailWithMessage("请携带商品ID", c)
	} else {
		uid := c.MustGet("id").(uint)
		if err := helpService.ApplyHelpGroupon(uid, h.ID, h.GoodsTotal); err != nil {
			//增加查看 人数
			helpService.UpdateGrouponLook(h.ID)
			response.FailWithMessage(err.Error(), c)
		} else {
			//增加跟团人数
			helpService.UpdateGrouponFollow(h.ID)

			//增加查看人数
			helpService.UpdateGrouponLook(h.ID)

			response.OkWithMessage("占位成功", c)
		}
	}
}

//获取组团购接龙商品
//id 订单的ID

func (a *HelpApi) GetHelpGroupon(c *gin.Context) {
	var pageInfo autocodeReq.HelpGrouponSearch
	_ = c.ShouldBindQuery(&pageInfo)

	if pageInfo.ID > 0 {
		if err, result := helpGrouponService.GetHelpGroupon(pageInfo.ID); err != nil {

			response.FailWithMessage("查询失败", c)
		} else {
			helpService.UpdateGrouponLook(pageInfo.ID)
			response.OkWithData(gin.H{"groupon": result}, c)
		}
		return
	}

	//获取我的团购接龙商品
	//pageInfo.Status = 1
	if err, list, total := helpGrouponService.GetMinAppHelpGrouponInfoList(pageInfo); err != nil {
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

//创建二手闲置商品发布
func (a *HelpApi) CreateHelpSecond(c *gin.Context) {
	var h autocode.HelpSecond
	_ = c.ShouldBindJSON(&h)

	if err := utils.Verify(h, utils.MinAppHelpSecond); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := c.MustGet("id").(uint)
	userNum := c.MustGet("userNum").(string)
	h.Uid = uid
	h.UserNum = userNum
	h.OrderNum = "SEC" + utils.UserNum()
	if err := helpSecondService.CreateHelpSecond(h); err != nil {
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

//二手闲置收藏
func (a *HelpApi) Collect(c *gin.Context) {
	var cq minAppRequest.CollectRequest
	_ = c.ShouldBindJSON(&cq)
	orderId := cq.OrderId
	if orderId < 1 {
		response.FailWithMessage("请填写正确的订单id", c)
		return
	}
	uid := c.MustGet("id").(uint)
	//如果等于create 创建收藏
	if cq.Way == "create" {
		if err := helpService.CreateCollect(uid, orderId); err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithMessage("收藏成功", c)
		}
	} else if cq.Way == "cancel" {
		if err := helpService.CancelCollect(uid, orderId); err != nil {
			response.FailWithMessage("取消失败", c)
		} else {
			response.OkWithMessage("取消成功", c)
		}
	}
}

//二手闲置评论留言
func (a *HelpApi) SecondMessage(c *gin.Context) {
	var message minapp.SecondMessage
	_ = c.ShouldBindJSON(&message)
	if err := utils.Verify(message, utils.SecondMessageVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := c.MustGet("id").(uint)
	message.Uid = uid
	message.MessageTime = utils.Today()
	if err := helpService.CreateSecondMessage(message); err != nil {
		response.FailWithMessage("失败", c)
	} else {
		//发送未读消息数量到首页
		_ = UnreadMessageCount([]string{strconv.Itoa(int(message.Tid))})

		response.OkWithMessage("成功", c)
	}
}

//获取二手闲置订单的所有留言
//订单ID
func (a *HelpApi) GetSecondOrderMessage(c *gin.Context) {
	var cq minAppRequest.CollectRequest
	_ = c.ShouldBindQuery(&cq)
	orderNum := cq.OrderNum
	if len(orderNum) == 0 {
		response.FailWithMessage("请填写正确的订单号", c)
		return
	}
	if err, result := helpService.GetSecondOrderMessage(orderNum); err != nil {
		response.FailWithMessage("失败", c)
		return
	} else {
		response.OkWithData(gin.H{"list": result}, c)
	}
}

//获取二手闲置商品
//id 订单的ID
func (a *HelpApi) GetHelpSecond(c *gin.Context) {
	var pageInfo autocodeReq.HelpSecondSearch
	_ = c.ShouldBindQuery(&pageInfo)
	uid := c.MustGet("id").(uint)
	if pageInfo.ID > 0 {
		if err, result := helpSecondService.GetHelpSecondStatus(uid, pageInfo.ID); err != nil {

			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"second": result}, c)
		}
		return
	}

	if err, list, total := helpSecondService.GetHelpSecondInfoList(pageInfo); err != nil {
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

//获取二手闲置所有的分类信息
func (a *HelpApi) GetCategoryList(c *gin.Context) {
	if err, result := helpSecondCategoryService.GetAllCategory(); err != nil {
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"list": result}, c)
	}
}

//获取团购商品的分类信息
func (a *HelpApi) GetGroupByCategory(c *gin.Context) {
	if err, result := helpSecondCategoryService.GetGroupByCategory(); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": result}, c)
	}
}
