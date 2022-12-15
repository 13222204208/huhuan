// 自动生成模板HelpGroupon
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

//团购接龙
// HelpGroupon 结构体
// 如果含有time.Time 请自行import time包
type HelpGroupon struct {
	global.GVA_MODEL
	OrderNum      string                `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:订单号;size:18;"`
	Title         string                `json:"title" form:"title" gorm:"column:title;comment:团购标题;"`
	GoodsName     string                `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名称;"`
	GoodsTotal    int                   `json:"goodsTotal" form:"goodsTotal" gorm:"column:goods_total;comment:商品数量;"`
	GoodsUnit     string                `json:"goodsUnit" form:"goodsUnit" gorm:"column:goods_unit;comment:商品的单位"`
	GrouponTotal  int                   `json:"grouponTotal" form:"grouponTotal" gorm:"column:groupon_total;default:0;comment:已团购占位的数量"`
	Price         float64               `json:"price" form:"price" gorm:"column:price;comment:商品价格;type:decimal(10,2);default:0"`
	StartTime     string                `json:"startTime" form:"startTime" gorm:"column:start_time;comment:组团的开始时间"`
	OverTime      string                `json:"overTime" form:"overTime" gorm:"column:over_time;comment:组团的结束时间"`
	Time          string                `json:"time" form:"time" gorm:"column:time;comment:团购时长;"`
	GoodsImg      string                `json:"goodsImg" form:"goodsImg" gorm:"column:goods_img;comment:商品图片;type:text;"`
	Look          int                   `json:"look" form:"look" gorm:"column:look;comment:查看的次数;default:0"`
	Follow        int                   `json:"follow" form:"follow" gorm:"column:follow;comment:跟团的人数;default:0"`
	GoodsDetail   string                `json:"goodsDetail" form:"goodsDetail" gorm:"column:goods_detail;comment:商品详情;type:text;"`
	Remark        string                `json:"remark" form:"remark" gorm:"column:remark;comment:留言;"`
	TakeName      string                `json:"takeName" form:"takeName" gorm:"column:take_name;comment:接单人姓名;"`
	TakePhone     string                `json:"takePhone" form:"takePhone" gorm:"column:take_phone;comment:接单人电话;"`
	TakeTime      string                `json:"takeTime" form:"takeTime" gorm:"column:take_time;comment:接单时间;"`
	Uid           uint                  `json:"uid" form:"uid" gorm:"column:uid;comment:发布人的ID"`
	UserNum       string                `json:"userNum" form:"userNum" gorm:"column:user_num;comment:发布人用户编号;"`
	Status        int                   `json:"status" form:"status" gorm:"column:status;comment:订单的状态;size:2;default:1"`
	PutAway       int                   `json:"putAway" form:"putAway" gorm:"column:putAway;default:1;comment:上架的状态 ，1 为上架，2 为下架"`
	City          string                `json:"city" form:"city" gorm:"column:city;comment:所属城市;"`
	StopTime      string                `json:"stopTime" form:"stopTime" gorm:"column:stop_time;comment:截止时间"`
	Area          string                `json:"area" form:"area" gorm:"column:area;comment:订单所属地区;"`
	Way           string                `json:"way" form:"way" gorm:"-"`
	UserInfo      MinappUser            `gorm:"foreignKey:Uid"`
	OnlineMessage []SecondOnlineMessage `gorm:"foreignKey:OrderNum"`
}
