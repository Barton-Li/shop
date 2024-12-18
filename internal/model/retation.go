package model

import "github.com/gogf/gf/v2/os/gtime"

type RotationCreateUpdateBase struct {
	// PicUrl 是图片的 URL 地址，是一个字符串
	PicUrl string
	// Link 是点击图片后跳转的链接，也是一个字符串
	Link string
	// Sort 是轮播图的排序，是一个整数
	Sort int
}

// RotationCreateOutput 是添加轮播图操作的响应
type RotationCreateOutput struct {
	// RotationId 是新添加的轮播图的 ID
	RotationId int `json:"rotation_id"`
}

// RotationCreateInput 是添加轮播图操作的请求
type RotationCreateInput struct {
	// RotationCreateUpdateBase 是添加轮播图操作的基础数据
	RotationCreateUpdateBase
}

// RotationUpdateInput 是更新轮播图操作的请求
type RotationUpdateInput struct {
	// RotationCreateUpdateBase 是更新轮播图操作的基础数据
	RotationCreateUpdateBase
	// Id 是要更新的轮播图的 ID，是一个整数
	Id uint
}

// RotationGetListInput 获取内容列表
type RotationGetListInput struct {
	// Page 是分页号码，默认值为 1
	Page int // 分页号码
	// Size 是分页数量，默认值为 10，最大值为 50
	Size int // 分页数量，最大50
	// Sort 是排序类型，默认值为 0，表示最新
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// RotationGetListOutput 查询列表结果
type RotationGetListOutput struct {
	// List 是轮播图列表数据
	List []RotationGetListOutputItem `json:"list" description:"列表"`
	// Page 是当前页码
	Page int `json:"page" description:"分页码"`
	// Size 是每页的条目数
	Size int `json:"size" description:"分页数量"`
	// Total 是数据总数
	Total int `json:"total" description:"数据总数"`
}

// RotationGetListOutputItem 是轮播图列表项的数据结构
type RotationGetListOutputItem struct {
	// Id 是轮播图的 ID
	Id uint `json:"id"` // 自增ID
	// PicUrl 是轮播图的图片链接
	PicUrl string `json:"pic_url"`
	// Link 是轮播图的跳转链接
	Link string `json:"link"`
	// Sort 是轮播图的排序，数值越低越靠前
	Sort uint `json:"sort"` // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	// CreatedAt 是轮播图的创建时间
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	// UpdatedAt 是轮播图的更新时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
