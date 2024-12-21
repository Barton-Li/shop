package order

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/internal/consts"
	"shop/internal/dao"
	"shop/internal/service"
	"shop/utility"

	"shop/internal/model"
)

type sOrder struct{}

func New() *sOrder {
	return &sOrder{}
}
func init() {
	service.RegisterOrder(New())
}

func (s *sOrder) Add(ctx context.Context, in model.OrderAddInput) (out *model.OrderAddOutput, err error) {
	// 从上下文中获取用户ID，并转换为uint类型
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	// 生成订单号
	in.Number = utility.GetOrderNum()
	// 初始化输出结构体
	out = &model.OrderAddOutput{}
	// 开启数据库事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 插入订单信息并获取插入的ID
		lastInserId, err := dao.OrderInfo.Ctx(ctx).InsertAndGetId(in)
		if err != nil {
			// 如果插入失败，返回错误
			return err
		}
		// 遍历订单商品信息
		for _, info := range in.OrderAddGoodsInfos {
			// 设置订单ID
			info.OrderId = gconv.Int(lastInserId)
			// 插入订单商品信息并获取插入的ID
			_, err := dao.OrderGoodsInfo.Ctx(ctx).InsertAndGetId(info)
			if err != nil {
				// 如果插入失败，返回错误
				return err
			}
		}
		// 遍历订单商品信息
		for _, info := range in.OrderAddGoodsInfos {
			// 增加商品销量
			_, err := dao.GoodsInfo.Ctx(ctx).WherePri(info.GoodsId).Increment(dao.GoodsInfo.Columns().Sale, info.Count)
			if err != nil {
				// 如果更新失败，返回错误
				return err
			}
			// 减少商品库存
			_, err2 := dao.GoodsInfo.Ctx(ctx).WherePri(info.GoodsId).Decrement(dao.GoodsInfo.Columns().Stock, info.Count)
			if err2 != nil {
				// 如果更新失败，返回错误
				return err
			}
			// 减少商品规格库存
			_, err3 := dao.GoodsOptionsInfo.Ctx(ctx).WherePri(info.GoodsOptionsId).Decrement(dao.GoodsOptionsInfo.Columns().Stock, info.Count)
			if err3 != nil {
				// 如果更新失败，返回错误
				return err
			}
		}
		// 设置输出结构体的ID
		out.Id = uint(lastInserId)
		// 返回nil表示事务成功
		return nil

	})
	// 如果事务执行失败，返回错误
	if err != nil {
		return out, err
	}
	// 返回输出结构体和nil表示成功
	return
}

// WherePri方法的功能同Where，但提供了对表主键的智能识别
func (s *sOrder) Detail(ctx context.Context, in model.OrderDetailInput) (out *model.OrderDetailOutput, err error) {
	err = dao.OrderInfo.Ctx(ctx).WithAll().WherePri(in.Id).Scan(&out)
	return
}
func (s *sOrder) List(ctx context.Context, in model.OrderListInput) (out *model.OrderListOutput, err error) {
	//1.获得*gdb.Model对象，方面后续调用
	whereCondition := s.orderListCondition(in)
	m := dao.OrderInfo.Ctx(ctx).Where(whereCondition)
	//2. 实例化响应结构体
	out = &model.OrderListOutput{
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
	out.List = make([]model.OrderListOutputItem, 0, in.Size)
	//6. 把查询到的结果赋值到响应结构体中
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// orderListCondition 方法根据输入参数构建订单列表的查询条件。
func (s *sOrder) orderListCondition(in model.OrderListInput) *gmap.Map {
	// 创建一个新的 gmap.Map 对象，用于存储查询条件。
	m := gmap.New()

	// 如果输入参数中的订单号不为空，则将订单号作为模糊查询条件添加到映射中。
	if in.Number != "" {
		m.Set(dao.OrderInfo.Columns().Number+" like ", "%"+in.Number+"%")
	}

	// 如果输入参数中的用户 ID 不为 0，则将用户 ID 作为查询条件添加到映射中。
	if in.UserId != 0 {
		m.Set(dao.OrderInfo.Columns().UserId, in.UserId)
	}

	// 如果输入参数中的支付类型不为 0，则将支付类型作为查询条件添加到映射中。
	if in.PayType != 0 {
		m.Set(dao.OrderInfo.Columns().PayType, in.PayType)
	}

	// 如果输入参数中的支付起始日期不为空，则将支付起始日期作为查询条件添加到映射中。
	if in.PayAtGte != "" {
		m.Set(dao.OrderInfo.Columns().PayAt+" >= ", gtime.New(in.PayAtGte).StartOfDay())
	}

	// 如果输入参数中的支付结束日期不为空，则将支付结束日期作为查询条件添加到映射中。
	if in.PayAtLte != "" {
		m.Set(dao.OrderInfo.Columns().PayAt+" <= ", gtime.New(in.PayAtLte).EndOfDay())
	}

	// 如果输入参数中的订单状态不为 0，则将订单状态作为查询条件添加到映射中。
	if in.Status != 0 {
		m.Set(dao.OrderInfo.Columns().Status, in.Status)
	}

	// 如果输入参数中的收货人电话不为空，则将收货人电话作为模糊查询条件添加到映射中。
	if in.ConsigneePhone != "" {
		m.Set(dao.OrderInfo.Columns().ConsigneePhone+" like ", "%"+in.ConsigneePhone+"%")
	}

	// 如果输入参数中的价格下限不为 0，则将价格下限作为查询条件添加到映射中。
	if in.PriceGte != 0 {
		m.Set(dao.OrderInfo.Columns().Price+" >= ", in.PriceGte)
	}

	// 如果输入参数中的价格上限不为 0，则将价格上限作为查询条件添加到映射中。
	if in.PriceLte != 0 {
		m.Set(dao.OrderInfo.Columns().Price+" <= ", in.PriceLte)
	}

	// 如果输入参数中的日期下限不为空，则将日期下限作为查询条件添加到映射中。
	if in.DateGte != "" {
		m.Set(dao.OrderInfo.Columns().CreatedAt+" >= ", gtime.New(in.DateGte).StartOfDay())
	}

	// 如果输入参数中的日期上限不为空，则将日期上限作为查询条件添加到映射中。
	if in.DateLte != "" {
		m.Set(dao.OrderInfo.Columns().CreatedAt+" <= ", gtime.New(in.DateLte).EndOfDay())
	}

	// 返回构建好的查询条件映射。
	return m
}
