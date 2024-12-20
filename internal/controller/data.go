package controller

import (
	"context"
	"shop/api/backend"
	"shop/internal/service"
)

type cData struct{}

var Data cData

func (c *cData) HeadCard(ctx context.Context, req *backend.DataHeadReq) (res *backend.DataHeadRes, err error) {
	card, err := service.Data().HeadCard(ctx)
	if err != nil {
		return &backend.DataHeadRes{}, err
	}
	return &backend.DataHeadRes{
		TadayOrderCount: card.TodayOrderCount,
		DAU:             card.DAU,
		ConversionRate:  card.ConversionRate,
	}, err
}
func (c *cData) Echarts(ctx context.Context, req *backend.DataEchartsReq) (res *backend.DataEchartsRes, err error) {
	echarts, err := service.Data().Echarts(ctx)
	if err != nil {
		return &backend.DataEchartsRes{}, err
	}
	return &backend.DataEchartsRes{
		OdderTotal:           echarts.OrderTotal,
		SalePriceTotal:       echarts.SalePriceTotal,
		ConsumptionPerPerson: echarts.ConsumptionPerPerson,
		NewOrder:             echarts.NewOrder,
	}, err
}
