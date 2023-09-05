// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"li17server/internal/model"
)

type (
	IEthTx interface {
		AnalzyTxs(ctx context.Context, signtxs *model.SignTx) (*model.AnalzyTx, error)
	}
)

var (
	localEthTx IEthTx
)

func EthTx() IEthTx {
	if localEthTx == nil {
		panic("implement not found for interface IEthTx, forgot register?")
	}
	return localEthTx
}

func RegisterEthTx(i IEthTx) {
	localEthTx = i
}
