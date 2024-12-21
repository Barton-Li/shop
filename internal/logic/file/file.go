package file

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/internal/consts"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"
	"time"
)

type sFile struct{}

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

/*
*
1. 定义图片上传位置
2. 校验：上传位置是否正确、安全性校验：1分钟内只能上传10次
3. 定义年月日 Ymd
4. 入库
5. 返回数据
*/
func (s *sFile) Upload(ctx context.Context, in model.FileUploadInput) (out *model.FileUploadOutput, err error) {
	// 从配置文件中获取上传路径
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	// 如果上传路径为空，返回错误
	if uploadPath == "" {
		return nil, gerror.New("读取配置文件失败，上传路径不存在")
	}
	// 如果文件名不为空，则设置文件名为指定的名称
	if in.Name != " " {
		in.File.Filename = in.Name
	}
	// 统计用户在一分钟内上传的文件数量
	count, err := dao.FileInfo.Ctx(ctx).
		Where(dao.FileInfo.Columns().UserId, gconv.Int(ctx.Value(consts.CtxAdminId))).
		WhereGT(dao.FileInfo.Columns().CreatedAt, gtime.Now().Add(-time.Minute)).
		Count()
	// 如果统计出错，返回错误
	if err != nil {
		return nil, err
	}
	// 如果上传次数超过限制，返回错误
	if count >= consts.FileMaxUploadCountMinute {
		return nil, gerror.New("上传频繁，1分钟内只能上传10次")
	}
	// 定义年月日 Ymd
	dateDirName := gtime.Now().Format("Ymd")
	// 保存文件到指定目录，并获取文件名
	fileName, err := in.File.Save(gfile.Join(uploadPath, dateDirName), in.RandomName)
	// 如果保存文件出错，返回错误
	if err != nil {
		return nil, err
	}
	// 4. 入库
	data := entity.FileInfo{
		Name:   fileName,
		Src:    gfile.Join(uploadPath, dateDirName, fileName),
		Url:    "/upload/" + dateDirName + "/" + fileName, //和上面gfile.Join()效果一样
		UserId: gconv.Int(ctx.Value(consts.CtxAdminId)),
	}
	// 将文件信息插入数据库，并获取插入的ID
	id, err := dao.FileInfo.Ctx(ctx).Data(data).OmitEmpty().InsertAndGetId()
	// 如果插入数据出错，返回错误
	if err != nil {
		return nil, err
	}
	// 返回上传文件的信息
	return &model.FileUploadOutput{
		Id:   uint(id),
		Name: data.Name,
		Src:  data.Src,
		Url:  data.Url,
	}, nil
}
