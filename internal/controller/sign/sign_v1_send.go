package sign

import (
	"context"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"
)

// recv zk_proof1 for cal zk_proof2
func (c *ControllerV1) SendZKProof1(ctx context.Context, req *v1.SendZKProof1Req) (res *v1.SendZKProof1Res, err error) {
	// todo: check sid
	sid := req.SessionId
	_, err = service.Cache().Get(ctx, sid)
	if err != nil {
		return
	}
	//
	ZKProof1 := req.ZKProof1
	err = service.Generator().CalZKProof2(ctx, sid, ZKProof1)

	return
}

// recv p1_hash_proof for cal p2_zk_proof
func (c *ControllerV1) SendHashProof(ctx context.Context, req *v1.SendHashProofReq) (res *v1.SendHashProofRes, err error) {
	sid := req.SessionId
	err = service.Generator().CalZKProofP2(ctx, sid, req.HashProof)

	return
}

// recv zk_proof for cal v2_public_key and update p2_context
func (c *ControllerV1) SendZKProof(ctx context.Context, req *v1.SendZKProofReq) (res *v1.SendZKProofRes, err error) {
	sid := req.SessionId
	err = service.Generator().CalPublicKey2(ctx, sid, req.ZKProof)

	return
}

func (c *ControllerV1) SendRequest(ctx context.Context, req *v1.SendRequestReq) (res *v1.SendRequestRes, err error) {
	sid := req.SessionId
	err = service.Generator().CalRequest(ctx, sid, req.Request)
	return
}

func (c *ControllerV1) SendMsg(ctx context.Context, req *v1.SendMsgReq) (res *v1.SendMsgRes, err error) {
	sid := req.SessionId
	err = service.Generator().CalSign(ctx, sid, req.Msg)
	return
}
