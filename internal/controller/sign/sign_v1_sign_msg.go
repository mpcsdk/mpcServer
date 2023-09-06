package sign

import (
	"bytes"
	"context"
	"encoding/hex"
	"strconv"
	"strings"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	///
	// "github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var prefix = "\x19Ethereum Signed Message:\n"

func (c *ControllerV1) hashMessage(ctx context.Context, SignData string) string {
	buf := bytes.Buffer{}
	msg := service.TxHash().DigestTxHash(ctx, SignData)
	msg = strings.TrimPrefix(msg, "0x")
	///
	bytemsg, _ := hex.DecodeString(msg)
	bytelen := strconv.Itoa(len(bytemsg))
	//
	buf.WriteString(prefix)
	buf.WriteString(bytelen)
	buf.WriteString(string(bytemsg))

	// msg = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(bytemsg), string(bytemsg))
	// hasher := sha3.NewLegacyKeccak256()
	// hasher.Write([]byte(msg))
	hash := crypto.Keccak256Hash(buf.Bytes())

	return hash.Hex()
}
func (c *ControllerV1) SignMsg(ctx context.Context, req *v1.SignMsgReq) (res *v1.SignMsgRes, err error) {
	g.Log().Debug(ctx, "SignMsg:", req)
	//todo: checksid
	_, err = service.Generator().Sid2Token(ctx, req.SessionId)
	if err != nil {
		g.Log().Error(ctx, "SignMsg no sid", req.SessionId, err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	// checkmsghash
	hash := c.hashMessage(ctx, req.SignData)
	hash = strings.TrimPrefix(hash, "0x")
	if hash != req.Msg {
		g.Log().Error(ctx, "SignMsg signMsg unmath", req.SessionId, err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	///

	//todo: nocheckrule
	err = service.Generator().CalSign(ctx, req, false) //, req.SessionId, req.Msg, req.Request, req.SignData)
	if err != nil {
		g.Log().Warning(ctx, "SignMsg:", err)
		return nil, err
	}

	return
}
