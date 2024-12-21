package model

import "shop/internal/model/entity"

type GoodsOptionsCreateUpdateBase struct {
	GoodsId    uint
	PicUrl     string
	Name       string
	Price      uint
	Stock      int
	Sale       uint
	Tags       string
	DetailInfo string
}

type GoodsOptionsCreateInput struct {
	GoodsOptionsCreateUpdateBase
}
type GoodsOptionsCreateOutput struct {
	Id uint `json:"id"`
}
type GoodsOptionsUpdateInput struct {
	GoodsOptionsCreateUpdateBase
	Id uint
}

type GoodsOptionsGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

type GoodsOptionsGetListOutput struct {
	List  []GoodsOptionsGetListOutputItem `json:"list" dc:"列表"`
	Page  int                             `json:"page" dc:"分页码"`
	Size  int                             `json:"size" dc:"分页数量"`
	Total int                             `json:"total" dc:"数据总数"`
}
type GoodsOptionsGetListOutputItem struct {
	entity.GoodsOptionsInfo
}
