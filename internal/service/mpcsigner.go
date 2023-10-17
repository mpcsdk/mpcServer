// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "li17server/api/sign/v1"
)

type (
	IMpcSigner interface {
		// GenContextP2
		GenContextP2(ctx context.Context, sid string, private_key2, public_key string, submit bool) error
		// 4.5.calculate p2_zk_proof by p1_hash_proof, need recal context_p2 by p1_hash_proof
		CalZKProofP2(ctx context.Context, sid string, p1_hash_proof string) error
		// 6.7.calculate v2_public_key by p1_zk_proof, recal context_p2 by p1_zk_proof
		CalPublicKey2(ctx context.Context, sid string, p1_zk_proof string) error
		// 8.calculate request, recal context_p2
		CalRequest(ctx context.Context, sid string, request string) error
		CalMsgSign(ctx context.Context, req *v1.SignMsgReq) error
		CalDomainSign(ctx context.Context, req *v1.SignMsgReq) error
		// 9.signature/
		// func (s *sMpcSigner) CheckCalSign(ctx context.Context, req *v1.SignMsgReq) error {
		// }
		CalSign(ctx context.Context, req *v1.SignMsgReq) error
		// /
		// /
		GetState(ctx context.Context, userId string) string
		// /
		FetchPubKey(ctx context.Context, sid string) (string, error)
		FetchZKProofp2(ctx context.Context, sid string) (string, error)
		FetchSignature(ctx context.Context, sid string) (string, error)
		CleanSignature(ctx context.Context, sid string) (string, error)
		// ///
		FetchTxs(ctx context.Context, sid string) (string, error)
		RecordTxs(ctx context.Context, sid string, val string) (string, error)
		// /
		// // key
		GenNewSid(ctx context.Context, userId string, token string) (string, error)
		Sid2UserId(ctx context.Context, sid string) (string, error)
		Sid2Token(ctx context.Context, sid string) (string, error)
		StateString(state int) string
		// 9.signature
		CalSignTask(ctx context.Context, sid string, msg string, request string) error
	}
)

var (
	localMpcSigner IMpcSigner
)

func MpcSigner() IMpcSigner {
	if localMpcSigner == nil {
		panic("implement not found for interface IMpcSigner, forgot register?")
	}
	return localMpcSigner
}

func RegisterMpcSigner(i IMpcSigner) {
	localMpcSigner = i
}
