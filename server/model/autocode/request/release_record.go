package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ReleaseRecordSearch struct {
	autocode.ReleaseRecord
	request.PageInfo
}

type SendReleaseRecord struct {
	Uid      []int `json:"uid"`
	CouponId uint  `json:"couponId" form:"couponId"`
	Total    int   `json:"total" form:"total"`
	Status   int   `json:"status" form:"status"`
}
