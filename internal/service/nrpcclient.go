// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "mpcServer/api/risk/nrpc/v1"
)

type (
	INrpcClient interface {
		RpcSendMailCode(ctx context.Context, token, serial string) error
		RpcSendSmsCode(ctx context.Context, token, serial string) error
		RpcVerifyCode(ctx context.Context, token, serial, phoneCode, mailCode string) error
		RpcRiskTxs(ctx context.Context, userId string, signTxData string) (*v1.TxRiskRes, error)
		RpcAlive(ctx context.Context) error
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
