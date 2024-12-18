package retation

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

type sRotation struct{}

func init() {
	service.RegisterRotation(New())
}
func New() *sRotation {
	return &sRotation{}
}

func (s *sRotation) Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error) {
	// 不允许HTML代码
	//if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
	//	return out, err
	//}
	lastInsertID, err := dao.RotationInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RotationCreateOutput{RotationId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sRotation) Delete(ctx context.Context, id uint) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容
		_, err := dao.RotationInfo.Ctx(ctx).Where(g.Map{
			dao.RotationInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// update 修改
func (s *sRotation) Update(ctx context.Context, in model.RotationUpdateInput) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 检查并过滤输入数据中的 HTML 特殊字符
		//if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
		//	return err
		//}
		// 更新轮播图信息，排除 ID 字段
		_, err := dao.RotationInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.RotationInfo.Columns().Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sRotation) GetList(ctx context.Context, in model.RotationGetListInput) (out *model.RotationGetListOutput, err error) {
	// 获取数据库操作对象
	m := dao.RotationInfo.Ctx(ctx)
	// 初始化输出对象
	out = &model.RotationGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分页查询轮播图信息
	listModel := m.Page(in.Page, in.Size)
	// 获取轮播图总数
	out.Total, err = m.Count()
	// 若发生错误或总数为0，则返回
	if err != nil || out.Total == 0 {
		return out, err
	}
	// 初始化列表切片
	out.List = make([]model.RotationGetListOutputItem, 0, in.Size)
	// 将查询结果扫描到列表中
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	// 返回结果
	return
}
