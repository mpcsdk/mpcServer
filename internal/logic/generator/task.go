package generator

import (
	"context"
	"li17server/internal/service"
)

// GenContextP2
func (s *sGenerator) genContextP2(ctx context.Context, sid string, private_key2, public_key string) error {
	p2 := service.Sign().GenContextP2(private_key2, public_key)
	err := s.RecordP2(ctx, sid, p2)
	s.UpGeneratorState(ctx, sid, "p2", err)
	return nil
}

// 1.2.3 cal zk_proof2 by zk_proof1, need recal private_key2_ and p2_context
func (s *sGenerator) calZKProof2(ctx context.Context, sid string, zk_proof1 string) (err error) {
	p2, err := s.FetchP2(ctx, sid)
	if err != nil {
		return
	}
	private_key2_ := service.Sign().RecvZKProofP2(p2, zk_proof1)
	s.RecordPrivateKey2(ctx, sid, private_key2_)

	p2_context := service.Sign().GenContextP2(private_key2_, "")
	s.RecordContextp2(ctx, sid, p2_context)

	zk_proof2 := service.Sign().SendZKProofP2(p2)
	s.RecordZKProof2(ctx, sid, zk_proof2)

	s.UpGeneratorState(ctx, sid, "zk_proof2", err)
	return err
}

// 4.5.calculate p2_zk_proof by p1_hash_proof, need recal p2_context by p1_hash_proof
func (s *sGenerator) calZKProofP2(ctx context.Context, sid string, p1_hash_proof string) error {
	p2_context, err := s.FetchContextp2(ctx, sid)
	p2_context = service.Sign().KeygenRecvHashProofP2(p2_context, p1_hash_proof)
	s.RecordContextp2(ctx, sid, p2_context)
	///
	p2_zk_proof := service.Sign().KeygenSendZKProofP2(p2_context)
	s.RecordZKProofP2(ctx, sid, p2_zk_proof)

	s.UpGeneratorState(ctx, sid, "p2_zk_proof", err)
	return err
}

// 6.7.calculate v2_public_key by p1_zk_proof, recal p2_context by p1_zk_proof
func (s *sGenerator) calPublicKey2(ctx context.Context, sid string, p1_zk_proof string) error {
	p2_context, err := s.FetchContextp2(ctx, sid)
	p2_context = service.Sign().KeygenRecvZKProofP2(p2_context, p1_zk_proof)
	s.RecordContextp2(ctx, sid, p2_context)
	///
	v2_public_key := service.Sign().PublicKeyP2(p2_context)
	s.RecordZKProofP2(ctx, sid, v2_public_key)

	s.UpGeneratorState(ctx, sid, "v2_public_key", err)
	return err
}

// 8.calculate request, recal p2_context
func (s *sGenerator) calRequest(ctx context.Context, sid string, request string) error {
	p2_context, err := s.FetchContextp2(ctx, sid)
	p2_context = service.Sign().KeygenRecvZKProofP2(p2_context, request)
	s.RecordContextp2(ctx, sid, p2_context)

	s.RecordRequest(ctx, sid, request)

	s.UpGeneratorState(ctx, sid, "request", err)
	return err
}

// 9.signature
func (s *sGenerator) calSign(ctx context.Context, sid string, msg string) error {
	p2_context, err := s.FetchContextp2(ctx, sid)
	p2_sign := service.Sign().SignSendPartialP2(p2_context, msg)
	s.RecordSignature(ctx, sid, p2_sign)
	///
	s.UpGeneratorState(ctx, sid, "signature", err)
	return err
}
