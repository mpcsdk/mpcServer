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
	IGenerator interface {
		// GenContextP2
		GenContextP2(ctx context.Context, sid string, private_key2, public_key string, submit bool) error
		// 4.5.calculate p2_zk_proof by p1_hash_proof, need recal context_p2 by p1_hash_proof
		CalZKProofP2(ctx context.Context, sid string, p1_hash_proof string) error
		// 6.7.calculate v2_public_key by p1_zk_proof, recal context_p2 by p1_zk_proof
		CalPublicKey2(ctx context.Context, sid string, p1_zk_proof string) error
		// 8.calculate request, recal context_p2
		CalRequest(ctx context.Context, sid string, request string) error
		CalMsgSign(ctx context.Context, req *v1.SignMsgReq) error
		// 9.signature/
		// func (s *sGenerator) CheckCalSign(ctx context.Context, req *v1.SignMsgReq) error {
		// }
		CalSign(ctx context.Context, req *v1.SignMsgReq) error
		StateNext(state int) int
		StatePrivate(state int) int
		StateInt(state string) int
		StateString(state int) string
		StateIs(state string, istate int) bool
		NextStateIs(curstate string) int
		UpState(ctx context.Context, userId string, state string, err error) error
		// /
		// /
		GetState(ctx context.Context, userId string) (string, error)
		// /
		RecordSid(ctx context.Context, sid string, key string, val string) error
		FetchSid(ctx context.Context, sid string, key string) (string, error)
		RecordUserId(ctx context.Context, userId string, key string, val string) error
		FetchUserId(ctx context.Context, userId string, key string) (string, error)
		// // key
		GenNewSid(ctx context.Context, userId string) (string, error)
		Sid2UserId(ctx context.Context, sid string) (string, error)
		Sid2Token(ctx context.Context, sid string) (string, error)
		// 9.signature
		CalSignTask(ctx context.Context, sid string, msg string, request string) error
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
