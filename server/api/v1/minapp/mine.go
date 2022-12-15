package minapp

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MineApi struct{}

//我的二手闲置
func (a *MineApi) MyHelpSecond(c *gin.Context) {
	var pageInfo autocodeReq.HelpSecondSearch
	_ = c.ShouldBindQuery(&pageInfo)

	uid := c.MustGet("id").(uint)
	pageInfo.Uid = uid
	//如果获取的是我收藏的二手闲置列表
	if pageInfo.Way == "collect" {
		if err, list, total := mineService.MyCollectHelpSecond(pageInfo); err != nil {
			response.FailWithMessage("获取失败", c)
			return
		} else {
			response.OkWithDetailed(response.PageResult{
				List:     list,
				Total:    total,
				Page:     pageInfo.Page,
				PageSize: pageInfo.PageSize,
			}, "获取成功", c)
			return
		}
	} else if pageInfo.Way == "take" {
		if err, list, total := mineService.MyTakeHelpSecond(pageInfo); err != nil {
			response.FailWithMessage("获取失败", c)
			return
		} else {
			response.OkWithDetailed(response.PageResult{
				List:     list,
				Total:    total,
				Page:     pageInfo.Page,
				PageSize: pageInfo.PageSize,
			}, "获取成功", c)
			return
		}
	}

	if err, list, total := mineService.MyHelpSecondInfoList(pageInfo); err != nil {
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

//我的组团拼车
//id 订单的ID
//address 搜索的内容信息
func (a *MineApi) MyHelpCarpool(c *gin.Context) {
	var pageInfo autocodeReq.HelpCarpoolPageInfo
	_ = c.ShouldBindQuery(&pageInfo)

	uid := c.MustGet("id").(uint)
	pageInfo.Uid = uid

	if pageInfo.ID > 0 {
		if err, result := helpCarpoolService.GetHelpCarpool(pageInfo.ID); err != nil {

			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"carpool": result}, c)
		}
		return
	}

	//获取我预约的组团拼车订单
	fmt.Println("信息", pageInfo.PageSize, pageInfo.Page)
	if pageInfo.Way == "take" {
		if err, list, total := mineService.MyHelpCarpoolTakeList(pageInfo); err != nil {
			response.FailWithMessage("获取失败", c)
			return
		} else {
			response.OkWithDetailed(response.PageResult{
				List:     list,
				Total:    total,
				Page:     pageInfo.Page,
				PageSize: pageInfo.PageSize,
			}, "获取成功", c)
			return
		}
	}

	if err, list, total := mineService.MyHelpCarpoolList(pageInfo); err != nil {
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

//我的二手闲置
func (a *MineApi) SetMyHelpSecond(c *gin.Context) {
	var second autocode.HelpSecond
	_ = c.ShouldBindJSON(&second)
	if second.ID < 1 {
		response.FailWithMessage("请填写订单ID", c)
		return
	}
	way := second.Way
	uid := c.MustGet("id").(uint)
	switch way {
	case "update":
		second.Uid = uid

		if err := mineService.UpdateHelpSecond(second); err != nil {
			response.FailWithMessage("更新失败", c)
		} else {
			response.OkWithMessage("更新成功", c)
		}
	case "cancel":
		second.Uid = uid
		second.Status = 6
		if err := mineService.UpdateHelpSecond(second); err != nil {
			response.FailWithMessage("取消失败", c)
		} else {
			response.OkWithMessage("取消成功", c)
		}
	case "take":
		if err := mineService.TakeHelpSecond(uid, second.ID); err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			if err, h := helpSecondService.GetHelpSecond(second.ID); err != nil {
				fmt.Println(err.Error())
			} else {
				//接单成功后发送一条我想要这件商品的私信消息
				var room minapp.SecondRoom
				room.OrderNum = h.OrderNum
				room.Sell = h.Uid
				room.Buy = uid
				_, roomId := messageService.SaveSecondRoom(room)
				if err := messageService.SaveSecondOnlineMessage(uid, h.Uid, strconv.Itoa(int(roomId)), "您好，我想要这件商品", h.OrderNum); err != nil {
					fmt.Println(err.Error())
				}
			}

			response.OkWithMessage("接单成功", c)
		}

	default:
		response.FailWithMessage("参数错误", c)
	}

}

//乘客发布
//取消或编辑组团拼车
func (a *MineApi) SetMyHelpCarpool(c *gin.Context) {
	var carpool autocode.HelpCarpool
	_ = c.ShouldBindJSON(&carpool)

	if carpool.ID < 1 {
		response.FailWithMessage("请填写订单ID", c)
		return
	}
	//判断是请求的类型是否为取消订单
	way := carpool.Way
	uid := c.MustGet("id").(uint)
	switch way {
	case "cancel":
		if err := mineService.CancelMyHelpCarpool(carpool.ID, uid); err != nil {
			response.FailWithMessage("取消失败", c)
		} else {
			response.OkWithMessage("取消成功", c)
			return
		}
	case "update":
		carpool.Uid = uid

		if err := mineService.UpdateHelpCarpool(carpool); err != nil {
			response.FailWithMessage("更新失败", c)
		} else {
			response.OkWithMessage("更新成功", c)
		}
	case "add":
		price := carpool.Price //加价金额
		if price < 1 {
			response.FailWithMessage("请填写加价金额", c)
			return
		}
		//我发布的组团拼车
		if err := mineService.AddPriceHelpCarpool(carpool.ID, price); err != nil {
			response.FailWithMessage("加价失败", c)
		} else {
			response.OkWithData("加价成功", c)
		}
	case "take":
		//组团拼车接单
		carpool.TakeId = uid
		if err := mineService.TakeHelpCarpool(carpool.ID, uid); err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithMessage("接单成功", c)
		}
	case "finish":
		//确认完成订单
		carpool.Status = 4
		if err := mineService.UpdateHelpCarpool(carpool); err != nil {
			response.FailWithMessage("订单完成失败", c)
		} else {
			response.OkWithMessage("订单完成", c)
		}
	case "reward":
		price := carpool.Price
		if price < 1 {
			response.FailWithMessage("请填写打赏金额", c)
			return
		}
		var reward minapp.Reward
		uid := c.MustGet("id").(uint)
		reward.Uid = uid
		reward.Money = price
		reward.OrderId = carpool.ID
		reward.OrderType = "helpCarpool"
		createIP := c.ClientIP()
		openid := c.MustGet("openid").(string)
		if err, config := mineService.RewardOrder(openid, createIP, reward); err != nil {
			response.FailWithMessage("失败", c)
		} else {
			response.OkWithData(gin.H{"params": config}, c)
		}
	default:
		response.FailWithMessage("参数错误", c)
	}

}

