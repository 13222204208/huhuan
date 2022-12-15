package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HelpCarpoolSearch struct {
	autocode.HelpCarpool
	request.PageInfo
	//判断是否是历史订单，如果是填写 1
	History uint `json:"history" form:"history"`
}

type HelpCarpoolPageInfo struct {
	request.PageInfo
	Uid uint `json:"uid" form:"uid"`
	ID  uint `json:"ID" form:"ID"`
	//判断是否是历史订单，如果是填写 1
	History uint   `json:"history" form:"history"`
	Way     string `json:"way" form:"way"`
}
