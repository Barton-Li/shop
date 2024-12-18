package controller

import (
	"context"
	"shop/api/backend"
	"shop/internal/model"
	"shop/internal/service"
)

type cPosition struct{}

var Position = cPosition{}

func (a *cPosition) Create(ctx context.Context, req *backend.PositionReq) (res *backend.PositionRes, err error) {
	out, err := service.Position().Create(ctx, model.PositionCreateInput{
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			Sort:      req.Sort,
			GoodsId:   req.GoodsId,
			GoodsName: req.GoodsName,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.PositionRes{PositionId: out.PositionId}, nil
}
