// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"shop/internal/model"
)

type (
	IFile interface {
		/*
			*
			1. 定义图片上传位置
			2. 校验：上传位置是否正确、安全性校验：1分钟内只能上传10次
			3. 定义年月日 Ymd
			4. 入库
			5. 返回数据
		*/
		Upload(ctx context.Context, in model.FileUploadInput) (out *model.FileUploadOutput, err error)
	}
)

var (
	localFile IFile
)

func File() IFile {
	if localFile == nil {
		panic("implement not found for interface IFile, forgot register?")
	}
	return localFile
}

func RegisterFile(i IFile) {
	localFile = i
}
