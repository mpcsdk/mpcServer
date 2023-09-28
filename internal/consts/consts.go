package consts

const (
	STATE_None int = iota
	STATE_Auth
	STATE_HandShake
)
const (
	RiskCodePass int32 = iota
	RiskCodeNeedVerification
	RiskCodeForbidden
	RiskCodeError
)