//我的帮我取订单
func (a *MineApi) MyHelpFetch(c *gin.Context) {
	var pageInfo autocodeReq.HelpFetchSearch
	_ = c.ShouldBindQuery(&pageInfo)

	uid := c.MustGet("id").(uint)
	pageInfo.Uid = uid
	if err, list, total := mineService.MyHelpFetchInfoList(pageInfo); err != nil {
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

//取消编辑加价，接单 帮我取订单
func (a *MineApi) SetMyHelpFetch(c *gin.Context) {
	var helpFetch autocode.HelpFetch
	_ = c.ShouldBindJSON(&helpFetch)
	way := helpFetch.Way
	if helpFetch.ID < 1 {
		response.FailWithMessage("请填写订单ID", c)
		return
	}
	uid := c.MustGet("id").(uint)
	switch way {
	case "take":
		helpFetch.TakeId = uid

		if err, result := mineService.FindUserVerifying(uid); err != nil {
			response.FailWithMessage("失败", c)
			return
		} else {
			if result == 0 {
				response.FailWithMessage("请您填写身份验证信息", c)
				return
			}
		}

		helpFetch.Status = 2
		if err := mineService.UpdateHelpFetch(helpFetch); err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithMessage("接单成功", c)
		}
	case "cancel":
		helpFetch.Status = 6
		if err := mineService.UpdateHelpFetch(helpFetch); err != nil {
			response.FailWithMessage("取消失败", c)
		} else {
			response.OkWithMessage("取消成功", c)
		}

	case "cancelTake":
		if err := global.GVA_DB.Model(&autocode.HelpFetch{}).Where("id", helpFetch.ID).Updates(map[string]interface{}{
			"take_id": 0,
			"status":  1,
		}).Error; err != nil {
			response.FailWithMessage("取消失败", c)
		} else {
			response.OkWithMessage("取消成功", c)
		}

	case "update":

		//判断修改时，是否修改了订单的金额
		var h autocode.HelpFetch
		global.GVA_DB.Where("id", helpFetch.ID).First(&h)
		price := utils.Decimal(helpFetch.Price)
		buyPrice := h.Price
		if price != buyPrice {
			if price < buyPrice {
				money := buyPrice - price
				mineService.RefundToUser(uid, h.OrderNum, money)
			} else {
				response.FailWithMessage("请使用加价", c)
				return
			}
		}

		if err := mineService.UpdateHelpFetch(helpFetch); err != nil {
			response.FailWithMessage("更新失败", c)
		} else {
			response.OkWithMessage("更新成功", c)
		}
	case "finish":
		helpFetch.Status = 3
		if err := mineService.UpdateHelpFetch(helpFetch); err != nil {
			response.FailWithMessage("失败", c)
		} else {
			response.OkWithMessage("成功", c)
		}
	case "add":
		price := helpFetch.Price
		if price < 1 {
			response.FailWithMessage("请填写加价金额", c)
			return
		}
		fmt.Println("加价订单ID", helpFetch.ID)
		if err := mineService.AddPriceHelpFetch(helpFetch.ID, price); err != nil {
			response.FailWithMessage("加价失败", c)
		} else {

			if helpFetch.PayType == "balancePay" {
				//余额支付的流程
				if err := payService.BalancePayMethod(helpFetch.OrderNum, price, uid); err != nil {
					response.FailWithMessage(err.Error(), c)

				} else {
					response.OkWithMessage("加价成功", c)
				}
				return
			}

			var reward minapp.Reward
			uid := c.MustGet("id").(uint)
			reward.Uid = uid
			reward.Money = price
			reward.OrderId = helpFetch.ID
			reward.OrderType = "helpFetch"
			createIP := c.ClientIP()
			openid := c.MustGet("openid").(string)
			if err, config := mineService.AddPriceOrder(openid, createIP, reward); err != nil {
				response.FailWithMessage("失败", c)
			} else {
				response.OkWithData(gin.H{"params": config}, c)
			}
		}
	case "reward":
		price := helpFetch.Price
		if price < 1 {
			response.FailWithMessage("请填写打赏金额", c)
			return
		}

		var reward minapp.Reward
		uid := c.MustGet("id").(uint)
		reward.Uid = uid
		reward.Money = price
		reward.OrderId = helpFetch.ID
		reward.OrderType = "helpFetch"
		createIP := c.ClientIP()
		openid := c.MustGet("openid").(string)

		//获取帮我取订单信息
		err, r := helpFetchService.GetHelpFetch(helpFetch.ID)
		if err != nil {
			fmt.Println(err)
		} else {
			if uid == r.TakeId {
				response.FailWithMessage("无法以打赏给自己", c)
				return
			}
			reward.FromOrderNum = r.OrderNum
			reward.Tid = r.TakeId
		}

		if err, config := mineService.RewardOrder(openid, createIP, reward); err != nil {
			response.FailWithMessage("失败", c)
		} else {
			response.OkWithData(gin.H{"params": config}, c)
		}
	default:
		response.FailWithMessage("参数不正确", c)
	}

}

//我的帮我买订单
func (a *MineApi) MyHelpBuy(c *gin.Context) {
	var pageInfo autocodeReq.HelpBuySearch
	_ = c.ShouldBindQuery(&pageInfo)

	uid := c.MustGet("id").(uint)
	pageInfo.Uid = uid
	if err, list, total := mineService.MyHelpBuyInfoList(pageInfo); err != nil {
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

//取消编辑加价，接单 帮我买订单
func (a *MineApi) SetMyHelpBuy(c *gin.Context) {
	var helpBuy autocode.HelpBuy
	_ = c.ShouldBindJSON(&helpBuy)
	way := helpBuy.Way
	if helpBuy.ID < 1 {
		response.FailWithMessage("请填写订单ID", c)
		return
	}
	uid := c.MustGet("id").(uint)
	switch way {
	case "take":
		helpBuy.TakeId = uid

		if err, result := mineService.FindUserVerifying(uid); err != nil {
			response.FailWithMessage("失败", c)
			return
		} else {
			if result == 0 {
				response.FailWithMessage("请您填写身份验证信息", c)
				return
			}
		}

		helpBuy.Status = 2
		if err := mineService.UpdateHelpBuy(helpBuy); err != nil {
			response.FailWithMessage("接单失败", c)
		} else {
			response.OkWithMessage("接单成功", c)
		}
	case "cancel":
		helpBuy.Status = 6
		if err := mineService.UpdateHelpBuy(helpBuy); err != nil {
			response.FailWithMessage("取消失败", c)
		} else {
			response.OkWithMessage("取消成功", c)
		}
	case "cancelTake":
		if err := global.GVA_DB.Model(&autocode.HelpBuy{}).Where("id", helpBuy.ID).Updates(map[string]interface{}{
			"take_id": 0,
			"status":  1,
		}).Error; err != nil {
			response.FailWithMessage("取消失败", c)
		} else {
			response.OkWithMessage("取消成功", c)
		}
	case "update":

		//判断修改时，是否修改了订单的金额
		var h autocode.HelpBuy
		global.GVA_DB.Where("id", helpBuy.ID).First(&h)
		price := utils.Decimal(helpBuy.Price)
		buyPrice := h.Price
		if price != buyPrice {
			if price < buyPrice {
				money := buyPrice - price
				mineService.RefundToUser(uid, h.OrderNum, money)
			} else {
				response.FailWithMessage("请使用加价", c)
				return
			}
		}

		if err := mineService.UpdateHelpBuy(helpBuy); err != nil {
			response.FailWithMessage("更新失败", c)
		} else {

			response.OkWithMessage("更新成功", c)
		}
	case "finish":
		helpBuy.Status = 4
		if err := mineService.UpdateHelpBuy(helpBuy); err != nil {
			response.FailWithMessage("失败", c)
		} else {
			response.OkWithMessage("成功", c)
		}
	case "add":
		price := helpBuy.Price
		if price < 1 {
			response.FailWithMessage("请填写加价金额", c)
			return
		}
		if err := mineService.AddPriceHelpBuy(helpBuy.ID, price); err != nil {
			response.FailWithMessage("加价失败", c)
		} else {

			if helpBuy.PayType == "balancePay" {
				//余额支付的流程
				if err := payService.BalancePayMethod(helpBuy.OrderNum, price, uid); err != nil {
					response.FailWithMessage(err.Error(), c)
					return
				} else {
					response.OkWithMessage("加价成功", c)
					return
				}
			}

			var reward minapp.Reward
			uid := c.MustGet("id").(uint)
			reward.Uid = uid
			reward.Money = price
			reward.OrderId = helpBuy.ID
			reward.OrderType = "helpBuy"
			createIP := c.ClientIP()
			openid := c.MustGet("openid").(string)
			if err, config := mineService.AddPriceOrder(openid, createIP, reward); err != nil {
				response.FailWithMessage("失败", c)
			} else {
				response.OkWithData(gin.H{"params": config}, c)
			}
		}
	case "reward":
		price := helpBuy.Price
		if price < 1 {
			response.FailWithMessage("请填写打赏金额", c)
			return
		}
		var reward minapp.Reward
		uid := c.MustGet("id").(uint)
		reward.Uid = uid
		reward.Money = price
		reward.OrderId = helpBuy.ID
		reward.OrderType = "helpBuy"
		createIP := c.ClientIP()
		openid := c.MustGet("openid").(string)

		//获取帮我买订单信息
		err, r := helpBuyService.GetHelpBuy(helpBuy.ID)
		if err != nil {
			fmt.Println(err)
		} else {
			if uid == r.TakeId {
				response.FailWithMessage("无法以打赏给自己", c)
				return
			}
			reward.FromOrderNum = r.OrderNum
			reward.Tid = r.TakeId
		}

		if err, config := mineService.RewardOrder(openid, createIP, reward); err != nil {
			response.FailWithMessage("失败", c)
		} else {
			response.OkWithData(gin.H{"params": config}, c)
		}
	default:
		response.FailWithMessage("参数不正确", c)
	}
}

//我的帮我做
func (a *MineApi) MyHelpWork(c *gin.Context) {
	var pageInfo autocodeReq.HelpWorkSearch
	_ = c.ShouldBindQuery(&pageInfo)

	uid := c.MustGet("id").(uint)
	pageInfo.Uid = uid
	if err, list, total := mineService.MyHelpWorkInfoList(pageInfo); err != nil {
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

//取消编辑加价，接单 帮我做订单
func (a *MineApi) SetMyHelpWork(c *gin.Context) {
	var helpWork autocode.HelpWork
	_ = c.ShouldBindJSON(&helpWork)
	way := helpWork.Way
	if helpWork.ID < 1 {
		response.FailWithMessage("请填写订单ID", c)
		return
	}
	uid := c.MustGet("id").(uint)
	switch way {
	case "take":
		helpWork.TakeId = uid
		if err, result := mineService.FindUserVerifying(uid); err != nil {
			response.FailWithMessage("失败", c)
			return
		} else {
			if result == 0 {
				response.FailWithMessage("请您填写身份验证信息", c)
				return
			}
		}
		helpWork.Status = 2
		if err := mineService.UpdateHelpWork(helpWork); err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithMessage("接单成功", c)
		}

	case "cancel":
		helpWork.Status = 6
		if err := mineService.UpdateHelpWork(helpWork); err != nil {
			response.FailWithMessage("取消失败", c)
		} else {
			response.OkWithMessage("取消成功", c)
		}
	case "cancelTake":
		if err := global.GVA_DB.Model(&autocode.HelpWork{}).Where("id", helpWork.ID).Updates(map[string]interface{}{
			"take_id": 0,
			"status":  1,
		}).Error; err != nil {
			response.FailWithMessage("取消失败", c)
		} else {
			response.OkWithMessage("取消成功", c)
		}
	case "update":

		//判断修改时，是否修改了订单的金额
		var h autocode.HelpWork
		global.GVA_DB.Where("id", helpWork.ID).First(&h)
		price := utils.Decimal(helpWork.Price)
		buyPrice := h.Price
		if price != buyPrice {
			if price < buyPrice {
				money := buyPrice - price
				mineService.RefundToUser(uid, h.OrderNum, money)
			} else {
				response.FailWithMessage("请使用加价", c)
				return
			}
		}

		if err := mineService.UpdateHelpWork(helpWork); err != nil {
			response.FailWithMessage("更新失败", c)
		} else {
			response.OkWithMessage("更新成功", c)
		}
	case "finish":
		helpWork.Status = 4
		if err := mineService.UpdateHelpWork(helpWork); err != nil {
			response.FailWithMessage("失败", c)
		} else {
			response.OkWithMessage("成功", c)
		}
	case "add":
		price := helpWork.Price
		if price < 1 {
			response.FailWithMessage("请填写加价金额", c)
			return
		}
		if err := mineService.AddPriceHelpWork(helpWork.ID, price); err != nil {
			response.FailWithMessage("加价失败", c)
		} else {
			//判断是否是余额支付
			if helpWork.PayType == "balancePay" {
				//余额支付的流程
				if err := payService.BalancePayMethod(helpWork.OrderNum, price, uid); err != nil {
					response.FailWithMessage(err.Error(), c)
					return
				} else {
					response.OkWithMessage("加价成功", c)
					return
				}
			}

			var reward minapp.Reward
			uid := c.MustGet("id").(uint)
			reward.Uid = uid
			reward.Money = price
			reward.OrderId = helpWork.ID
			reward.OrderType = "helpWork"
			createIP := c.ClientIP()
			openid := c.MustGet("openid").(string)
			if err, config := mineService.AddPriceOrder(openid, createIP, reward); err != nil {
				response.FailWithMessage("失败", c)
			} else {
				response.OkWithData(gin.H{"params": config}, c)
			}
		}
	case "reward":
		price := helpWork.Price
		if price < 1 {
			response.FailWithMessage("请填写打赏金额", c)
			return
		}
		var reward minapp.Reward
		uid := c.MustGet("id").(uint)
		reward.Uid = uid
		reward.Money = price
		reward.OrderId = helpWork.ID
		reward.OrderType = "helpWork"
		createIP := c.ClientIP()
		openid := c.MustGet("openid").(string)

		//获取帮我做订单信息
		err, r := helpWorkService.GetHelpWork(helpWork.ID)
		if err != nil {
			fmt.Println(err)
		} else {
			if uid == r.TakeId {
				response.FailWithMessage("无法以打赏给自己", c)
				return
			}
			reward.FromOrderNum = r.OrderNum
			reward.Tid = r.TakeId
		}

		if err, config := mineService.RewardOrder(openid, createIP, reward); err != nil {
			response.FailWithMessage("失败", c)
		} else {
			response.OkWithData(gin.H{"params": config}, c)
		}
	default:
		response.FailWithMessage("参数不正确", c)
	}
}

//我的帮我修
func (a *MineApi) MyHelpRepair(c *gin.Context) {
	var pageInfo autocodeReq.HelpRepairSearch
	_ = c.ShouldBindQuery(&pageInfo)

	uid := c.MustGet("id").(uint)
	pageInfo.Uid = uid
	if err, list, total := mineService.MyHelpRepairInfoList(pageInfo); err != nil {
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

//取消编辑加价，接单 帮我修订单
func (a *MineApi) SetMyHelpRepair(c *gin.Context) {
	var helpRepair autocode.HelpRepair
	_ = c.ShouldBindJSON(&helpRepair)
	way := helpRepair.Way

	if helpRepair.ID < 1 {
		response.FailWithMessage("请填写订单ID", c)
		return
	}
	uid := c.MustGet("id").(uint)

	switch way {
	case "take":
		helpRepair.TakeId = uid

		if err, result := mineService.FindUserVerifying(uid); err != nil {
			response.FailWithMessage("失败", c)
			return
		} else {
			if result == 0 {
				response.FailWithMessage("请您填写身份验证信息", c)
				return
			}
		}

		helpRepair.Status = 2
		if err := mineService.UpdateHelpRepair(helpRepair); err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithMessage("接单成功", c)
		}
	case "cancel":
		helpRepair.Status = 6
		if err := mineService.UpdateHelpRepair(helpRepair); err != nil {
			response.FailWithMessage("取消失败", c)
		} else {
			response.OkWithMessage("取消成功", c)
		}
	case "cancelTake":
		if err := global.GVA_DB.Model(&autocode.HelpRepair{}).Where("id", helpRepair.ID).Updates(map[string]interface{}{
			"take_id": 0,
			"status":  1,
		}).Error; err != nil {
			response.FailWithMessage("取消失败", c)
		} else {
			response.OkWithMessage("取消成功", c)
		}
	case "update":

		//判断修改时，是否修改了订单的金额
		var h autocode.HelpRepair
		global.GVA_DB.Where("id", helpRepair.ID).First(&h)
		price := utils.Decimal(helpRepair.Price)
		buyPrice := h.Price
		if price != buyPrice {
			if price < buyPrice {
				money := buyPrice - price
				mineService.RefundToUser(uid, h.OrderNum, money)
			} else {
				response.FailWithMessage("请使用加价", c)
				return
			}
		}

		if err := mineService.UpdateHelpRepair(helpRepair); err != nil {
			response.FailWithMessage("更新失败", c)
		} else {
			response.OkWithMessage("更新成功", c)
		}
	case "finish":
		helpRepair.Status = 4
		if err := mineService.UpdateHelpRepair(helpRepair); err != nil {
			response.FailWithMessage("失败", c)
		} else {
			response.OkWithMessage("成功", c)
		}
	case "add":
		price := helpRepair.Price
		if price < 1 {
			response.FailWithMessage("请填写加价金额", c)
			return
		}
		if err := mineService.AddPriceHelpRepair(helpRepair.ID, price); err != nil {
			response.FailWithMessage("加价失败", c)
		} else {
			if helpRepair.PayType == "balancePay" {
				//余额支付的流程
				if err := payService.BalancePayMethod(helpRepair.OrderNum, price, uid); err != nil {
					response.FailWithMessage(err.Error(), c)
					return
				} else {
					response.OkWithMessage("加价成功", c)
					return
				}
			}

			var reward minapp.Reward
			uid := c.MustGet("id").(uint)
			reward.Uid = uid
			reward.Money = price
			reward.OrderId = helpRepair.ID
			reward.OrderType = "helpRepair"
			createIP := c.ClientIP()
			openid := c.MustGet("openid").(string)
			if err, config := mineService.AddPriceOrder(openid, createIP, reward); err != nil {
				response.FailWithMessage("失败", c)
			} else {
				response.OkWithData(gin.H{"params": config}, c)
			}
		}
	case "reward":
		price := helpRepair.Price
		if price < 1 {
			response.FailWithMessage("请填写打赏金额", c)
			return
		}
		var reward minapp.Reward
		uid := c.MustGet("id").(uint)
		reward.Uid = uid
		reward.Money = price
		reward.OrderId = helpRepair.ID
		reward.OrderType = "helpRepair"
		createIP := c.ClientIP()
		openid := c.MustGet("openid").(string)

		//获取帮我修订单信息
		err, r := helpRepairService.GetHelpRepair(helpRepair.ID)
		if err != nil {
			fmt.Println(err)
		} else {
			if uid == r.TakeId {
				response.FailWithMessage("无法以打赏给自己", c)
				return
			}
			reward.FromOrderNum = r.OrderNum
			reward.Tid = r.TakeId
		}

		if err, config := mineService.RewardOrder(openid, createIP, reward); err != nil {
			response.FailWithMessage("失败", c)
		} else {
			response.OkWithData(gin.H{"params": config}, c)
		}
	default:
		response.FailWithMessage("参数不正确", c)
	}
}

//我的帮我修
func (a *MineApi) MyHelpIncidentally(c *gin.Context) {
	var pageInfo autocodeReq.HelpIncidentallySearch
	_ = c.ShouldBindQuery(&pageInfo)

	uid := c.MustGet("id").(uint)
	pageInfo.Uid = uid
	if err, list, total := mineService.MyHelpIncidentallyInfoList(pageInfo); err != nil {
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

//取消编辑加价，接单 国内捎带订单
func (a *MineApi) SetMyHelpIncidentally(c *gin.Context) {
	var helpIncidentally autocode.HelpIncidentally
	_ = c.ShouldBindJSON(&helpIncidentally)
	way := helpIncidentally.Way
	if helpIncidentally.ID < 1 {
		response.FailWithMessage("请填写订单ID", c)
		return
	}
	uid := c.MustGet("id").(uint)
	switch way {
	case "cancel":
		helpIncidentally.Status = 5
		if err := mineService.UpdateHelpIncidentally(helpIncidentally); err != nil {
			response.FailWithMessage("取消失败", c)
		} else {
			response.OkWithMessage("取消成功", c)
		}

	case "update":
		if err := mineService.UpdateHelpIncidentally(helpIncidentally); err != nil {
			response.FailWithMessage("更新失败", c)
		} else {
			response.OkWithMessage("更新成功", c)
		}
	case "finish":
		helpIncidentally.Status = 4
		if err := mineService.UpdateHelpIncidentally(helpIncidentally); err != nil {
			response.FailWithMessage("失败", c)
		} else {
			response.OkWithMessage("成功", c)
		}

	case "take":
		//更改接单人ID
		helpIncidentally.TakeId = uid
		if helpIncidentally.Price > 0 {
			response.FailWithMessage("不能填写金额", c)
			return
		}
		helpIncidentally.Status = 2
		if err := mineService.UpdateHelpIncidentally(helpIncidentally); err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithMessage("接单成功", c)
		}

	case "add":
		price := helpIncidentally.Price
		if price < 1 {
			response.FailWithMessage("请填写加价金额", c)
			return
		}
		if err := mineService.AddPriceHelpIncidentally(helpIncidentally.ID, price); err != nil {
			response.FailWithMessage("加价失败", c)
		} else {
			response.OkWithMessage("加价成功", c)
		}

	case "reward":
		price := helpIncidentally.Price
		if price < 1 {
			response.FailWithMessage("请填写打赏金额", c)
			return
		}
		var reward minapp.Reward
		uid := c.MustGet("id").(uint)
		reward.Uid = uid
		reward.Money = price
		reward.OrderId = helpIncidentally.ID
		reward.OrderType = "helpIncidentally"
		createIP := c.ClientIP()
		openid := c.MustGet("openid").(string)
		if err, config := mineService.RewardOrder(openid, createIP, reward); err != nil {
			response.FailWithMessage("失败", c)
		} else {
			response.OkWithData(gin.H{"params": config}, c)
		}
	default:
		response.FailWithMessage("参数不正确", c)
	}
}

//我的团购接龙
func (a *MineApi) MyHelpGroupon(c *gin.Context) {
	var pageInfo autocodeReq.HelpGrouponSearch
	_ = c.ShouldBindQuery(&pageInfo)

	uid := c.MustGet("id").(uint)
	pageInfo.Uid = uid

	//获取我预约的团购接龙商品
	if pageInfo.Way == "take" {

		if err, list, total := helpGrouponService.GetTakeHelpGrouponInfoList(pageInfo); err != nil {
			response.FailWithMessage("获取失败", c)
			return
		} else {
			response.OkWithDetailed(response.PageResult{
				List:     list,
				Total:    total,
				Page:     pageInfo.Page,
				PageSize: pageInfo.PageSize,
			}, "获取成功", c)
			return
		}
	}
	fmt.Println("请求的数据", pageInfo)
	if err, list, total := mineService.MyHelpGrouponInfoList(pageInfo); err != nil {
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
