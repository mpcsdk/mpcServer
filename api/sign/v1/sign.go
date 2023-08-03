package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ///
type AuthUserReq struct {
	g.Meta    `path:"/AuthUser" tags:"AuthUser" method:"post" summary:"AuthUser"`
	UserToken string `name:"userToken"`
	PubKey    string `name:"pubKey"`
}
type AuthUserRes struct {
	SessionId string `name:"sessionId"`
}

// ///
type GetStateReq struct {
	g.Meta    `path:"/GetState" tags:"GetState" method:"post" summary:"GetState"`
	SessionId string `name:"sessionId"`
}
type GetStateRes struct {
	State string `name:"state"`
}

// //
type SendZKProof1Req struct {
	g.Meta    `path:"/SendZKProof1" tags:"SendZKProof1" method:"post" summary:"SendZKProof1"`
	SessionId string `name:"sessionId"`
	ZKProof1  string `name:"ZKProof1"`
}
type SendZKProof1Res struct {
	// g.Meta `mime:"text/html" example:"string"`
}

// ///
type GetZKProof2Req struct {
	g.Meta    `path:"/GetZKProof2" tags:"GetZKProof2" method:"post" summary:"GetZKProof2"`
	SessionId string `name:"sessionId"`
}
type GetZKProof2Res struct {
	Proof2 string `name:"proof2"`
}

// ///
type SendHashProofReq struct {
	g.Meta    `path:"/SendHashProof" tags:"SendHashProof" method:"post" summary:"SendHashProof"`
	SessionId string `name:"sessionId"`
	HashProof string `name:"hashProof"`
}
type SendHashProofRes struct {
}

// ///
type GetZKProofP2Req struct {
	g.Meta    `path:"/GetZKProofP2" tags:"GetZKProofP2" method:"post" summary:"GetZKProofP2"`
	SessionId string `name:"sessionId"`
}
type GetZKProofP2Res struct {
	ZKProofP2 string `name:"ZKProofP2"`
}

// ///
type SendZKProofReq struct {
	g.Meta    `path:"/SendZKProof" tags:"SendZKProof" method:"post" summary:"SendZKProof"`
	SessionId string `name:"sessionId"`
	ZKProof   string `name:"ZKProof"`
}
type SendZKProofRes struct {
}

// ///
type SendRequestReq struct {
	g.Meta    `path:"/SendRequest" tags:"SendRequest" method:"post" summary:"SendRequest"`
	SessionId string `name:"sessionId"`
	Request   string `name:"request"`
}
type SendRequestRes struct {
}

// ///
type SendMsgReq struct {
	g.Meta    `path:"/SendMsg" tags:"SendMsg" method:"post" summary:"SendMsg"`
	SessionId string `name:"sessionId"`
	Msg       string `name:"msg"`
}
type SendMsgRes struct {
}

// ///
type GetSignatureReq struct {
	g.Meta    `path:"/GetSignature" tags:"GetSignature" method:"post" summary:"GetSignature"`
	SessionId string `name:"sessionId"`
}
type GetSignatureRes struct {
	Signature string `name:"signature"`
}
