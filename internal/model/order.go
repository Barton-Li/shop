package model

import (
	"shop/internal/model/do"
	"shop/internal/model/entity"
)

type OrderListInput struct {
	Page           int
	Size           int
	Number         string  `json:"number" dc:"订单编号" v:"required#订单编号不能为空"`
	UserId         int     `json:"user_id" dc:"用户id"`
	PayType        int     `json:"pay_type" dc:"支付类型 1 微信 2 支付宝 3 云闪付"`
	PayAtGte       string  `json:"pay_at_gte" dc:"支付时间大于等于"`
	PayAtLte       string  `json:"pay_at_lte" dc:"支付时间小于等于"`
	Status         int     `json:"status" dc:"订单状态1 待支付 2 已支付 3 已发货 4 已收货"`
	ConsigneePhone string  `json:"consignee_phone" dc:"收货人手机号"`
	PriceGte       float64 `json:"price_gte" dc:"订单金额大于等于"`
	PriceLte       float64 `json:"price_lte" dc:"订单金额小于等于"`
	DateGte        string  `json:"date_gte" dc:"下单时间大于等于"`
	DateLte        string  `json:"date_lte" dc:"下单时间小于等于"`
}
type OrderListOutput struct {
	List  []OrderListOutputItem
	Page  int
	Size  int
	Total int
}
type OrderListOutputItem struct {
	entity.OrderInfo
}
type OrderDetailInput struct {
	Id uint
}
type OrderDetailOutput struct {
	do.OrderInfo
	GoodsInfo []*do.OrderGoodsInfo `orm:"with:order_info"`
}
type OrderAddInput struct {
	UserId           uint
	Number           string
	Remark           string `description:"备注"`
	Price            int    `description:"订单金额 单位分"`
	CouponPrice      int    `description:"优惠券金额 单位分"`
	ActualPrice      int    `description:"实际支付金额 单位分"`
	ConsigneeName    string `description:"收货人姓名"`
	ConsigneePhone   string `description:"收货人手机号"`
	ConsigneeAddress string `description:"收货人详细地址"`
	//商品订单维度
	OrderAddGoodsInfos []*OrderAddGoodsInfo
}

type OrderAddGoodsInfo struct {
	Id             int
	OrderId        int
	GoodsId        int
	GoodsOptionsId int
	Count          int
	Remark         string
	Price          int
	CouponPrice    int
	ActualPrice    int
}

type OrderAddOutput struct {
	Id uint `json:"id"`
}
