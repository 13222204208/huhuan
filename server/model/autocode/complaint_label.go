// 自动生成模板ComplaintLabel
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ComplaintLabel 结构体
// 如果含有time.Time 请自行import time包
type ComplaintLabel struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标签名称;"`
}


