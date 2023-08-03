package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type NewSessionReq struct {
	g.Meta `path:"/newSession" tags:"NewSession" method:"get" summary:"NewSession"`
	UserId string `name:"userId"`
	PubKey string `name:"pubKey"`
	Proof  string `name:"proof"`
}
type NewSessionRes struct {
	// g.Meta `mime:"text/html" example:"string"`
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
type GetZkProofP2Req struct {
	g.Meta    `path:"/GetZkProofP2" tags:"GetZkProofP2" method:"post" summary:"GetZkProofP2"`
	SessionId string `name:"sessionId"`
}
type GetZkProofP2Res struct {
	ZkProofP2 string `name:"zkProofP2"`
}

// ///
type SendZkProofReq struct {
	g.Meta    `path:"/SendZkProof" tags:"SendZkProof" method:"post" summary:"SendZkProof"`
	SessionId string `name:"sessionId"`
	ZkProof   string `name:"zkProof"`
}
type SendZkProofRes struct {
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
