package model

type RotationCreateUpdateBase struct {
	PicUrl string
	Link   string
	Sort   int
}

type RotationCreateOutput struct {
	RotationId int `json:"rotation_id"`
}
type RotationCreateInput struct {
	RotationCreateUpdateBase
}
