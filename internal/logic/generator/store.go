package generator

import (
	"context"
	"li17server/internal/service"
	"time"
)

var duration time.Duration = 0

func (s *sGenerator) UpGeneratorState(ctx context.Context, sid string, state string, err error) error {
	stat := string(state)
	if err != nil {
		stat = stat + ":err:"
		stat += err.Error()
	}
	service.Cache().Set(ctx, sid, stat, duration)
	return nil
}

func (s *sGenerator) GetGeneratorState(ctx context.Context, sid string) (string, error) {
	stat, err := service.Cache().Get(ctx, sid)
	return stat.String(), err
}

func (s *sGenerator) GetStateData(ctx context.Context, sid, state string) (string, error) {
	data, err := service.Cache().Get(ctx, sid+state)
	return data.String(), err
}

// pubkey
func (s *sGenerator) FetchPubKey(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"pubkey")
	return p2.String(), err
}
func (s *sGenerator) RecordPubKey(ctx context.Context, sid string, pubkey string) error {
	err := service.Cache().Set(ctx, sid+"pubkey", pubkey, duration)
	if err == nil {
		service.Generator().UpGeneratorState(ctx, sid, "pubkey", err)
	}
	return err
}

// privatekey
func (s *sGenerator) FetchPrivateKey(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"privatekey")
	return p2.String(), err
}
func (s *sGenerator) RecordPrivateKey(ctx context.Context, sid string, privatekey string) error {
	err := service.Cache().Set(ctx, sid+"privatekey", privatekey, duration)
	if err == nil {
		service.Generator().UpGeneratorState(ctx, sid, "privatekey", err)
	}
	return err
}

// //p2
func (s *sGenerator) FetchP2(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"p2")
	return p2.String(), err
}
func (s *sGenerator) RecordP2(ctx context.Context, sid string, p2 string) error {
	err := service.Cache().Set(ctx, sid+"p2", p2, duration)
	if err == nil {
		service.Generator().UpGeneratorState(ctx, sid, "p2", err)
	}
	return err
}

// //private_key2_
func (s *sGenerator) RecordPrivateKey2(ctx context.Context, sid string, pkey string) error {
	err := service.Cache().Set(ctx, sid+"pkey2", pkey, duration)
	if err == nil {
		service.Generator().UpGeneratorState(ctx, sid, "pkey2", err)
	}
	return err
}
func (s *sGenerator) FetchPrivateKey2(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"pkey2")
	return p2.String(), err
}

// //zk_proof2
func (s *sGenerator) RecordZKProof2(ctx context.Context, sid string, zkproof2 string) error {
	err := service.Cache().Set(ctx, sid+"zk_proof2", zkproof2, duration)
	if err == nil {
		service.Generator().UpGeneratorState(ctx, sid, "zk_proof2", err)
	}
	return err
}
func (s *sGenerator) FetchZKProof2(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"zk_proof2")
	return p2.String(), err
}

// //p2_context
func (s *sGenerator) RecordContextp2(ctx context.Context, sid string, p2_context string) error {
	err := service.Cache().Set(ctx, sid+"p2_context", p2_context, duration)
	if err == nil {
		service.Generator().UpGeneratorState(ctx, sid, "p2_context", err)
	}
	return err
}
func (s *sGenerator) FetchContextp2(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"p2_context")
	return p2.String(), err
}

// //p1_hash_proof
func (s *sGenerator) RecordHashProofP1(ctx context.Context, sid string, hashproofp1 string) error {
	err := service.Cache().Set(ctx, sid+"p1_hash_proof", hashproofp1, duration)
	if err == nil {
		service.Generator().UpGeneratorState(ctx, sid, "p1_hash_proof", err)
	}
	return err
}
func (s *sGenerator) FetchHashProofP1(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"p1_hash_proof")
	return p2.String(), err
}

// //p2_zk_proof
func (s *sGenerator) RecordZKProofP2(ctx context.Context, sid string, p2_zk_proof string) error {
	err := service.Cache().Set(ctx, sid+"p2_zk_proof", p2_zk_proof, duration)
	if err == nil {
		service.Generator().UpGeneratorState(ctx, sid, "p2_zk_proof", err)
	}
	return err
}
func (s *sGenerator) FetchZKProofP2(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"p2_zk_proof")
	return p2.String(), err
}

// //p1_zk_proof
func (s *sGenerator) RecordZKProofP1(ctx context.Context, sid string, p1_zk_proof string) error {
	err := service.Cache().Set(ctx, sid+"p1_zk_proof", p1_zk_proof, duration)
	if err == nil {
		service.Generator().UpGeneratorState(ctx, sid, "p1_zk_proof", err)
	}
	return err
}
func (s *sGenerator) FetchZKProofP1(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"p1_zk_proof")
	return p2.String(), err
}

// v1_public_key
// v2_public_key
func (s *sGenerator) RecordPublicKey2(ctx context.Context, sid string, v2_public_key string) error {
	err := service.Cache().Set(ctx, sid+"v2_public_key", v2_public_key, duration)
	if err == nil {
		service.Generator().UpGeneratorState(ctx, sid, "v2_public_key", err)
	}
	return err
}
func (s *sGenerator) FetchPublicKey2(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"v2_public_key")
	return p2.String(), err
}

// //request
func (s *sGenerator) RecordRequest(ctx context.Context, sid string, request string) error {
	err := service.Cache().Set(ctx, sid+"request", request, duration)
	if err == nil {
		service.Generator().UpGeneratorState(ctx, sid, "request", err)
	}
	return err
}
func (s *sGenerator) FetchRequest(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"request")
	return p2.String(), err
}

// //msg
func (s *sGenerator) RecordMsg(ctx context.Context, sid string, msg string) error {
	err := service.Cache().Set(ctx, sid+"msg", msg, duration)
	if err == nil {
		service.Generator().UpGeneratorState(ctx, sid, "msg", err)
	}
	return err
}
func (s *sGenerator) FetchMsg(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"msg")
	return p2.String(), err
}

// //msg
func (s *sGenerator) RecordSignature(ctx context.Context, sid string, signature string) error {
	err := service.Cache().Set(ctx, sid+"signature", signature, duration)
	if err == nil {
		service.Generator().UpGeneratorState(ctx, sid, "signature", err)
	}
	return err
}
func (s *sGenerator) FetchSignature(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"signature")
	return p2.String(), err
}
