package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 订单接口
type OrderListReq struct {
	g.Meta `path:"order/list" tage:"订单列表" method:"get" summary:"订单列表"`
	CommonPaginationReq
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
type OrderListRes struct {
	CommonPaginationRes
}

type OrderDetailReq struct {
	g.Meta `path:"order/detail" tage:"订单详情" method:"get" summary:"订单详情"`
	Id     uint `json:"id" dc:"订单id" `
}
type OrderDetailRes struct {
	OrderInfoBase
	GoodsInfo []OrderGoodsInfoBase `json:"goods_info" dc:"商品信息"`
}
type OrderInfoBase struct {
	Id               int         `json:"id"               dc:""`
	Number           string      `json:"number"           dc:"订单编号"`
	UserId           int         `json:"user_id"          dc:"用户id"`
	PayType          int         `json:"pay_type"         dc:"支付类型 1 微信 2 支付宝 3 云闪付"`
	Remark           string      `json:"remark"           dc:"备注"`
	PayAt            *gtime.Time `json:"pay_at"           dc:"支付时间"`
	Status           int         `json:"status"           dc:"订单状态1 待支付 2 已支付 3 已发货 4 已收货"`
	ConsigneeName    string      `json:"consignee_name"   dc:"收货人姓名"`
	ConsigneePhone   string      `json:"consignee_phone"  dc:"收货人手机号"`
	ConsigneeAddress string      `json:"consignee_address" dc:"收货地址"`
	Price            int         `json:"price"            dc:"订单金额"`
	CouponPrice      int         `json:"coupon_price"     dc:"优惠券金额"`
	ActualPrice      int         `json:"actual_price"     dc:"实际支付金额"`
	CreatedAt        *gtime.Time `json:"created_at"       dc:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updated_at"       dc:"更新时间"`
}

type OrderGoodsInfoBase struct {
	Id          int         `json:"id"        dc:"商品维度的订单表"`
	OrderId     int         `json:"order_id"  dc:"订单id"`
	GoodsId     int         `json:"goods_id"  dc:"商品id"`
	Count       int         `json:"count"     dc:"商品数量"`
	Price       int         `json:"price"     dc:"商品价格"`
	Remark      string      `json:"remark"    dc:"备注"`
	CouponPrice int         `json:"coupon_price" dc:"优惠券金额"`
	ActualPrice int         `json:"actual_price" dc:"实际支付金额"`
	CreatedAt   *gtime.Time `json:"created_at"       dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updated_at"       dc:"更新时间"`
}
