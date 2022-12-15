package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HelpFetchSearch struct {
	autocode.HelpFetch
	request.PageInfo
}

type Order struct {
	OrderNum string `json:"orderNum" form:"orderNum"`
}
