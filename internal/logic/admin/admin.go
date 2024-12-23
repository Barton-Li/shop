package admin

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"
	"shop/utility"
)

type sAdmin struct{}

func init() {
	service.RegisterAdmin(New())
}

func New() *sAdmin {
	return &sAdmin{}
}

func (s *sAdmin) Create(ctx context.Context, in model.AdminCreateInput) (out model.AdminCreateOutput, err error) {
	// 检查输入数据中是否包含 HTML 特殊字符，防止 XSS 攻击
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		// 如果输入数据中包含特殊字符，返回错误
		return out, err
	}
	// 生成一个 10 位的随机字符串作为用户盐值
	UserSalt := grand.S(10)
	// 使用用户盐值对输入的密码进行加密
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	// 将用户盐值保存到输入结构体中
	in.UserSalt = UserSalt
	// 将新管理员账户信息插入数据库，并获取插入的 ID
	lastInserID, err := dao.AddressInfo.Ctx(ctx).Data(in).InsertAndGetId()
	// 如果插入过程中发生错误，返回错误
	if err != nil {
		return out, err
	}
	// 返回创建成功的管理员账户 ID
	return model.AdminCreateOutput{AdminId: int(lastInserID)}, err
}

// Delete 删除
func (s *sAdmin) Delete(ctx context.Context, id uint) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容
		_, err := dao.AdminInfo.Ctx(ctx).Where(g.Map{
			dao.AdminInfo.Columns().Id: id,
		}).Delete()
		return err
	})
}

// Update 修改
func (s *sAdmin) Update(ctx context.Context, in model.AdminUpdateInput) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		//判断是否修改了密码
		if in.Password != "" {
			UserSalt := grand.S(10)
			in.Password = utility.EncryptPassword(in.Password, UserSalt)
			in.UserSalt = UserSalt
		}
		//更新操作
		_, err := dao.AdminInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.AdminInfo.Columns().Id).
			Where(dao.AdminInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sAdmin) GetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error) {
	var (
		m = dao.AdminInfo.Ctx(ctx)
	)
	out = &model.AdminGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.AdminInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Admin
	//指定item的键名用：ScanList
	//if err := listModel.ScanList(&out.List, "Admin"); err != nil {
	//不指定item的键名用：Scan
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
func (s *sAdmin) GetUserByUserNamePassword(ctx context.Context, in model.UserLoginInput) map[string]interface{} {
	//验证账号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, in.Name).Scan(&adminInfo)
	if err != nil {
		return nil
	}
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return nil
	} else {
		return g.Map{
			"id":       adminInfo.Id,
			"username": adminInfo.Name,
		}
	}
}
func (s *sAdmin) GetAdminByNamePassword(ctx context.Context, in model.UserLoginInput) map[string]interface{} {
	//验证账号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return nil
	}
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return nil
	} else {
		return g.Map{
			"id":       adminInfo.Id,
			"username": adminInfo.Name,
		}
	}
}
