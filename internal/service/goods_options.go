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
	IGoodsOptions interface {
		Create(ctx context.Context, in model.GoodsOptionsCreateInput) (out model.GoodsOptionsCreateOutput, err error)
		Update(ctx context.Context, in model.GoodsOptionsUpdateInput) (err error)
		Delete(ctx context.Context, id uint) (err error)
		// GetList 查询分类列表
		GetList(ctx context.Context, in model.GoodsOptionsGetListInput) (out *model.GoodsOptionsGetListOutput, err error)
	}
)

var (
	localGoodsOptions IGoodsOptions
)

func GoodsOptions() IGoodsOptions {
	if localGoodsOptions == nil {
		panic("implement not found for interface IGoodsOptions, forgot register?")
	}
	return localGoodsOptions
}

func RegisterGoodsOptions(i IGoodsOptions) {
	localGoodsOptions = i
}
