package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HelpSecondSearch struct {
	autocode.HelpSecond
	request.PageInfo
	//判断是否是历史订单，如果是填写 1
	History uint `json:"history" form:"history"`
}
