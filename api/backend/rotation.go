package backend

import "github.com/gogf/gf/v2/frame/g"

type RotationReq struct {
	// g.Meta 包含了请求的元数据，如路径、方法、标签和摘要
	g.Meta `path:"/rotation/add" tags:"Rotation" method:"post" summary:"You first rotation api"`
	// PicUrl 是图片的 URL 地址，是一个字符串，必须提供
	PicUrl string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	// Link 是点击图片后跳转的链接，也是一个字符串，必须提供
	Link string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	// Sort 是轮播图的排序，是一个整数，可选
	Sort int `json:"sort"    dc:"排序"`
}

// RotationRes 是添加轮播图操作的响应
type RotationRes struct {
	// RotationId 是新添加的轮播图的 ID
	RotationId int `json:"rotation_id"`
}

// RotationDeleteReq 是删除轮播图操作的请求
type RotationDeleteReq struct {
	// g.Meta 包含了请求的元数据，如路径、方法、标签和摘要
	g.Meta `path:"/rotation/delete" method:"delete" tags:"轮播图" summary:"删除轮播图接口"`
	// Id 是要删除的轮播图的 ID，必须提供，且值必须大于等于 1
	Id uint `v:"min:1#请选择需要删除的轮播图" dc:"轮播图id"`
}

// RotationDeleteRes 是删除轮播图操作的响应
type RotationDeleteRes struct{}

type RotationUpdateReq struct {
	g.Meta `path:"/rotation/update/" method:"post" tags:"轮播图" summary:"修改轮播图接口"`
	Id     uint   `json:"id"      v:"min:1#请选择需要修改的轮播图" dc:"轮播图Id"`
	PicUrl string `json:"pic_url" v:"required#轮播图图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort"    dc:"跳转链接"`
}

type RotationUpdateRes struct {
	Id uint `json:"id"`
}
type RotationGetListCommonReq struct {
	g.Meta `path:"/rotation/list" method:"get" tags:"轮播图" summary:"轮播图列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}

type RotationGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
