// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"mpcServer/internal/model/do"
	"mpcServer/internal/model/entity"
)

type (
	IDB interface {
		InertContext(ctx context.Context, userId string, data *do.MpcContext) error
		UpdateContext(ctx context.Context, userId string, data *do.MpcContext) error
		FetchContext(ctx context.Context, userId string) (*entity.MpcContext, error)
	}
)

var (
	localDB IDB
)

func DB() IDB {
	if localDB == nil {
		panic("implement not found for interface IDB, forgot register?")
	}
	return localDB
}

func RegisterDB(i IDB) {
	localDB = i
}
