package generator

import "fmt"

const (
	STATE_None int = iota
	STATE_HandShake
	STATE_ContextP2
	STATE_ZKProof2
	STATE_ZKProofP2
	STATE_PublicKey2
	STATE_Request
	STATE_Signature
	STATE_Err
)

// const STATE_None string = "none"
// const STATE_HandShake string = "handshake"
// const STATE_ContextP2 string = "p2"
// const STATE_ZKProof2 string = "zk_proof2"
// const STATE_ZKProofP2 string = "p2_zk_proof"
// const STATE_PublicKey2 string = "v2_public_key"
// const STATE_Request string = "request"
// const STATE_Signature string = "signature"

// const STATE_Err string = "error"

type Stater struct {
	curStat int
}

func newStater() *Stater {
	return &Stater{
		curStat: STATE_None,
	}
}

//	func (a *Stater) Next() string {
//		return
//	}
//
//	func (a *Stater) Previous() string {
//		return
//	}
func (a *Stater) Step() {
	if a.curStat == STATE_Err {
		return
	}
	a.curStat = a.curStat + 1

}

func (a *Stater) Current() int {
	return a.curStat
}

func (a *Stater) StateString(state int) string {
	switch a.curStat {
	case STATE_None:
		return "none"
	case STATE_HandShake:
		return "handshake"
	case STATE_ContextP2:
		return "p2"
	case STATE_ZKProof2:
		return "zk_proof2"
	case STATE_ZKProofP2:
		return "p2_zk_proof"
	case STATE_PublicKey2:
		return "v2_public_key"
	case STATE_Request:
		return "request"
	case STATE_Signature:
		return "signature"
	case STATE_Err:
		return "error"
	default:
		return fmt.Sprintf("unknow state:%d", a.curStat)
	}
}
