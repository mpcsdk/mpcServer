package sign

import "li17server/internal/logic/sign/util/li17"

func (a *sSign) KeygenSendHashProofP1(context1 string) string {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	return li17.KeygenSendHashProofP1(context1)
}

func (a *sSign) KeygenRecvHashProofP2(context2, proof1 string) string {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	return li17.KeygenRecvHashProofP2(context2, proof1)
}

func (a *sSign) KeygenSendZKProofP1(context1 string) string {
	return li17.KeygenSendZKProofP1(context1)
}

func (a *sSign) KeygenRecvZKProofP1(context1, proof2 string) string {
	return li17.KeygenRecvZKProofP1(context1, proof2)
}

func (a *sSign) KeygenSendZKProofP2(context1 string) string {
	return li17.KeygenSendZKProofP2(context1)
}
func (a *sSign) KeygenRecvZKProofP2(context2, proof1 string) string {
	return li17.KeygenRecvZKProofP2(context2, proof1)
}

func (a *sSign) PublicKeyP1(context1 string) string {
	return li17.PublicKeyP1(context1)
}
func (a *sSign) PublicKeyP2(context2 string) string {
	return li17.PublicKeyP2(context2)
}
