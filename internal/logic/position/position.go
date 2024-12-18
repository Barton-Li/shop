package position

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

type sPosition struct{}

func init() {
	service.RegisterPosition(New())
}
func New() *sPosition {

	return &sPosition{}

}

func (s *sPosition) Create(ctx context.Context, in model.PositionCreateInput) (out model.PositionCreateOutput, err error) {
	// 不允许HTML代码
	//if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
	//	return out, err
	//}
	lastInsertID, err := dao.PositionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PositionCreateOutput{PositionId: int(lastInsertID)}, err
}
