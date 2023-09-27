// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "li17server/api/risk/v1"
)

type (
	IRPC interface {
		PerformMailCode(ctx context.Context, token, serial string) error
		PerformSmsCode(ctx context.Context, token, serial string) error
		PerformVerifyCode(ctx context.Context, token, serial, code string) error
		PerformRiskTxs(ctx context.Context, userId string, signTxData string) (*v1.TxRiskRes, error)
	}
)

var (
	localRPC IRPC
)

func RPC() IRPC {
	if localRPC == nil {
		panic("implement not found for interface IRPC, forgot register?")
	}
	return localRPC
}

func RegisterRPC(i IRPC) {
	localRPC = i
}
