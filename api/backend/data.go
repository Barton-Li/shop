package backend

import "github.com/gogf/gf/v2/frame/g"

type DataHeadReq struct {
	g.Meta `path:"/data/head/"method:"get" tags:"数据管理" summary:"数据大屏头部卡片"`
}
type DataHeadRes struct {
	TadayOrderCount int `json:"taday_order_count" dc: "今日订单量"` // 今日订单数
	DAU             int `json:"dau" dc: "日活"`
	ConversionRate  int `json:"conversion_rate" dc: "转化率"`
}
type DataEchartsReq struct {
	g.Meta `path:"/data/echarts/"method:"get" tags:"数据管理" summary:"数据大屏echarts"`
}
type DataEchartsRes struct {
	OdderTotal           []int `json:"odder_total" dc: "订单总量"`
	SalePriceTotal       []int `json:"sale_price_total" desc:"销售价格"`
	ConsumptionPerPerson []int `json:"consumption_per_person" desc:"人均消费"`
	NewOrder             []int `json:"new_order" desc:"新增订单"`
}
