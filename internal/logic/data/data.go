package data

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
	"shop/utility"
	"time"
)

type sData struct{}

func New() *sData { return &sData{} }
func init() {
	service.RegisterData(New())
}
func (s *sData) HeadCard(ctx context.Context) (out *model.HeadDataOutput, err error) {
	return &model.HeadDataOutput{
		TodayOrderCount: TodayOderCount(ctx),
		DAU:             utility.RandInt(200),
		ConversionRate:  utility.RandInt(80),
	}, nil
}

func TodayOderCount(ctx context.Context) (count int) {
	count, err := dao.OrderInfo.Ctx(ctx).
		WhereBetween(dao.OrderInfo.Columns().CreatedAt, gtime.New(time.Now()).StartOfDay(), gtime.New(time.Now()).EndOfDay()).
		Count("id")
	if err != nil {
		return 0
	}
	return
}
func (s *sData) Echarts(ctx context.Context) (out *model.EchartsOutput, err error) {
	return &model.EchartsOutput{
		OrderTotal:           OderTotal(ctx),
		SalePriceTotal:       SalePriceTotal(ctx),
		ConsumptionPerPerson: OderTotal(ctx),
		NewOrder:             OderTotal(ctx),
	}, nil

}

//select date_format(created_at, '%Y-%m-%d') today, count(*) as cnt from order_info group by today
/**
gf官方示例：
// SELECT COUNT(*) total,age FROM `user` GROUP BY age
db.Model("user").Fields("COUNT(*) total,age").Group("age").All()
*/
func OderTotal(ctx context.Context) (counts []int) {
	// 初始化一个包含7个0的切片，用于存储最近7天的订单总数
	counts = []int{0, 0, 0, 0, 0, 0, 0}
	// 获取最近7天的日期
	recent7Dates := utility.GetRecent7Date()
	// 初始化一个切片，用于存储从数据库查询到的每天的订单总数
	TodayTotals := []model.TodayTotal{}
	// 查询数据库，获取最近7天的订单总数，按照日期分组
	err := dao.OrderInfo.Ctx(ctx).
		Where(dao.OrderInfo.Columns().CreatedAt+">=", utility.GetBefore7Date()).
		Fields("count(*) total,data_format(created_at,‘%Y-%m-%d’) today").
		Group("today").
		Scan(&TodayTotals)
	// 打印查询结果
	fmt.Printf("result:%v", TodayTotals)
	// 遍历最近7天的日期
	for i, date := range recent7Dates {
		// 遍历从数据库查询到的每天的订单总数
		for _, todayTotal := range TodayTotals {
			// 如果日期匹配，则将订单总数添加到对应的位置
			if todayTotal.Today == date {
				counts[i] = todayTotal.Total
			}
		}
	}
	// 如果查询过程中发生错误，则返回当前的订单总数
	if err != nil {
		return counts
	}
	// 返回最终的订单总数
	return
}

func SalePriceTotal(ctx context.Context) (totals []int) {
	// 初始化一个包含7个0的切片，用于存储最近7天的销售总额
	totals = []int{0, 0, 0, 0, 0, 0, 0}
	// 获取最近7天的日期
	recent7Dates := utility.GetRecent7Date()
	// 初始化一个切片，用于存储从数据库查询到的每天的销售总额
	TodayTotals := []model.TodayTotal{}
	// 查询数据库，获取最近7天的销售总额，按照日期分组
	err := dao.OrderInfo.Ctx(ctx).Where(dao.OrderInfo.Columns().CreatedAt+" >= ", utility.GetBefore7Date()).Fields("sum(actual_price) total,date_format(created_at, '%Y-%m-%d') today").Group("today").Scan(&TodayTotals)
	// 打印查询结果
	fmt.Printf("result:%v", TodayTotals)
	// 遍历最近7天的日期
	for i, date := range recent7Dates {
		// 遍历从数据库查询到的每天的销售总额
		for _, todayTotal := range TodayTotals {
			// 如果日期匹配，则将销售总额添加到对应的位置
			if date == todayTotal.Today {
				totals[i] = todayTotal.Total
			}
		}
	}
	// 如果查询过程中发生错误，则返回当前的销售总额
	if err != nil {
		return totals
	}
	// 返回最终的销售总额
	return
}
