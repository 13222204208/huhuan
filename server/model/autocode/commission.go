// 自动生成模板Commission
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Commission 结构体
// 如果含有time.Time 请自行import time包
type Commission struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;"`
      Money  *int `json:"money" form:"money" gorm:"column:money;comment:小于多少金额;"`
      Genre  *int `json:"genre" form:"genre" gorm:"column:genre;comment:佣金的类型;"`
      Commission  *int `json:"commission" form:"commission" gorm:"column:commission;comment:多少佣金;"`
}


