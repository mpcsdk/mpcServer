package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ///
type AuthUserReq struct {
	g.Meta    `path:"/AuthUser" tags:"AuthUser" method:"post" summary:"AuthUser"`
	UserToken string `json:"userToken"`
}
type AuthUserRes struct {
	SessionId string `json:"sessionId"`
}

// ///
type GetStateReq struct {
	g.Meta    `path:"/GetState" tags:"GetState" method:"post" summary:"GetState"`
	SessionId string `json:"sessionId"`
}
type GetStateRes struct {
	State string `json:"state"`
}

// ///
type SendHashProofReq struct {
	g.Meta    `path:"/SendHashProof" tags:"SendHashProof" method:"post" summary:"SendHashProof"`
	SessionId string `json:"sessionId"`
	HashProof string `json:"hashProof"`
}
type SendHashProofRes struct {
}

// ///
type GetZKProofP2Req struct {
	g.Meta    `path:"/GetZKProofP2" tags:"GetZKProofP2" method:"post" summary:"GetZKProofP2"`
	SessionId string `json:"sessionId"`
}
type GetZKProofP2Res struct {
	ZKProofP2 string `json:"zkProofP2"`
}

// ///
type SendZKProofP1Req struct {
	g.Meta    `path:"/SendZKProofP1" tags:"SendZKProofP1" method:"post" summary:"SendZKProofP1"`
	SessionId string `json:"sessionId"`
	ZKProofP1 string `json:"zkProofP1"`
}
type SendZKProofP1Res struct {
}

// ///
type SignMsgReq struct {
	g.Meta    `path:"/SignMsg" tags:"SendMsg" method:"post" summary:"SendMsg"`
	SessionId string `json:"sessionId"`
	Msg       string `json:"msg"`
	Request   string `json:"request,omitempty"`
	SignData  string `json:"signData,omitempty"`
}
type SignTx struct {
	ChainId uint64        `json:"chainId,omitempty"`
	Address string        `json:"address,omitempty"`
	Number  uint64        `json:"number,omitempty"`
	Txs     []*SignTxData `json:"txs,omitempty"`
	TxHash  string        `json:"txHash,omitempty"`
}
type SignTxData struct {
	To   string `json:"target,omitempty"`
	From string `json:"from,omitempty"`
	Data string `json:"data,omitempty"`
}
type SignMsgRes struct {
}

// ///
type GetSignatureReq struct {
	g.Meta    `path:"/GetSignature" tags:"GetSignature" method:"post" summary:"GetSignature"`
	SessionId string `json:"sessionId"`
}
type GetSignatureRes struct {
	Signature string `json:"signature"`
}

// ///
type GetInfoReq struct {
	g.Meta    `path:"/GetInfo" tags:"GetInfo" method:"post" summary:"GetInfo"`
	SessionId string `json:"sessionId"`
}
type GetInfoRes struct {
	PublicKey string `json:"publicKey"`
}

// //
type SendSmsCodeReq struct {
	g.Meta    `path:"/SendSmsCode" tags:"SendSmsCode" method:"post" summary:"SendSmsCode"`
	SessionId string `json:"sessionId"`
}
type SendSmsCodeRes struct {
}

// /
type VerifySmsCodeReq struct {
	g.Meta    `path:"/VerifySmsCode" tags:"VerifySmsCode" method:"post" summary:"VerifySmsCode"`
	SessionId string `json:"sessionId"`
	Code      string `json:"code"`
}
type VerifySmsCodeRes struct {
}
