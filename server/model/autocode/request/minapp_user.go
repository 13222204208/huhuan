package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MinappUserSearch struct {
	autocode.MinappUser
	request.PageInfo
}

//获取每日新增用户，日期范围
type NewMinAppUserEveryday struct {
	Start string `json:"start" form:"start"`
	End   string `json:"end" form:"end"`
	Way   string `json:"way" form:"way"` //请的类型，如用户排行 userTop
}
