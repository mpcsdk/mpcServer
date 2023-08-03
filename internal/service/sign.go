// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	IGenerator interface {
		KeygenSendHashProofP1(context1 string) string
		KeygenRecvHashProofP2(context2, proof1 string) string
		KeygenSendZkProofP1(context1 string) string
		KeygenRecvZkProofP1(context1, proof2 string) string
		KeygenSendZkProofP2(context1 string) string
		KeygenRecvZkProofP2(context2, proof1 string) string
		PublicKeyP1(context1 string) string
		PublicKeyP2(context2 string) string
		GenContextP1(preivateKey, publicKey string) string
		GenContextP2(preivateKey, publicKey string) string
		SendZkProofP1(p1 string) string
		RecvZkProofP1(p1, zkproof2 string) string
		SendZkProofP2(p2 string) string
		RecvZkProofP2(p2, zkproof1 string) string
		SignSendRequestP1(context1 string) string
		SignRecvRequestP2(context2 string, request string) string
		SignSendPartialP2(context2, msg string) string
		SignSendPartialP1(context1, sign2, msg string) string
	}
)

var (
	localGenerator IGenerator
)

func Generator() IGenerator {
	if localGenerator == nil {
		panic("implement not found for interface IGenerator, forgot register?")
	}
	return localGenerator
}

func RegisterGenerator(i IGenerator) {
	localGenerator = i
}
