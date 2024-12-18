package controller

import (
	"context"
	"shop/api/backend"
	"shop/internal/model"
	"shop/internal/service"
)

var Rotation = cRotation{}

type cRotation struct{}

func (a *cRotation) Create(ctx context.Context, req *backend.RotationReq) (res *backend.RotationRes, err error) {
	// 调用 service 层的 Create 方法创建轮播图
	out, err := service.Rotation().Create(ctx, model.RotationCreateInput{
		// 初始化创建或更新轮播图的基础数据
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			// 设置图片链接
			PicUrl: req.PicUrl,
			// 设置跳转链接
			Link: req.Link,
			// 设置排序
			Sort: req.Sort,
		},
	})
	// 如果发生错误，返回 nil 和错误信息
	if err != nil {
		return nil, err
	}
	// 创建响应对象
	return &backend.RotationRes{RotationId: out.RotationId}, nil
}

func (a *cRotation) Delete(ctx context.Context, req *backend.RotationDeleteReq) (res *backend.RotationDeleteRes, err error) {
	// 调用 service 层的 Delete 方法删除轮播图
	err = service.Rotation().Delete(ctx, req.Id)
	// 如果发生错误，返回 nil 和错误信息
	if err != nil {
		return nil, err
	}
	// 创建响应对象
	return &backend.RotationDeleteRes{}, nil
}

func (a *cRotation) Update(ctx context.Context, req *backend.RotationUpdateReq) (res *backend.RotationUpdateRes, err error) {
	// 调用 service 层的 Update 方法更新轮播图
	err = service.Rotation().Update(ctx, model.RotationUpdateInput{
		// 初始化创建或更新轮播图的基础数据
		Id: req.Id,
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	// 如果发生错误，返回 nil 和错误信息
	if err != nil {
		return nil, err
	}
	// 创建响应对象
	return &backend.RotationUpdateRes{Id: req.Id}, nil
}
func (a *cRotation) List(ctx context.Context, req *backend.RotationGetListCommonReq) (res *backend.RotationGetListCommonRes, err error) {
	// 调用 service 层的 GetList 方法获取轮播图列表
	getListRes, err := service.Rotation().GetList(ctx, model.RotationGetListInput{
		// 设置分页号码
		Page: req.Page,
		// 设置分页数量
		Size: req.Size,
		// 设置排序类型
		Sort: req.Sort,
	})
	// 如果发生错误，返回 nil 和错误信息
	if err != nil {
		return nil, err
	}

	// 创建响应对象
	return &backend.RotationGetListCommonRes{List: getListRes.List,
		// 设置当前页码
		Page: getListRes.Page,
		// 设置每页的条目数
		Size: getListRes.Size,
		// 设置数据总数
		Total: getListRes.Total}, nil
}
