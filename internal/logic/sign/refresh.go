package sign

import "li17server/internal/logic/sign/util/li17"

func (a *sGenerator) GenContextP1(preivateKey, publicKey string) string {
	return li17.GenContextP1(preivateKey, publicKey)
}

func (a *sGenerator) GenContextP2(preivateKey, publicKey string) string {
	return li17.GenContextP2(preivateKey, publicKey)
}

func (a *sGenerator) SendZkProofP1(p1 string) string {
	return li17.SendZkProofP1(p1)
}

func (a *sGenerator) RecvZkProofP1(p1, zkproof2 string) string {
	return li17.RecvZkProofP1(p1, zkproof2)
}

func (a *sGenerator) SendZkProofP2(p2 string) string {
	return li17.SendZkProofP2(p2)
}

func (a *sGenerator) RecvZkProofP2(p2, zkproof1 string) string {
	return li17.RecvZkProofP2(p2, zkproof1)
}
