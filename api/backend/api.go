package backend

// CommonPaginationReq 结构体定义了通用分页请求的参数
type CommonPaginationReq struct {
	// Page 表示页码，默认值为 1，最小值为 0
	Page int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	// Size 表示每页的条目数，默认值为 10，最大值为 50
	Size int `json:"size" in:"query" d:"10" v:"max:50#分页数量最大50条" dc:"分页数量，最大50"`
}

// CommonPaginationRes 结构体定义了通用分页响应的参数
type CommonPaginationRes struct {
	// List 表示分页数据列表
	List interface{} `dc:"列表数据"`
	// Total 表示数据总数
	Total int `dc:"总数"`
	// Page 表示当前页码
	Page int `dc:"分页号码"`
	// Size 表示每页的条目数
	Size int `dc:"分页数量"`
}
