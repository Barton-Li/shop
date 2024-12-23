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
	ICoupon interface {
		Create(ctx context.Context, in model.CouponCreateInput) (out model.CouponCreateOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, id uint) (err error)
		// Update 修改
		Update(ctx context.Context, in model.CouponUpdateInput) error
		// GetList 查询分类列表
		GetList(ctx context.Context, in model.CouponGetListInput) (out *model.CouponGetListOutput, err error)
		// GetList 查询分类全部信息-不翻页
		GetListAll(ctx context.Context, in model.CouponGetListInput) (out *model.CouponGetListOutput, err error)
	}
)

var (
	localCoupon ICoupon
)

func Coupon() ICoupon {
	if localCoupon == nil {
		panic("implement not found for interface ICoupon, forgot register?")
	}
	return localCoupon
}

func RegisterCoupon(i ICoupon) {
	localCoupon = i
}
