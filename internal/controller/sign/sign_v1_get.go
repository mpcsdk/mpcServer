package sign

import (
	"context"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"
)

func (c *ControllerV1) GetState(ctx context.Context, req *v1.GetStateReq) (res *v1.GetStateRes, err error) {
	state, err := service.Generator().GetGeneratorState(ctx, req.SessionId)
	data, err := service.Generator().GetStateData(ctx, req.SessionId, state)
	res = &v1.GetStateRes{
		State: state,
		Data:  data,
	}
	return
}

func (c *ControllerV1) GetZKProof2(ctx context.Context, req *v1.GetZKProof2Req) (res *v1.GetZKProof2Res, err error) {
	ZKProof2, err := service.Generator().FetchZKProof2(ctx, req.SessionId)
	res = &v1.GetZKProof2Res{
		Proof2: ZKProof2,
	}
	return
}

func (c *ControllerV1) GetZKProofP2(ctx context.Context, req *v1.GetZKProofP2Req) (res *v1.GetZKProofP2Res, err error) {
	ZKProofp2, err := service.Generator().FetchZKProofP2(ctx, req.SessionId)

	res = &v1.GetZKProofP2Res{
		ZKProofP2: ZKProofp2,
	}
	return
}

func (c *ControllerV1) GetSignature(ctx context.Context, req *v1.GetSignatureReq) (res *v1.GetSignatureRes, err error) {
	signature, err := service.Generator().FetchSignature(ctx, req.SessionId)

	res = &v1.GetSignatureRes{
		Signature: signature,
	}
	return
}
