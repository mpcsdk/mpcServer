package generator

import (
	"fmt"
	"li17server/internal/consts"
)

type Stater struct {
	curStat int
}

func newStater() *Stater {
	return &Stater{
		curStat: consts.STATE_None,
	}
}

func (s *sGenerator) StateNext(state int) int {
	return state + 1
}
func (s *sGenerator) StatePrivate(state int) int {
	if state > consts.STATE_HandShake {
		return state - 1
	}
	return state
}

func (s *sGenerator) StateInt(state string) int {
	switch state {
	case "none":
		return consts.STATE_None
	case "auth":
		return consts.STATE_Auth
	case "handshake":
		return consts.STATE_HandShake
	case "done":
		return consts.STATE_Done
	case "error":
		return consts.STATE_Err
	default:
		return consts.STATE_Err
	}
}

func (s *sGenerator) StateString(state int) string {
	switch state {
	case consts.STATE_None:
		return "none"
	case consts.STATE_Auth:
		return "auth"
	case consts.STATE_HandShake:
		return "handshake"
	case consts.STATE_Done:
		return "done"
	case consts.STATE_Err:
		return "error"
	default:
		return fmt.Sprintf("unknow state:%d", state)
	}
}

func (a *Stater) Step() {
	if a.curStat == consts.STATE_Err || a.curStat == consts.STATE_Done {
		return
	}
	a.curStat = a.curStat + 1

}

func (s *sGenerator) StateIs(state string, istate int) bool {
	if state == s.StateString(istate) {
		return true
	}
	return false
}
func (s *sGenerator) NextStateIs(curstate string) int {
	return s.StateNext(s.StateInt(curstate))
}

/////
