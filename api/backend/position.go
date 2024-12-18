package backend

import "github.com/gogf/gf/v2/frame/g"

// PositionReq 结构体定义了添加轮播图操作的请求参数
type PositionReq struct {
	// g.Meta 包含了请求的元数据，如路径、方法、标签和摘要
	g.Meta `path:"/position/add" tags:"Position" method:"post" summary:"You first position api"`
	// PicUrl 是图片的 URL 地址，是一个字符串
	PicUrl string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	// Link 是点击图片后跳转的链接，也是一个字符串
	Link string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	// GoodsName 是商品的名称，是一个字符串
	GoodsName string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"` //冗余设计
	// GoodsId 是商品的 ID，是一个整数
	GoodsId uint `json:"goods_id"  v:"required#商品Id不能为空" dc:"商品ID"` //mysql三范式
	// Sort 是轮播图的排序，是一个整数
	Sort int `json:"sort"    dc:"排序"`
}

// PositionRes 结构体定义了添加轮播图操作的响应参数
type PositionRes struct {
	// PositionId 是新添加的轮播图的 ID
	PositionId int `json:"position_id"`
}

// PositionDeleteReq 结构体定义了删除轮播图操作的请求参数
type PositionDeleteReq struct {
	// g.Meta 包含了请求的元数据，如路径、方法、标签和摘要
	g.Meta `path:"/position/delete" method:"delete" tags:"手工位图" summary:"删除手工位图接口"`
	// Id 是要删除的轮播图的 ID，是一个整数
	Id uint `v:"min:1#请选择需要删除的手工位图" dc:"手工位图id"`
}

// PositionDeleteRes 是删除轮播图操作的响应
type PositionDeleteRes struct{}

// PositionUpdateReq 结构体定义了更新轮播图操作的请求参数
type PositionUpdateReq struct {
	// g.Meta 包含了请求的元数据，如路径、方法、标签和摘要
	g.Meta `path:"/position/update/{Id}" method:"post" tags:"手工位图" summary:"修改手工位图接口"`
	// Id 是要更新的轮播图的 ID，是一个整数
	Id uint `json:"id"      v:"min:1#请选择需要修改的手工位图" dc:"手工位图Id"`
	// PicUrl 是图片的 URL 地址，是一个字符串
	PicUrl string `json:"pic_url" v:"required#手工位图图片链接不能为空" dc:"图片链接"`
	// Link 是点击图片后跳转的链接，也是一个字符串
	Link string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	// Sort 是轮播图的排序，是一个整数
	Sort int `json:"sort"    dc:"跳转链接"`
	// GoodsName 是商品的名称，是一个字符串
	GoodsName string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"` //冗余设计
	// GoodsId 是商品的 ID，是一个整数
	GoodsId uint `json:"goods_id"  v:"required#商品Id不能为空" dc:"商品ID"` //mysql三范式
}

// PositionUpdateRes 结构体定义了更新轮播图操作的响应参数
type PositionUpdateRes struct {
	// Id 是更新后的轮播图的 ID
	Id uint `json:"id"`
}

// PositionGetListCommonReq 结构体定义了获取轮播图列表的通用请求参数
type PositionGetListCommonReq struct {
	// g.Meta 包含了请求的元数据，如路径、方法、标签和摘要
	g.Meta `path:"/position/list" method:"get" tags:"手工位图" summary:"手工位图列表接口"`
	// Sort 是轮播图的排序类型，可选
	Sort int `json:"sort"   in:"query" dc:"排序类型"`
	// CommonPaginationReq 是通用分页请求参数
	CommonPaginationReq
}

// PositionGetListCommonRes 结构体定义了获取轮播图列表的通用响应参数
type PositionGetListCommonRes struct {
	// List 是轮播图列表数据
	List interface{} `json:"list" description:"列表"`
	// Page 是当前页码
	Page int `json:"page" description:"分页码"`
	// Size 是每页的条目数
	Size int `json:"size" description:"分页数量"`
	// Total 是数据总数
	Total int `json:"total" description:"数据总数"`
}
