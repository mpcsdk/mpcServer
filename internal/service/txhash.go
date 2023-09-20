// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ITxHash interface {
		DigestTxHash(ctx context.Context, msg string) (string, error)
		HasDomain(ctx context.Context, msg string) string
		TypedDataEncoderHash(ctx context.Context, msg string) (string, error)
	}
)

var (
	localTxHash ITxHash
)

func TxHash() ITxHash {
	if localTxHash == nil {
		panic("implement not found for interface ITxHash, forgot register?")
	}
	return localTxHash
}

func RegisterTxHash(i ITxHash) {
	localTxHash = i
}
