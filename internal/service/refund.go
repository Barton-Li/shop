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
	IRefund interface {
		Create(ctx context.Context, in model.RefundAddInput) (out model.RefundAddOutput, err error)
		// GetList 查询分类列表
		GetList(ctx context.Context, in model.RefundListInput) (out *model.RefundListOutput, err error)
		// 详情
		Detail(ctx context.Context, in model.RefundDetailInput) (out *model.RefundDetailOutput, err error)
	}
)

var (
	localRefund IRefund
)

func Refund() IRefund {
	if localRefund == nil {
		panic("implement not found for interface IRefund, forgot register?")
	}
	return localRefund
}

func RegisterRefund(i IRefund) {
	localRefund = i
}
