package sign

type errCode struct {
	code    int
	message string
	detail  interface{}
}

func (e *errCode) Message() string {
	return e.message
}
func (e *errCode) Code() int {
	return e.code
}
func (e *errCode) Detail() interface{} {
	return e.detail
}

var (
	CodeNil = &errCode{-1, "", nil}  // No error code specified.
	CodeOK  = &errCode{0, "ok", nil} // It is OK.

	CodeInternalError = &errCode{50, "Internal Error", nil} // An error occurred internally.

)

const ErrSessionNotExist string = "sessionId not exist"
const ErrStateIncorrect string = "state is incorrect"

const ErrZKProofP2NotExist string = "zkproof2 not exist"
const ErrSignatureNotExist string = "signature not exist"

func CodeStateError(msg string) *errCode {

	return &errCode{1, msg, nil}
}

// //
func CodeGetGeneratorError(msg string) *errCode {
	return &errCode{2, msg, nil}
}

// //
// /
func CalZKProofP2Error(msg string) *errCode {

	return &errCode{11, "failed to cal zkproof2", nil}
}

func CalPublicKey2Error(msg string) *errCode {

	return &errCode{12, "failed to cal PublicKey2", nil}
}

func CalSignError(msg string) *errCode {

	return &errCode{13, "failed to sign", nil}
}
func NeedSmsCodeError(msg string) *errCode {
	return &errCode{1, "send smscode", nil}
}
func SmsCodeError(msg string) *errCode {
	return &errCode{2, "smscode code", nil}
}
