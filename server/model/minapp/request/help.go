package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HelpRepairSearch struct {
	autocode.HelpRepair
	request.PageInfo
}

//二手闲置的收藏与取消收藏
type CollectRequest struct {
	Way      string `json:"way" form:"way"`
	OrderNum string `json:"orderNum" form:"orderNum"`
	OrderId  uint   `json:"orderId" form:"orderId"`
}
