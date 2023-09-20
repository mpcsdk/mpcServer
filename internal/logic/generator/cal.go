package generator

import (
	"bytes"
	"context"
	"encoding/hex"
	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// GenContextP2
func (s *sGenerator) GenContextP2(ctx context.Context, sid string, private_key2, public_key string, submit bool) error {
	if submit {
		s.pool.Submit(func() {
			s.genContextP2(s.ctx, sid, private_key2, public_key)
		})
	} else {
		s.genContextP2(s.ctx, sid, private_key2, public_key)
	}
	return nil
}

// 4.5.calculate p2_zk_proof by p1_hash_proof, need recal context_p2 by p1_hash_proof
func (s *sGenerator) CalZKProofP2(ctx context.Context, sid string, p1_hash_proof string) error {
	s.pool.Submit(func() {
		s.calZKProofP2(s.ctx, sid, p1_hash_proof)
	})

	return nil
}

// 6.7.calculate v2_public_key by p1_zk_proof, recal context_p2 by p1_zk_proof
func (s *sGenerator) CalPublicKey2(ctx context.Context, sid string, p1_zk_proof string) error {

	s.pool.Submit(func() {
		s.calPublicKey2(s.ctx, sid, p1_zk_proof)
	})

	return nil
}

// 8.calculate request, recal context_p2
func (s *sGenerator) CalRequest(ctx context.Context, sid string, request string) error {
	s.pool.Submit(func() {
		s.pool.Submit(func() {
			s.calRequest(s.ctx, sid, request)
		})
	})

	return nil
}

var prefix = "\x19Ethereum Signed Message:\n"

func (c *sGenerator) hashMessage(ctx context.Context, msg string) string {
	bytemsg, err := hex.DecodeString(msg)
	if err == nil {
		buf := bytes.Buffer{}

		///
		// bytemsg, _ := hex.DecodeString(msg)
		bytelen := strconv.Itoa(len(bytemsg))
		//
		buf.WriteString(prefix)
		buf.WriteString(bytelen)
		buf.WriteString(string(bytemsg))

		hash := crypto.Keccak256Hash(buf.Bytes())

		return hash.Hex()
	} else {
		buf := bytes.Buffer{}

		///
		bytelen := strconv.Itoa(len(msg))
		//
		buf.WriteString(prefix)
		buf.WriteString(bytelen)
		buf.WriteString(msg)

		hash := crypto.Keccak256Hash(buf.Bytes())
		///
		hstr := hash.Hex()
		hstr = strings.TrimPrefix(hstr, "0x")

		return c.hashMessage(ctx, hstr)
	}
}

func (c *sGenerator) digestTxHash(ctx context.Context, SignData string) (string, error) {
	msg, err := service.TxHash().DigestTxHash(ctx, SignData)
	msg = strings.TrimPrefix(msg, "0x")
	return msg, err
}

func (s *sGenerator) CalMsgSign(ctx context.Context, req *v1.SignMsgReq) error {
	hash := s.hashMessage(ctx, req.Msg)
	hash = strings.TrimPrefix(hash, "0x")
	g.Log().Info(ctx, "CalMsgSign:", hash, req.Msg)
	signMsg := hash

	// /////sign
	s.pool.Submit(func() {
		s.CalSignTask(s.ctx, req.SessionId, signMsg, req.Request)
	})
	return nil
}
func (s *sGenerator) CalDomainSign(ctx context.Context, req *v1.SignMsgReq) error {

	///
	msg, err := service.TxHash().TypedDataEncoderHash(ctx, req.SignData)
	if err != nil {

	}
	if msg != "" {

	}
	msg = strings.TrimPrefix(msg, "0x")
	////

	// check domainhash
	hash := s.hashMessage(ctx, msg)
	hash = strings.TrimPrefix(hash, "0x")
	if hash != req.Msg {
		g.Log().Error(ctx, "CalDomainSign unmath", req.SessionId, err, hash)
		return gerror.NewCode(consts.CodeInternalError)
	}

	// /////sign
	s.pool.Submit(func() {
		s.CalSignTask(s.ctx, req.SessionId, req.Msg, req.Request)
	})
	return nil
}

// 9.signature/
// func (s *sGenerator) CheckCalSign(ctx context.Context, req *v1.SignMsgReq) error {
// }
func (s *sGenerator) CalSign(ctx context.Context, req *v1.SignMsgReq) error {
	var err error
	///
	if len(req.Msg) < 10 {
		///impossible
		panic("<10?")
	}
	// checkmsghash
	msg, err := s.digestTxHash(ctx, req.SignData)
	if err != nil {
		//todo:
	}
	hash := s.hashMessage(ctx, msg)
	hash = strings.TrimPrefix(hash, "0x")
	if hash != req.Msg {
		g.Log().Error(ctx, "SignMsg signMsg unmath", req.SessionId, err, hash)
		return gerror.NewCode(consts.CodeInternalError)
	}

	// /////sign
	s.pool.Submit(func() {
		s.CalSignTask(s.ctx, req.SessionId, req.Msg, req.Request)
	})

	return nil
}
