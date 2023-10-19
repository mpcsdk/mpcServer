// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"mpcServer/internal/model"
)

type (
	ISigner interface {
		KeygenSendHashProofP1(context1 string) model.ISignerPromise
		KeygenRecvHashProofP2(context2, proof1 string) model.ISignerPromise
		KeygenSendZKProofP1(context1 string) model.ISignerPromise
		KeygenRecvZKProofP1(context1, proof2 string) model.ISignerPromise
		KeygenSendZKProofP2(context1 string) model.ISignerPromise
		KeygenRecvZKProofP2(context2, proof1 string) model.ISignerPromise
		PublicKeyP1(context1 string) model.ISignerPromise
		PublicKeyP2(context2 string) model.ISignerPromise
		AddTask(task func())
		TaskLen() int
		Stop()
		GenContextP1(preivateKey, publicKey string) model.ISignerPromise
		GenContextP2(preivateKey, publicKey string) model.ISignerPromise
		SendZKProofP1(p1 string) model.ISignerPromise
		RecvZKProofP1(p1, ZKProof2 string) model.ISignerPromise
		SendZKProofP2(p2 string) model.ISignerPromise
		RecvZKProofP2(p2, ZKProof1 string) model.ISignerPromise
		SignSendRequestP1(context1 string) model.ISignerPromise
		SignRecvRequestP2(context2 string, request string) model.ISignerPromise
		SignSendPartialP2(context2, msg string) model.ISignerPromise
		SignSendPartialP1(context1, sign2, msg string) model.ISignerPromise
	}
)

var (
	localSigner ISigner
)

func Signer() ISigner {
	if localSigner == nil {
		panic("implement not found for interface ISigner, forgot register?")
	}
	return localSigner
}

func RegisterSigner(i ISigner) {
	localSigner = i
}
