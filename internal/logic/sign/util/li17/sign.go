package li17

//#cgo CFLAGS: -I./libs/
//#cgo LDFLAGS: -L${SRCDIR}/libs -lli17
//
//#include "li17.h"
//#include <stdlib.h>
import "C"
import "unsafe"

func SignSendRequestP1(context1 string) string {

	c_context1 := C.CString(context1)
	var cptr = C.li17_p1_signature_send_signature_request(c_context1)

	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}

func SignRecvRequestP2(context2 string, request string) string {

	c_context2 := C.CString(context2)
	c_request := C.CString(request)
	var cptr = C.li17_p2_signature_recv_signature_request(c_context2, c_request)

	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}

func SignSendPartialP2(context2, msg string) string {
	c_context2 := C.CString(context2)
	c_msg := C.CString(msg)
	var cptr = C.li17_p2_signature_send_signature_partial(c_context2, c_msg)

	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}

func SignSendPartialP1(context1, sign2, msg string) string {
	c_context1 := C.CString(context1)
	c_sign := C.CString(sign2)
	c_msg := C.CString(msg)
	var cptr = C.li17_p1_signature_recv_signature_partial(c_context1, c_sign, c_msg)

	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}
