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
	SendHashProof(ctx context.Context, req *v1.SendHashProofReq) (res *v1.SendHashProofRes, err error)
	GetZKProofP2(ctx context.Context, req *v1.GetZKProofP2Req) (res *v1.GetZKProofP2Res, err error)
	SendZKProofP1(ctx context.Context, req *v1.SendZKProofP1Req) (res *v1.SendZKProofP1Res, err error)
	SignMsg(ctx context.Context, req *v1.SignMsgReq) (res *v1.SignMsgRes, err error)
	GetSignature(ctx context.Context, req *v1.GetSignatureReq) (res *v1.GetSignatureRes, err error)
	GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error)
	SendSmsCode(ctx context.Context, req *v1.SendSmsCodeReq) (res *v1.SendSmsCodeRes, err error)
	VerifySms(ctx context.Context, req *v1.VerifySmsCodeReq) (res *v1.VerifySmsCodeRes, err error)
}


