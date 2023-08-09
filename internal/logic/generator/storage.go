package generator

import (
	"context"
	"errors"
	"li17server/internal/service"
	"time"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/os/glog"
	"github.com/yitter/idgenerator-go/idgen"
)

var duration time.Duration = 0
var emptyErr error = errors.New("empty value")

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
	if stat.IsEmpty() {
		return "", emptyErr
	}
	return stat.String(), err
}

func (s *sGenerator) GetStateData(ctx context.Context, sid, state string) (string, error) {
	data, err := service.Cache().Get(ctx, sid+state)
	if data.IsEmpty() {
		return "", emptyErr
	}
	return data.String(), err
}

// pubkey
func (s *sGenerator) FetchPubKey(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"pubkey")
	if p2.IsEmpty() {
		return "", emptyErr
	}
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
	if p2.IsEmpty() {
		return "", emptyErr
	}
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
	p2, err := service.Cache().Get(ctx, sid+"context_p2")
	if p2.IsEmpty() {
		return "", emptyErr
	}
	return p2.String(), err
}
func (s *sGenerator) RecordP2(ctx context.Context, sid string, p2 string) error {
	err := service.Cache().Set(ctx, sid+"context_p2", p2, duration)
	return err
}

// //private_key2_
func (s *sGenerator) RecordPrivateKey2(ctx context.Context, sid string, pkey string) error {
	err := service.Cache().Set(ctx, sid+"pkey2", pkey, duration)

	return err
}
func (s *sGenerator) FetchPrivateKey2(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"pkey2")
	if p2.IsEmpty() {
		return "", emptyErr
	}
	return p2.String(), err
}

// //zk_proof2
func (s *sGenerator) RecordZKProof2(ctx context.Context, sid string, zkproof2 string) error {
	err := service.Cache().Set(ctx, sid+"zk_proof2", zkproof2, duration)

	return err
}
func (s *sGenerator) FetchZKProof2(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"zk_proof2")
	if p2.IsEmpty() {
		return "", emptyErr
	}
	return p2.String(), err
}

// //context_p2
func (s *sGenerator) RecordContextp2(ctx context.Context, sid string, context_p2 string) error {
	err := service.Cache().Set(ctx, sid+"context_p2", context_p2, duration)

	return err
}
func (s *sGenerator) FetchContextp2(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"context_p2")
	if p2.IsEmpty() {
		return "", emptyErr
	}
	return p2.String(), err
}

// //p1_hash_proof
func (s *sGenerator) RecordHashProofP1(ctx context.Context, sid string, hashproofp1 string) error {
	err := service.Cache().Set(ctx, sid+"p1_hash_proof", hashproofp1, duration)

	return err
}
func (s *sGenerator) FetchHashProofP1(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"p1_hash_proof")
	if p2.IsEmpty() {
		return "", emptyErr
	}
	return p2.String(), err
}

// //p2_zk_proof
func (s *sGenerator) RecordZKProofP2(ctx context.Context, sid string, p2_zk_proof string) error {
	err := service.Cache().Set(ctx, sid+"p2_zk_proof", p2_zk_proof, duration)

	return err
}
func (s *sGenerator) FetchZKProofP2(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"p2_zk_proof")
	if p2.IsEmpty() {
		return "", emptyErr
	}
	return p2.String(), err
}

// //p1_zk_proof
func (s *sGenerator) RecordZKProofP1(ctx context.Context, sid string, p1_zk_proof string) error {
	err := service.Cache().Set(ctx, sid+"p1_zk_proof", p1_zk_proof, duration)

	return err
}
func (s *sGenerator) FetchZKProofP1(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"p1_zk_proof")
	if p2.IsEmpty() {
		return "", emptyErr
	}
	return p2.String(), err
}

// v1_public_key
// v2_public_key
func (s *sGenerator) RecordPublicKey2(ctx context.Context, sid string, v2_public_key string) error {
	err := service.Cache().Set(ctx, sid+"v2_public_key", v2_public_key, duration)

	return err
}
func (s *sGenerator) FetchPublicKey2(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"v2_public_key")
	if p2.IsEmpty() {
		return "", emptyErr
	}
	return p2.String(), err
}

// //request
func (s *sGenerator) RecordRequest(ctx context.Context, sid string, request string) error {
	err := service.Cache().Set(ctx, sid+"request", request, duration)

	return err
}

// func (s *sGenerator) FetchRequest(ctx context.Context, sid string) (string, error) {
// 	p2, err := service.Cache().Get(ctx, sid+"request")
// 	if p2.IsEmpty() {
// 		return "", emptyErr
// 	}
// 	return p2.String(), err
// }

// //msg
func (s *sGenerator) RecordMsg(ctx context.Context, sid string, msg string) error {
	err := service.Cache().Set(ctx, sid+"msg", msg, duration)

	return err
}
func (s *sGenerator) FetchMsg(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"msg")
	if p2.IsEmpty() {
		return "", emptyErr
	}
	return p2.String(), err
}

// //msg
func (s *sGenerator) RecordSignature(ctx context.Context, sid string, signature string) error {
	err := service.Cache().Set(ctx, sid+"signature", signature, duration)

	return err
}
func (s *sGenerator) FetchSignature(ctx context.Context, sid string) (string, error) {
	p2, err := service.Cache().Get(ctx, sid+"signature")
	if p2.IsEmpty() {
		return "", emptyErr
	}
	return p2.String(), err
}

// // sid
func (s *sGenerator) GenNewSid(ctx context.Context, userToken string) (string, error) {
	var genid gvar.Var
	genid.Set(idgen.NextId())
	sid := genid.String()
	// bind sid with token
	//todo: duration
	err := service.Cache().Set(ctx, sid, userToken, 0)
	if err != nil {
		glog.Warning(ctx, err)
		return "", err
	}
	return sid, nil
}

func (s *sGenerator) Sid2Token(ctx context.Context, sid string) (string, error) {
	token, err := service.Cache().Get(ctx, sid)
	if token.IsEmpty() {
		return "", emptyErr
	}
	return token.String(), err
}
