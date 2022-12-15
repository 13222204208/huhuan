// 自动生成模板HelpCarpool
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HelpCarpool 结构体
// 如果含有time.Time 请自行import time包
type HelpCarpool struct {
	global.GVA_MODEL
	OrderNum     string   `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:订单号;size:18;"`
	OwnerName    string   `json:"ownerName" form:"ownerName" gorm:"column:owner_name;comment:车主姓名;"`
	RiderName    string   `json:"riderName" form:"riderName" gorm:"column:rider_name;comment:乘车人姓名;"`
	Gender       string   `json:"gender" form:"gender" gorm:"column:gender;comment:姓别;size:4;"`
	ContactPhone string   `json:"contactPhone" form:"contactPhone" gorm:"column:contact_phone;comment:联系电话;"`
	Time         string   `json:"time" form:"time" gorm:"column:time;comment:接单时间;"`
	Remark       string   `json:"remark" form:"remark" gorm:"column:remark;comment:留言;size:6000;"`
	CouponId     *int     `json:"couponId" form:"couponId" gorm:"column:coupon_id;comment:优惠券ID;"`
	CouponAmount *float64 `json:"couponAmount" form:"couponAmount" gorm:"column:coupon_amount;comment:优惠券优惠金额;"`
	Price        float64  `json:"price" form:"price" gorm:"column:price;comment:费用;type:decimal(10,2);default:0"`
	TakeName     string   `json:"takeName" form:"takeName" gorm:"column:take_name;comment:接单人姓名;"`
	TakePhone    string   `json:"takePhone" form:"takePhone" gorm:"column:take_phone;comment:接单人电话;"`
	TakeId       uint     `json:"takeId" form:"takeId" gorm:"column:take_id;comment:接单人ID"`
	Uid          uint     `json:"uid" form:"uid" gorm:"column:uid;comment:发布人的ID"`
	TakeTime     string   `json:"takeTime" form:"takeTime" gorm:"column:take_time;comment:接单时间;"`
	UserNum      string   `json:"userNum" form:"userNum" gorm:"column:user_num;comment:发布人用户编号;"`
	StopTime     string   `json:"stopTime" form:"stopTime" gorm:"column:stop_time;comment:截止时间"`
	//RiderNum      int        `json:"riderNum" form:"riderNum" gorm:"column:rider_num;comment:乘车人数;size:2;"`
	Status         int                   `json:"status" form:"status" gorm:"column:status;comment:订单的状态;size:2;default:1"`
	PayStatus      int                   `json:"payStatus" form:"payStatus" gorm:"column:pay_status;comment:订单支付的状态;default:0"`
	DoneTime       string                `json:"doneTime" form:"doneTime" gorm:"column:done_time;comment:订单完成时间;"`
	City           string                `json:"city" form:"city" gorm:"column:city;comment:所属城市;"`
	Area           string                `json:"area" form:"area" gorm:"column:area;comment:订单所属地区;"`
	Type           *int                  `json:"type" form:"type" gorm:"column:type;comment:1我是车主，2我是乘车人;size:2;"`
	StartAddress   string                `json:"startAddress" form:"startAddress" gorm:"column:start_address;comment:出发地址;"`
	ArriveAddress  string                `json:"arriveAddress" form:"arriveAddress" gorm:"column:arrive_address;comment:送达地址;"`
	CarNumber      string                `json:"carNumber" form:"carNumber" gorm:"column:car_number;comment:车牌号;"`
	CarType        string                `json:"carType" form:"carType" gorm:"column:car_type;comment:车型;"`
	Brand          string                `json:"brand" form:"brand" gorm:"column:brand;comment:车的品牌"`
	UseNum         int                   `json:"useNum" form:"useNum" gorm:"column:use_num;comment:用车人数;size:2;"`
	UseTime        string                `json:"useTime" form:"useTime" gorm:"column:use_time;comment:用车时间;"`
	CarUsedNum     int                   `json:"carUsedNum" form:"carUsedNum" gorm:"column:car_used_num;comment:车辆已预约人数，已经坐了几人;size:2;default:0"`
	Way            string                `json:"way" form:"way" gorm:"-"`
	CouponInfo     Coupon                `json:"couponInfo" form:"couponInfo" gorm:"foreignKey:CouponId"`
	Images         string                `json:"images" form:"images" gorm:"column:images;type:text;comment:留言下的图片"`
	UserInfo       MinappUser            `gorm:"foreignKey:Uid"`
	OnlineMessage  []SecondOnlineMessage `gorm:"foreignKey:OrderNum"`
	CarUsedNumInfo []CarUsedNum          `json:"carUsedNumAvatar" gorm:"foreignKey:OrderId"`
}
