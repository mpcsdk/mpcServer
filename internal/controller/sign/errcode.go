package sign

type errCode struct {
	code    int
	message string
	detail  interface{}
}

func (e *errCode) Message() string {
	return ""
}
func (e *errCode) Code() int {
	return 1
}
func (e *errCode) Detail() interface{} {
	return ""
}

var (
	CodeNil           = &errCode{-1, "", nil}               // No error code specified.
	CodeOK            = &errCode{0, "OK", nil}              // It is OK.
	CodeInternalError = &errCode{50, "Internal Error", nil} // An error occurred internally.
)
