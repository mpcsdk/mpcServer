package mpcsigner

import (
	"bytes"
	"context"
	"encoding/hex"
	v1 "mpcServer/api/sign/v1"
	"mpcServer/internal/consts"
	"mpcServer/internal/service"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// GenContextP2
func (s *sMpcSigner) GenContextP2(ctx context.Context, sid string, private_key2, public_key string, submit bool) error {
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
func (s *sMpcSigner) CalZKProofP2(ctx context.Context, sid string, p1_hash_proof string) error {
	s.pool.Submit(func() {
		s.calZKProofP2(s.ctx, sid, p1_hash_proof)
	})

	return nil
}

// 6.7.calculate v2_public_key by p1_zk_proof, recal context_p2 by p1_zk_proof
func (s *sMpcSigner) CalPublicKey2(ctx context.Context, sid string, p1_zk_proof string) error {

	s.pool.Submit(func() {
		if err := s.calPublicKey2(s.ctx, sid, p1_zk_proof); err != nil {
			g.Log().Error(s.ctx, "CalPublicKey2:", err)
		}
	})

	return nil
}

// 8.calculate request, recal context_p2
func (s *sMpcSigner) CalRequest(ctx context.Context, sid string, request string) error {
	// s.pool.Submit(func() {
	// 	s.calRequest(s.ctx, sid, request)
	// })

	// return nil
	_, err := s.calRequest(s.ctx, sid, request)
	return err
}

var prefix = "\x19Ethereum Signed Message:\n"

func (c *sMpcSigner) hashMsg(ctx context.Context, msg string) string {
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

	return hstr
}
func (c *sMpcSigner) hashMessage(ctx context.Context, msg string) string {
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

		return hstr
	}
}

func (c *sMpcSigner) digestTxHash(ctx context.Context, SignData string) (string, error) {
	msg, err := service.TxHash().DigestTxHash(ctx, SignData)
	msg = strings.TrimPrefix(msg, "0x")
	return msg, err
}

func (s *sMpcSigner) CalMsgSign(ctx context.Context, req *v1.SignMsgReq) error {
	hmsg := s.hashMsg(ctx, req.Msg)
	hmsg = strings.TrimPrefix(hmsg, "0x")
	///
	hash := s.hashMessage(ctx, hmsg)
	hash = strings.TrimPrefix(hash, "0x")
	//
	g.Log().Info(ctx, "CalMsgSign:", hash, req.Msg)
	signMsg := hash

	// /////sign
	s.pool.Submit(func() {
		s.CalSignTask(s.ctx, req.SessionId, signMsg, req.Request)
	})
	return nil
}
func (s *sMpcSigner) CalDomainSign(ctx context.Context, req *v1.SignMsgReq) error {

	///
	hash, err := service.TxHash().TypedDataEncoderHash(ctx, req.SignData)
	if err != nil {
		return err
	}
	// check domainhash
	// msg := s.hashMessage(ctx, hash)
	msg := strings.TrimPrefix(hash, "0x")
	if msg != req.Msg {
		g.Log().Error(ctx, "CalDomainSign unmath", req.SessionId, err, msg, req.Msg)
		return gerror.NewCode(consts.CodeInternalError)
	}

	// /////sign
	s.pool.Submit(func() {
		s.CalSignTask(s.ctx, req.SessionId, req.Msg, req.Request)
	})
	return nil
}

// 9.signature/
// func (s *sMpcSigner) CheckCalSign(ctx context.Context, req *v1.SignMsgReq) error {
// }
func (s *sMpcSigner) CalSign(ctx context.Context, req *v1.SignMsgReq) error {
	var err error
	///
	if len(req.Msg) < 10 {
		///impossible
		panic("<10?")
	}
	// checkmsghash
	msg, err := s.digestTxHash(ctx, req.SignData)
	if err != nil {
		g.Log().Warning(ctx, "CalSign digestTxHash err", err)
		return gerror.NewCode(consts.CodeInternalError)
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
