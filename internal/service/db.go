// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IDb interface {
		GetAbi(ctx context.Context, addr string) (string, error)
	}
)

var (
	localDb IDb
)

func Db() IDb {
	if localDb == nil {
		panic("implement not found for interface IDb, forgot register?")
	}
	return localDb
}

func RegisterDb(i IDb) {
	localDb = i
}
