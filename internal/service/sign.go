// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	ISign interface {
		KeygenSendHashProofP1(context1 string) string
		KeygenRecvHashProofP2(context2, proof1 string) string
		KeygenSendZKProofP1(context1 string) string
		KeygenRecvZKProofP1(context1, proof2 string) string
		KeygenSendZKProofP2(context1 string) string
		KeygenRecvZKProofP2(context2, proof1 string) string
		PublicKeyP1(context1 string) string
		PublicKeyP2(context2 string) string
		GenContextP1(preivateKey, publicKey string) string
		GenContextP2(preivateKey, publicKey string) string
		SendZKProofP1(p1 string) string
		RecvZKProofP1(p1, ZKProof2 string) string
		SendZKProofP2(p2 string) string
		RecvZKProofP2(p2, ZKProof1 string) string
		SignSendRequestP1(context1 string) string
		SignRecvRequestP2(context2 string, request string) string
		SignSendPartialP2(context2, msg string) string
		SignSendPartialP1(context1, sign2, msg string) string
	}
)

var (
	localSign ISign
)

func Sign() ISign {
	if localSign == nil {
		panic("implement not found for interface ISign, forgot register?")
	}
	return localSign
}

func RegisterSign(i ISign) {
	localSign = i
}
