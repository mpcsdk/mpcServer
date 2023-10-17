package mpcsigner

import (
	"fmt"
	"li17server/internal/consts"
)

func (s *sMpcSigner) StateString(state int) string {
	switch state {
	case consts.STATE_None:
		return "none"
	case consts.STATE_Auth:
		return "auth"
	case consts.STATE_HandShake:
		return "handshake"
	// case consts.STATE_Done:
	// 	return "done"
	// case consts.STATE_Err:
	// 	return "error"
	default:
		return fmt.Sprintf("unknow state:%d", state)
	}
}
