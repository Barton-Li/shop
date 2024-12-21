package goods_options

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

type sGoodsOptions struct{}

func New() *sGoodsOptions {
	return &sGoodsOptions{}
}
func init() {

	service.RegisterGoodsOptions(New())
}
func (s *sGoodsOptions) Create(ctx context.Context, in model.GoodsOptionsCreateInput) (out model.GoodsOptionsCreateOutput, err error) {
	lastInsertID, err := dao.GoodsOptionsInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.GoodsOptionsCreateOutput{Id: uint(lastInsertID)}, err
}
func (s *sGoodsOptions) Update(ctx context.Context, in model.GoodsOptionsUpdateInput) (err error) {
	_, err = dao.GoodsOptionsInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.GoodsOptionsInfo.Columns().Id).
		Where(dao.GoodsOptionsInfo.Columns().Id, in.Id).
		Update()
	return err
}
func (s *sGoodsOptions) Delete(ctx context.Context, id uint) (err error) {
	_, err = dao.GoodsOptionsInfo.
		Ctx(ctx).
		Where(dao.GoodsOptionsInfo.Columns().Id, id).
		Delete()
	if err != nil {
		return err
	}
	return
}

// GetList 查询分类列表
func (s *sGoodsOptions) GetList(ctx context.Context, in model.GoodsOptionsGetListInput) (out *model.GoodsOptionsGetListOutput, err error) {
	//1.获得*gdb.Model对象，方面后续调用
	m := dao.GoodsOptionsInfo.Ctx(ctx)
	//2. 实例化响应结构体
	out = &model.GoodsOptionsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	//3. 分页查询
	listModel := m.Page(in.Page, in.Size)
	//4. 再查询count，判断有无数据
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	//5. 延迟初始化list切片 确定有数据，再按期望大小初始化切片容量
	out.List = make([]model.GoodsOptionsGetListOutputItem, 0, in.Size)
	//6. 把查询到的结果赋值到响应结构体中
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
