// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sign

import (
	"context"
	
	"li17server/api/sign/v1"
)

type ISignV1 interface {
	AuthUser(ctx context.Context, req *v1.AuthUserReq) (res *v1.AuthUserRes, err error)
	GetState(ctx context.Context, req *v1.GetStateReq) (res *v1.GetStateRes, err error)
	SendZKProof1(ctx context.Context, req *v1.SendZKProof1Req) (res *v1.SendZKProof1Res, err error)
	GetZKProof2(ctx context.Context, req *v1.GetZKProof2Req) (res *v1.GetZKProof2Res, err error)
	SendHashProof(ctx context.Context, req *v1.SendHashProofReq) (res *v1.SendHashProofRes, err error)
	GetZKProofP2(ctx context.Context, req *v1.GetZKProofP2Req) (res *v1.GetZKProofP2Res, err error)
	SendZKProof(ctx context.Context, req *v1.SendZKProofReq) (res *v1.SendZKProofRes, err error)
	SendRequest(ctx context.Context, req *v1.SendRequestReq) (res *v1.SendRequestRes, err error)
	SendMsg(ctx context.Context, req *v1.SendMsgReq) (res *v1.SendMsgRes, err error)
	GetSignature(ctx context.Context, req *v1.GetSignatureReq) (res *v1.GetSignatureRes, err error)
}


