package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MinappUserApi struct {
}

var minappUserService = service.ServiceGroupApp.AutoCodeServiceGroup.MinappUserService
var recordService = service.ServiceGroupApp.MinAppServiceGroup.RecordService
var userService = service.ServiceGroupApp.MinAppServiceGroup.UserService

// CreateMinappUser 创建MinappUser
// @Tags MinappUser
// @Summary 创建MinappUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.MinappUser true "创建MinappUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /minappUser/createMinappUser [post]
func (minappUserApi *MinappUserApi) CreateMinappUser(c *gin.Context) {
	var minappUser autocode.MinappUser
	_ = c.ShouldBindJSON(&minappUser)
	if err := minappUserService.CreateMinappUser(minappUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMinappUser 删除MinappUser
// @Tags MinappUser
// @Summary 删除MinappUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.MinappUser true "删除MinappUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /minappUser/deleteMinappUser [delete]
func (minappUserApi *MinappUserApi) DeleteMinappUser(c *gin.Context) {
	var minappUser autocode.MinappUser
	_ = c.ShouldBindJSON(&minappUser)
	if err := minappUserService.DeleteMinappUser(minappUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMinappUserByIds 批量删除MinappUser
// @Tags MinappUser
// @Summary 批量删除MinappUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MinappUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /minappUser/deleteMinappUserByIds [delete]
func (minappUserApi *MinappUserApi) DeleteMinappUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := minappUserService.DeleteMinappUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMinappUser 更新MinappUser
// @Tags MinappUser
// @Summary 更新MinappUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.MinappUser true "更新MinappUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /minappUser/updateMinappUser [put]
func (minappUserApi *MinappUserApi) UpdateMinappUser(c *gin.Context) {
	//var minappUser autocode.MinappUser
	//_ = c.ShouldBindJSON(&minappUser)
	//if err := minappUserService.UpdateMinappUser(minappUser); err != nil {
	//    global.GVA_LOG.Error("更新失败!", zap.Error(err))
	//	response.FailWithMessage("更新失败", c)
	//} else {
	//	response.OkWithMessage("更新成功", c)
	//}
	var record minapp.Record
	_ = c.ShouldBindJSON(&record)
	if record.Uid > 0 && record.Money > 0 {
		//record.Type = 1   //1 表示用户充值"
		record.Status = 1 //表示充值成功
		if record.Type == 1 {
			record.Remark = "后台增加金额"
		} else {
			record.Remark = "后台减少金额"
		}
		if err := recordService.CreateUserTopUpRecord(record); err != nil {
			response.FailWithMessage("充值失败", c)
		} else {
			if err := userService.BackgroundTopUpBalance(record.Uid, record.Money, record.Type); err != nil {
				response.FailWithMessage("用户充值失败", c)
			} else {
				response.OkWithMessage("用户充值成功", c)
				return
			}
			response.OkWithMessage("成功", c)
		}
	} else {
		response.FailWithMessage("参数错误", c)
	}
}

// FindMinappUser 用id查询MinappUser
// @Tags MinappUser
// @Summary 用id查询MinappUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.MinappUser true "用id查询MinappUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /minappUser/findMinappUser [get]
func (minappUserApi *MinappUserApi) FindMinappUser(c *gin.Context) {
	var minappUser autocode.MinappUser
	_ = c.ShouldBindQuery(&minappUser)
	if err, reminappUser := minappUserService.GetMinappUser(minappUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reminappUser": reminappUser}, c)
	}
}

// GetMinappUserList 分页获取MinappUser列表
// @Tags MinappUser
// @Summary 分页获取MinappUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.MinappUserSearch true "分页获取MinappUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /minappUser/getMinappUserList [get]
func (minappUserApi *MinappUserApi) GetMinappUserList(c *gin.Context) {
	var pageInfo autocodeReq.MinappUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := minappUserService.GetMinappUserInfoList(pageInfo); err != nil {
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

//获取新增用户的数据
func (minappUserApi *MinappUserApi) GetNewUserCount(c *gin.Context) {
	var res autocodeReq.NewMinAppUserEveryday
	_ = c.ShouldBindQuery(&res)

	if res.Way == "userTop" {
		if err, result := minappUserService.UserTop(); err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithData(gin.H{"data": result}, c)
		}
		return
	}
	//统计每月用户访问次数
	if res.Way == "monthly" {
		if err, result := redisService.GetMonthlyUsers(); err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithData(gin.H{"data": result}, c)
		}
		return
	}

	if err, result := redisService.GetUserCount(res.Start, res.End); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"data": result}, c)
	}

}
