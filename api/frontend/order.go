package frontend

import "github.com/gogf/gf/v2/frame/g"

type AddOrderReq struct {
	g.Meta             `path:"/add/order" method:"post" tags:"前台订单" summary:"创建订单"`
	Price              int                  `json:"price" dc:"订单金额 单位分"`
	CouponPrice        int                  `json:"coupon_price" dc:"优惠券金额 单位分"`
	ActualPrice        int                  `json:"actual_price" dc:"实际支付金额 单位分"`
	ConsigneeName      string               `json:"consignee_name" dc:"收货人姓名"`
	ConsigneePhone     string               `json:"consignee_phone" dc:"收货人电话"`
	ConsigneeAddress   string               `json:"consignee_address" dc:"收货人地址"`
	Remark             string               `json:"remark" dc:"订单备注"`
	OrderAddGoodsInfos []*OrderAddGoodsInfo `json:"order_add_goods_infos"`
}
type OrderAddGoodsInfo struct {
	GoodsId        int    `json:"goods_id" dc:"商品id"`
	GoodsOptionsId int    `json:"goods_options_id" dc:"商品规格id"`
	Count          int    `json:"count" dc:"商品数量"`
	Remark         string `json:"remark"`
	Price          int    `json:"price"`
	CouponPrice    int    `json:"coupon_price"`
	ActualPrice    int    `json:"actual_price"`
}

type AddOrderRes struct {
	Id uint `json:"id"`
}
type OrderGoodsInfo struct {
	Id             int `json:"id,omitempty"`
	OrderId        int `json:"order_id"`
	GoodsId        int `json:"goods_id"`
	GoodsOptionsId int `json:"goods_options_id"`
	//商品详情
	GoodsInfo *BaseGoodsColumns
	//注意：api层不需要做orm关联 关联了也没有意义
	//GoodsInfo   *BaseGoodsColumns `orm:"with:id=goods_id" json:"goods_info"`
	Count       int    `json:"count"`
	Remark      string `json:"remark"`
	Status      int    `json:"status"`
	Price       int    `json:"price"`
	CouponPrice int    `json:"coupon_price"`
	ActualPrice int    `json:"actual_price"`
}
