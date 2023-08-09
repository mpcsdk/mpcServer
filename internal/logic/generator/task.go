package generator

import (
	"context"
	"li17server/internal/service"
)

// GenContextP2
func (s *sGenerator) genContextP2(ctx context.Context, sid string, private_key2, public_key string) error {
	p2 := service.Sign().GenContextP2(private_key2, public_key)
	err := s.RecordP2(ctx, sid, p2)
	return err
}

// 1.2.3 cal zk_proof2 by zk_proof1, need recal private_key2_ and context_p2
func (s *sGenerator) calZKProof2(ctx context.Context, sid string, zk_proof1 string) (err error) {
	p2, err := s.FetchP2(ctx, sid)
	if err != nil {
		return
	}
	private_key2_ := service.Sign().RecvZKProofP2(p2, zk_proof1)
	s.RecordPrivateKey2(ctx, sid, private_key2_)

	context_p2 := service.Sign().GenContextP2(private_key2_, "")
	s.RecordContextp2(ctx, sid, context_p2)

	zk_proof2 := service.Sign().SendZKProofP2(p2)
	s.RecordZKProof2(ctx, sid, zk_proof2)

	return err
}

// 4.5.calculate p2_zk_proof by p1_hash_proof, need recal context_p2 by p1_hash_proof
func (s *sGenerator) calZKProofP2(ctx context.Context, sid string, p1_hash_proof string) error {
	context_p2, err := s.FetchContextp2(ctx, sid)
	context_p2 = service.Sign().KeygenRecvHashProofP2(context_p2, p1_hash_proof)
	s.RecordContextp2(ctx, sid, context_p2)
	///
	p2_zk_proof := service.Sign().KeygenSendZKProofP2(context_p2)
	s.RecordZKProofP2(ctx, sid, p2_zk_proof)
	// s.UpGeneratorState(ctx, sid, s.StateString(service.STATE_HandShake), err)
	return err
}

// 6.7.calculate v2_public_key by p1_zk_proof, recal context_p2 by p1_zk_proof
func (s *sGenerator) calPublicKey2(ctx context.Context, sid string, p1_zk_proof string) error {
	context_p2, err := s.FetchContextp2(ctx, sid)
	context_p2 = service.Sign().KeygenRecvZKProofP2(context_p2, p1_zk_proof)
	s.RecordContextp2(ctx, sid, context_p2)
	///
	v2_public_key := service.Sign().PublicKeyP2(context_p2)
	s.RecordPublicKey2(ctx, sid, v2_public_key)

	s.UpGeneratorState(ctx, sid, s.StateString(service.STATE_HandShake), err)
	return err
}

// 8.calculate request, recal context_p2
func (s *sGenerator) calRequest(ctx context.Context, sid string, request string) (string, error) {
	context_p2, err := s.FetchContextp2(ctx, sid)
	context_p2 = service.Sign().SignRecvRequestP2(context_p2, request)
	s.RecordContextp2(ctx, sid, context_p2)

	s.RecordRequest(ctx, sid, request)

	return context_p2, err
}

// 9.signature
func (s *sGenerator) CalSignTask(ctx context.Context, sid string, msg string, request string) error {
	context_p2, err := s.FetchContextp2(ctx, sid)
	if request != "" {
		context_p2, err = s.calRequest(ctx, sid, request)
		if err != nil {
			return err
		}
	}
	p2_sign := service.Sign().SignSendPartialP2(context_p2, msg)
	s.RecordSignature(ctx, sid, p2_sign)
	///
	s.UpGeneratorState(ctx, sid, s.StateString(service.STATE_HandShake), err)
	return err
}
