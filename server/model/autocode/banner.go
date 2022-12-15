// 自动生成模板Banner
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Banner 结构体
// 如果含有time.Time 请自行import time包
type Banner struct {
	global.GVA_MODEL
	Title  string `json:"title" form:"title" gorm:"column:title;comment:轮播图标题;"`
	City   string `json:"city" form:"city" gorm:"column:city;comment:城市"`
	Class  int    `json:"class" form:"class" gorm:"column:class;comment:轮播图分类"`
	Img    string `json:"img" form:"img" gorm:"column:img;comment:轮播图图片;"`
	Url    string `json:"url" form:"url" gorm:"column:url;comment:轮播图跳转的链接;"`
	Hidden bool   `json:"hidden" gorm:"comment:是否在首页隐藏"` // 是否在列表隐藏
}
