package sign

import (
	"context"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"
)

var privkey2 string = "0ac7d64995c6b4daac2688c0e40d25af50887ada5b7a4cbe197ada0bdef32375"

func (c *ControllerV1) NewSession(ctx context.Context, req *v1.NewSessionReq) (res *v1.NewSessionRes, err error) {
	pubkey := req.PubKey

	res.SessionId = service.Generator().GenContextP2(privkey2, pubkey)
	return
}
func (c *ControllerV1) GetState(ctx context.Context, req *v1.GetStateReq) (res *v1.GetStateRes, err error) {
	return
}
func (c *ControllerV1) GetZKProof2(ctx context.Context, req *v1.GetZKProof2Req) (res *v1.GetZKProof2Res, err error) {
	return
}
func (c *ControllerV1) SendHashProof(ctx context.Context, req *v1.SendHashProofReq) (res *v1.SendHashProofRes, err error) {
	return
}
func (c *ControllerV1) GetZkProofP2(ctx context.Context, req *v1.GetZkProofP2Req) (res *v1.GetZkProofP2Res, err error) {
	return
}
func (c *ControllerV1) SendZkProof(ctx context.Context, req *v1.SendZkProofReq) (res *v1.SendZkProofRes, err error) {
	return
}
func (c *ControllerV1) SendRequest(ctx context.Context, req *v1.SendRequestReq) (res *v1.SendRequestRes, err error) {
	return
}
func (c *ControllerV1) SendMsg(ctx context.Context, req *v1.SendMsgReq) (res *v1.SendMsgRes, err error) {
	return
}
func (c *ControllerV1) GetSignature(ctx context.Context, req *v1.GetSignatureReq) (res *v1.GetSignatureRes, err error) {
	return
}
