package sign

import "li17server/internal/logic/sign/util/li17"

func (a *sSign) GenContextP1(preivateKey, publicKey string) string {
	return li17.GenContextP1(preivateKey, publicKey)
}

func (a *sSign) GenContextP2(preivateKey, publicKey string) string {
	return li17.GenContextP2(preivateKey, publicKey)
}

func (a *sSign) SendZKProofP1(p1 string) string {
	return li17.SendZKProofP1(p1)
}

func (a *sSign) RecvZKProofP1(p1, ZKProof2 string) string {
	return li17.RecvZKProofP1(p1, ZKProof2)
}

func (a *sSign) SendZKProofP2(p2 string) string {
	return li17.SendZKProofP2(p2)
}

func (a *sSign) RecvZKProofP2(p2, ZKProof1 string) string {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	return li17.RecvZKProofP2(p2, ZKProof1)
}
