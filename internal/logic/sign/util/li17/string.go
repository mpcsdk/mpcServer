package li17

import "C"

type Li17String struct {
	Str *C.char
	Len int
}
