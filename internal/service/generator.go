// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

const (
	STATE_None int = iota
	STATE_Auth
	STATE_HandShake
	// STATE_ContextP2
	// STATE_ZKProof2
	// STATE_ZKProofP2
	// STATE_PublicKey2
	// STATE_Request
	STATE_Signature
	STATE_Done
	STATE_Err
	
)

type (
	IGenerator interface {
		// GenContextP2
		GenContextP2(ctx context.Context, sid string, private_key2, public_key string) error
		// 1.2.3 cal zk_proof2 by zk_proof1, need recal private_key2_ and context_p2
		CalZKProof2(ctx context.Context, sid string, zk_proof1 string) (err error)
		// 4.5.calculate p2_zk_proof by p1_hash_proof, need recal context_p2 by p1_hash_proof
		CalZKProofP2(ctx context.Context, sid string, p1_hash_proof string) error
		// 6.7.calculate v2_public_key by p1_zk_proof, recal context_p2 by p1_zk_proof
		CalPublicKey2(ctx context.Context, sid string, p1_zk_proof string) error
		// 8.calculate request, recal context_p2
		CalRequest(ctx context.Context, sid string, request string) error
		// 9.signature
		CalSign(ctx context.Context, sid string, msg string, request string) error
		StateNext(state int) int
		StatePrivate(state int) int
		StateInt(state string) int
		StateString(state int) string
		StateIs(state string, istate int) bool
		NextStateIs(curstate string) int
		UpGeneratorState(ctx context.Context, sid string, state string, err error) error
		GetGeneratorState(ctx context.Context, sid string) (string, error)
		GetStateData(ctx context.Context, sid, state string) (string, error)
		// pubkey
		FetchPubKey(ctx context.Context, sid string) (string, error)
		RecordPubKey(ctx context.Context, sid string, pubkey string) error
		// privatekey
		FetchPrivateKey(ctx context.Context, sid string) (string, error)
		RecordPrivateKey(ctx context.Context, sid string, privatekey string) error
		// //p2
		FetchP2(ctx context.Context, sid string) (string, error)
		RecordP2(ctx context.Context, sid string, p2 string) error
		// //private_key2_
		RecordPrivateKey2(ctx context.Context, sid string, pkey string) error
		FetchPrivateKey2(ctx context.Context, sid string) (string, error)
		// //zk_proof2
		RecordZKProof2(ctx context.Context, sid string, zkproof2 string) error
		FetchZKProof2(ctx context.Context, sid string) (string, error)
		// //context_p2
		RecordContextp2(ctx context.Context, sid string, context_p2 string) error
		FetchContextp2(ctx context.Context, sid string) (string, error)
		// //p1_hash_proof
		RecordHashProofP1(ctx context.Context, sid string, hashproofp1 string) error
		FetchHashProofP1(ctx context.Context, sid string) (string, error)
		// //p2_zk_proof
		RecordZKProofP2(ctx context.Context, sid string, p2_zk_proof string) error
		FetchZKProofP2(ctx context.Context, sid string) (string, error)
		// //p1_zk_proof
		RecordZKProofP1(ctx context.Context, sid string, p1_zk_proof string) error
		FetchZKProofP1(ctx context.Context, sid string) (string, error)
		// v1_public_key
		// v2_public_key
		RecordPublicKey2(ctx context.Context, sid string, v2_public_key string) error
		FetchPublicKey2(ctx context.Context, sid string) (string, error)
		// //request
		RecordRequest(ctx context.Context, sid string, request string) error
		FetchRequest(ctx context.Context, sid string) (string, error)
		// //msg
		RecordMsg(ctx context.Context, sid string, msg string) error
		FetchMsg(ctx context.Context, sid string) (string, error)
		// //msg
		RecordSignature(ctx context.Context, sid string, signature string) error
		FetchSignature(ctx context.Context, sid string) (string, error)
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
