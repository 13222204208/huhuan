// 自动生成模板HelpIncidentally
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HelpIncidentally 结构体
// 如果含有time.Time 请自行import time包
type HelpIncidentally struct {
	global.GVA_MODEL
	OrderNum       string  `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:订单号;size:18;"`
	GoodsName      string  `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:物品名称;"`
	Weight         string  `json:"weight" form:"weight" gorm:"column:weight;comment:物品重量"`
	GoodsCount     string  `json:"goodsCount" form:"goodsCount" gorm:"column:goods_count;comment:物品数量;"`
	Linkman        string  `json:"linkman" form:"linkman" gorm:"column:linkman;comment:联系人姓名;"`
	ContactPhone   string  `json:"contactPhone" form:"contactPhone" gorm:"column:contact_phone;comment:联系电话;"`
	FetchAddress   string  `json:"fetchAddress" form:"fetchAddress" gorm:"column:fetch_address;comment:取货地址;"`
	ReceiveAddress string  `json:"receiveAddress" form:"receiveAddress" gorm:"column:receive_address;comment:收货地址;"`
	Time           string  `json:"time" form:"time" gorm:"column:time;comment:出发时间;"`
	Remark         string  `json:"remark" form:"remark" gorm:"column:remark;comment:留言;type:text;"`
	Images         string  `json:"images" form:"images" gorm:"column:images;type:text;comment:留言下的图片"`
	CouponId       int     `json:"couponId" form:"couponId" gorm:"column:coupon_id;comment:优惠券ID;"`
	CouponAmount   float64 `json:"couponAmount" form:"couponAmount" gorm:"column:coupon_amount;comment:优惠券优惠金额;"`
	Price          float64 `json:"price" form:"price" gorm:"column:price;comment:赏金;type: decimal(10,2);default: 0"`
	TakeName       string  `json:"takeName" form:"takeName" gorm:"column:take_name;comment:接单人姓名;"`
	TakeId         uint    `json:"takeId" form:"takeId" gorm:"column:take_id;comment:接单人ID"`
	Uid            uint    `json:"uid" form:"uid" gorm:"column:uid;comment:发布人的ID"`
	TakePhone      string  `json:"takePhone" form:"takePhone" gorm:"column:take_phone;comment:接单人电话;"`
	TakeTime       string  `json:"takeTime" form:"takeTime" gorm:"column:take_time;comment:接单时间;"`
	UserNum        string  `json:"userNum" form:"userNum" gorm:"column:user_num;comment:发布人用户编号;"`
	GoodsPrice     float64 `json:"goodsPrice" form:"goodsPrice" gorm:"column:goods_price;comment:商品价值;"`
	StopTime       string  `json:"stopTime" form:"stopTime" gorm:"column:stop_time;comment:截止时间"`
	Status         int     `json:"status" form:"status" gorm:"column:status;comment:订单的状态;size:2;default:1"`
	DoneTime       string  `json:"doneTime" form:"doneTime" gorm:"column:done_time;comment:订单完成时间;"`
	City           string  `json:"city" form:"city" gorm:"column:city;comment:所属城市;"`
	Area           string  `json:"area" form:"area" gorm:"column:area;comment:订单所属地区;"`
	Type           int     `json:"type" form:"type" gorm:"column:type;comment:1寻找乘机人，2我是乘机人;size:2;"`
	StartCity      string  `json:"startCity" form:"startCity" gorm:"column:start_city;comment:出发城市;"`
	ArriveCity     string  `json:"arriveCity" form:"arriveCity" gorm:"column:arrive_city;comment:抵达城市;"`
	FlightNumber   string  `json:"flightNumber" form:"flightNumber" gorm:"column:flight_number;comment:航班号;"`
	Way            string  `json:"way" form:"way" gorm:"-"`
	//一对多关系，这个订单的消息
	OnlineMessage []SecondOnlineMessage `gorm:"foreignKey:OrderNum"`
	CouponInfo    Coupon                `json:"couponInfo" form:"couponInfo" gorm:"foreignKey:CouponId"`
	UserInfo      MinappUser            `gorm:"foreignKey:Uid"`
}
