package model

import "github.com/gogf/gf/v2/os/gtime"

type PositionCreateUpdateBase struct {
	PicUrl    string
	Link      string
	Sort      int
	GoodsName string
	GoodsId   uint
}

type PositionCreateInput struct {
	PositionCreateUpdateBase
}

type PositionCreateOutput struct {
	PositionId int `json:"position_id"`
}
type PositionUpdateInput struct {
	PositionCreateUpdateBase
	Id uint
}

// PositionGetListInput 获取内容列表
type PositionGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// PositionGetListOutput 查询列表结果
type PositionGetListOutput struct {
	List  []PositionGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}
type PositionGetListOutputItem struct {
	//指定item的键名的方式
	//Position *PositionListItem `json:"Position"`
	//不指定item键名的方式
	Id        uint        `json:"id"`         // 自增ID
	PicUrl    string      `json:"pic_url"`    //图片链接
	Name      string      `json:"name"`       //手工位名称
	GoodsId   uint        `json:"goods_id"`   //商品编号（id）
	Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	Link      string      `json:"brief"`      // 跳转链接
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type PositionListItem struct {
	Id        uint        `json:"id"`         // 自增ID
	PicUrl    string      `json:"pic_url"`    //图片链接
	Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	GoodsId   uint        `json:"goods_id"`   // 商品id todo
	Link      string      `json:"brief"`      // 跳转链接
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
