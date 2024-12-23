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
	ICategory interface {
		Create(ctx context.Context, in model.CategoryCreateInput) (out model.CategoryCreateOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, id uint) (err error)
		// Update 修改
		Update(ctx context.Context, in model.CategoryUpdateInput) error
		// GetList 查询分类列表
		GetList(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error)
		// GetList 查询分类全部信息-不翻页
		GetListAll(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error)
	}
)

var (
	localCategory ICategory
)

func Category() ICategory {
	if localCategory == nil {
		panic("implement not found for interface ICategory, forgot register?")
	}
	return localCategory
}

func RegisterCategory(i ICategory) {
	localCategory = i
}
