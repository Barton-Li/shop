// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"shop/internal/model"
)

type (
	IOrder interface {
		Add(ctx context.Context, in model.OrderAddInput) (out *model.OrderAddOutput, err error)
		// WherePri方法的功能同Where，但提供了对表主键的智能识别
		Detail(ctx context.Context, in model.OrderDetailInput) (out *model.OrderDetailOutput, err error)
		List(ctx context.Context, in model.OrderListInput) (out *model.OrderListOutput, err error)
	}
)

var (
	localOrder IOrder
)

func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

func RegisterOrder(i IOrder) {
	localOrder = i
}
