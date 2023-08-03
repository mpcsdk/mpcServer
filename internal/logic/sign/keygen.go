package sign

import "li17server/internal/logic/sign/util/li17"

func (a *sGenerator) KeygenSendHashProofP1(context1 string) string {
	return li17.KeygenSendHashProofP1(context1)
}

func (a *sGenerator) KeygenRecvHashProofP2(context2, proof1 string) string {
	return li17.KeygenRecvHashProofP2(context2, proof1)
}

func (a *sGenerator) KeygenSendZkProofP1(context1 string) string {
	return li17.KeygenSendZkProofP1(context1)
}

func (a *sGenerator) KeygenRecvZkProofP1(context1, proof2 string) string {
	return li17.KeygenRecvZkProofP1(context1, proof2)
}

func (a *sGenerator) KeygenSendZkProofP2(context1 string) string {
	return li17.KeygenSendZkProofP2(context1)
}
func (a *sGenerator) KeygenRecvZkProofP2(context2, proof1 string) string {
	return li17.KeygenRecvZkProofP2(context2, proof1)
}

func (a *sGenerator) PublicKeyP1(context1 string) string {
	return li17.PublicKeyP1(context1)
}
func (a *sGenerator) PublicKeyP2(context2 string) string {
	return li17.PublicKeyP2(context2)
}
