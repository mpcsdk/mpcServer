package generator

import (
	"context"
)

// GenContextP2
func (s *sGenerator) GenContextP2(ctx context.Context, token string, private_key2, public_key string, submit bool) error {
	if submit {
		s.pool.Submit(func() {
			s.genContextP2(s.ctx, token, private_key2, public_key)
		})
	} else {
		s.genContextP2(s.ctx, token, private_key2, public_key)
	}
	return nil
}

// 1.2.3 cal zk_proof2 by zk_proof1, need recal private_key2_ and context_p2
func (s *sGenerator) CalZKProof2(ctx context.Context, token string, zk_proof1 string) (err error) {

	s.pool.Submit(func() {
		s.calZKProof2(s.ctx, token, zk_proof1)
	})

	return
}

// 4.5.calculate p2_zk_proof by p1_hash_proof, need recal context_p2 by p1_hash_proof
func (s *sGenerator) CalZKProofP2(ctx context.Context, token string, p1_hash_proof string) error {
	s.pool.Submit(func() {
		s.calZKProofP2(s.ctx, token, p1_hash_proof)
	})

	return nil
}

// 6.7.calculate v2_public_key by p1_zk_proof, recal context_p2 by p1_zk_proof
func (s *sGenerator) CalPublicKey2(ctx context.Context, token string, p1_zk_proof string) error {

	s.pool.Submit(func() {
		s.calPublicKey2(s.ctx, token, p1_zk_proof)
	})

	return nil
}

// 8.calculate request, recal context_p2
func (s *sGenerator) CalRequest(ctx context.Context, token string, request string) error {
	s.pool.Submit(func() {
		s.pool.Submit(func() {
			s.calRequest(s.ctx, token, request)
		})
	})

	return nil
}

// 9.signature
func (s *sGenerator) CalSign(ctx context.Context, token string, msg string, request string) error {
	// s.pool.Submit(func() {
	// 	s.CalSignTask(s.ctx, sid, msg, request)
	// })

	s.CalSignTask(s.ctx, token, msg, request)
	return nil
}
