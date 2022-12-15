package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MoneyRecordRouter struct {
}

// InitMoneyRecordRouter 初始化 MoneyRecord 路由信息
func (s *MoneyRecordRouter) InitMoneyRecordRouter(Router *gin.RouterGroup) {
	moneyRecordRouter := Router.Group("moneyRecord").Use(middleware.OperationRecord())
	moneyRecordRouterWithoutRecord := Router.Group("moneyRecord")
	var moneyRecordApi = v1.ApiGroupApp.AutoCodeApiGroup.MoneyRecordApi
	{
		moneyRecordRouter.POST("createMoneyRecord", moneyRecordApi.CreateMoneyRecord)   // 新建MoneyRecord
		moneyRecordRouter.DELETE("deleteMoneyRecord", moneyRecordApi.DeleteMoneyRecord) // 删除MoneyRecord
		moneyRecordRouter.DELETE("deleteMoneyRecordByIds", moneyRecordApi.DeleteMoneyRecordByIds) // 批量删除MoneyRecord
		moneyRecordRouter.PUT("updateMoneyRecord", moneyRecordApi.UpdateMoneyRecord)    // 更新MoneyRecord
	}
	{
		moneyRecordRouterWithoutRecord.GET("findMoneyRecord", moneyRecordApi.FindMoneyRecord)        // 根据ID获取MoneyRecord
		moneyRecordRouterWithoutRecord.GET("getMoneyRecordList", moneyRecordApi.GetMoneyRecordList)  // 获取MoneyRecord列表
	}
}
