// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"mpcServer/api/riskctrl"
)

type (
	INrpcClient interface {
		Flush()
		RpcRiskTxs(ctx context.Context, userId string, signTxData string) (*riskctrl.TxRequestRes, error)
		RpcAlive(ctx context.Context) error
		RpcSendMailCode(ctx context.Context, token, serial string) error
		RpcSendSmsCode(ctx context.Context, token, serial string) error
		RpcVerifyCode(ctx context.Context, token, serial, phoneCode, mailCode string) error
	}
)

var (
	localNrpcClient INrpcClient
)

func NrpcClient() INrpcClient {
	if localNrpcClient == nil {
		panic("implement not found for interface INrpcClient, forgot register?")
	}
	return localNrpcClient
}

func RegisterNrpcClient(i INrpcClient) {
	localNrpcClient = i
}
