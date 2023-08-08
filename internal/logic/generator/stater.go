package generator

import (
	"fmt"
	"li17server/internal/service"
)

// const (
// 	STATE_None int = iota
// 	STATE_Auth
// 	STATE_HandShake_Request
// 	STATE_HandShake_NoRequest
// 	STATE_Done
// 	STATE_Err
// )

type Stater struct {
	curStat int
}

func newStater() *Stater {
	return &Stater{
		curStat: service.STATE_None,
	}
}

func (s *sGenerator) StateNext(state int) int {
	return state + 1
}
func (s *sGenerator) StatePrivate(state int) int {
	if state > service.STATE_HandShake {
		return state - 1
	}
	return state
}

func (s *sGenerator) StateInt(state string) int {
	switch state {
	case "none":
		return service.STATE_None
	case "auth":
		return service.STATE_Auth
	case "handshake":
		return service.STATE_HandShake
	case "done":
		return service.STATE_Done
	case "error":
		return service.STATE_Err
	default:
		return service.STATE_Err
	}
}

func (s *sGenerator) StateString(state int) string {
	switch state {
	case service.STATE_None:
		return "none"
	case service.STATE_Auth:
		return "auth"
	case service.STATE_HandShake:
		return "handshake"
	case service.STATE_Done:
		return "done"
	case service.STATE_Err:
		return "error"
	default:
		return fmt.Sprintf("unknow state:%d", state)
	}
}

func (a *Stater) Step() {
	if a.curStat == service.STATE_Err || a.curStat == service.STATE_Done {
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
