// 自动生成模板HelpSecond
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HelpSecond 结构体
// 如果含有time.Time 请自行import time包
type HelpSecond struct {
	global.GVA_MODEL
	OrderNum      string   `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:订单号;size:18;"`
	GoodsDetail   string   `json:"goodsDetail" form:"goodsDetail" gorm:"column:goods_detail;comment:商品详情;type:text;"`
	GoodsImg      string   `json:"goodsImg" form:"goodsImg" gorm:"column:goods_img;comment:商品图片;type:text;"`
	GoodsName     string   `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名称;"`
	OriginalPrice *float64 `json:"originalPrice" form:"originalPrice" gorm:"column:original_price;comment:商品原价;"`
	Price         float64  `json:"price" form:"price" gorm:"column:price;type:decimal(10,2);comment:商品价格;"`
	Address       string   `json:"address" form:"address" gorm:"column:address;comment:我的地址;"`
	TakeName      string   `json:"takeName" form:"takeName" gorm:"column:take_name;comment:接单人姓名;"`
	TakePhone     string   `json:"takePhone" form:"takePhone" gorm:"column:take_phone;comment:接单人电话;"`
	TakeTime      string   `json:"takeTime" form:"takeTime" gorm:"column:take_time;comment:接单时间;"`
	StopTime      string   `json:"stopTime" form:"stopTime" gorm:"column:stop_time;comment:截止时间"`
	Uid           uint     `json:"uid" form:"uid" gorm:"column:uid;comment:发布人的ID"`
	UserNum       string   `json:"userNum" form:"userNum" gorm:"column:user_num;comment:发布人用户编号;"`
	Status        int      `json:"status" form:"status" gorm:"column:status;comment:订单的状态;size:2;default:1"`
	PutAway       int      `json:"putAway" form:"putAway" gorm:"column:putAway;default:1;comment:上架的状态 ，1 为上架，2 为下架"`
	//收藏状态
	Collect int `json:"collect" gorm:"-"`
	//是否已经为我想要状态
	Take int `json:"take" gorm:"-"`

	City     string `json:"city" form:"city" gorm:"column:city;comment:所属城市;"`
	Area     string `json:"area" form:"area" gorm:"column:area;comment:订单所属地区;"`
	Category string `json:"category" form:"category" gorm:"column:category;comment:分类商品的类别;size:5000;"`
	Way      string `json:"way" form:"way" gorm:"-"`
	//一对多关系，这个订单的消息
	OnlineMessage []SecondOnlineMessage `gorm:"foreignKey:OrderNum"`
	//SecondMessage []minapp.SecondMessage `gorm:"foreignKey:orderId"`
	UserInfo MinappUser `gorm:"foreignKey:Uid"`
}

type TakeHelpSecond struct {
	global.GVA_MODEL
	Uid        uint       `json:"uid" form:"uid" gorm:"column:uid;"`
	OrderId    uint       `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单id"`
	HelpSecond HelpSecond `gorm:"foreignKey:OrderId"`
}
