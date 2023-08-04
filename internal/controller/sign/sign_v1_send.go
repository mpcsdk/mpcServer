package sign

import (
	"context"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"
)

// recv p1_hash_proof for cal p2_zk_proof
func (c *ControllerV1) SendHashProof(ctx context.Context, req *v1.SendHashProofReq) (res *v1.SendHashProofRes, err error) {
	//todo: check sid
	sid := req.SessionId

	state, err := service.Generator().GetGeneratorState(ctx, sid)
	if state != service.Generator().StateString(service.STATE_None) {
		// todo: must none state
		return
	}
	err = service.Generator().CalZKProofP2(ctx, sid, req.HashProof)

	return
}

func (c *ControllerV1) SendZKProofP1(ctx context.Context, req *v1.SendZKProofP1Req) (res *v1.SendZKProofP1Res, err error) {
	// todo: check sid and p2_zk_proof
	sid := req.SessionId
	_, err = service.Cache().Get(ctx, sid)
	if err != nil {
		return
	}

	state, err := service.Generator().GetGeneratorState(ctx, sid)
	if state != service.Generator().StateString(service.STATE_None) {
		// todo: must none state
		return
	}

	_, err = service.Generator().FetchZKProofP2(ctx, sid)
	if err != nil {
		// todo: no zkproofp2
		return
	}

	err = service.Generator().CalPublicKey2(ctx, sid, req.ZKProofP1)
	/// todo: publicKeyV1
	return
}
