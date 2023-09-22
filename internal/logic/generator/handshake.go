package generator

import (
	"context"
	"li17server/internal/consts"
	"li17server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sGenerator) FetchPubKey(ctx context.Context, sid string) (string, error) {
	////
	///
	userId, err := service.Generator().Sid2UserId(ctx, sid)
	if err != nil {
		return "", gerror.NewCode(consts.CodeInternalError)
	}
	////

	///
	pubkey, err := s.fetchUserId(ctx, userId, consts.KEY_publickey2)
	if err != nil {
		return "", gerror.NewCode(consts.CodeInternalError)
	}

	return pubkey, err
}
func (s *sGenerator) FetchZKProofp2(ctx context.Context, sid string) (string, error) {
	////
	ZKProofp2, err := s.fetchSid(ctx, sid, consts.KEY_zkproof2)
	if err != nil {
		return "", gerror.NewCode(consts.CodeInternalError)
	}

	return ZKProofp2, err
}
func (s *sGenerator) FetchSignature(ctx context.Context, sid string) (string, error) {
	////
	signature, err := s.fetchSid(ctx, sid, consts.KEY_signature)
	if err != nil {
		return "", gerror.NewCode(consts.CodeInternalError)
	}
	if signature == "" {
		return "", gerror.NewCode(consts.CodeInternalError)
	}

	return signature, err
}
func (s *sGenerator) CleanSignature(ctx context.Context, sid string) (string, error) {
	////
	s.recordSid(ctx, sid, consts.KEY_signature, "")
	return "", nil
}

// ///
func (s *sGenerator) FetchTxs(ctx context.Context, sid string) (string, error) {
	////
	signature, err := s.fetchSid(ctx, sid, consts.KEY_txs)
	if err != nil {
		return "", gerror.NewCode(consts.CodeInternalError)
	}
	if signature == "" {
		return "", gerror.NewCode(consts.CodeInternalError)
	}

	return signature, err
}
func (s *sGenerator) RecordTxs(ctx context.Context, sid string, val string) (string, error) {
	////
	s.recordSid(ctx, sid, consts.KEY_txs, val)
	return "", nil
	// if err != nil {
	// 	return "", gerror.NewCode(consts.CodeInternalError)
	// }
	// if signature == "" {
	// 	return "", gerror.NewCode(consts.CodeInternalError)
	// }

	// return signature, err
}
