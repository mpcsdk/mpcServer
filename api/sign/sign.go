// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sign

import (
	"context"
	
	"li17server/api/sign/v1"
)

type ISignV1 interface {
	NewSession(ctx context.Context, req *v1.NewSessionReq) (res *v1.NewSessionRes, err error)
	GetState(ctx context.Context, req *v1.GetStateReq) (res *v1.GetStateRes, err error)
	GetZKProof2(ctx context.Context, req *v1.GetZKProof2Req) (res *v1.GetZKProof2Res, err error)
	SendHashProof(ctx context.Context, req *v1.SendHashProofReq) (res *v1.SendHashProofRes, err error)
	GetZkProofP2(ctx context.Context, req *v1.GetZkProofP2Req) (res *v1.GetZkProofP2Res, err error)
	SendZkProof(ctx context.Context, req *v1.SendZkProofReq) (res *v1.SendZkProofRes, err error)
	SendRequest(ctx context.Context, req *v1.SendRequestReq) (res *v1.SendRequestRes, err error)
	SendMsg(ctx context.Context, req *v1.SendMsgReq) (res *v1.SendMsgRes, err error)
	GetSignature(ctx context.Context, req *v1.GetSignatureReq) (res *v1.GetSignatureRes, err error)
}


