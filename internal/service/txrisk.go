// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "li17server/api/risk/v1"
	"li17server/internal/model"
)

type (
	ITxRisk interface {
		VerifyMail(ctx context.Context, sid, serial string) error
		VerifyPhone(ctx context.Context, sid, serial string) error
		VerifyCode(ctx context.Context, sid, serial, code string) error
		CheckTxs(ctx context.Context, sid string, from string, txs []*model.SignTxData) (*v1.TxRiskRes, error)
	}
)

var (
	localTxRisk ITxRisk
)

func TxRisk() ITxRisk {
	if localTxRisk == nil {
		panic("implement not found for interface ITxRisk, forgot register?")
	}
	return localTxRisk
}

func RegisterTxRisk(i ITxRisk) {
	localTxRisk = i
}
