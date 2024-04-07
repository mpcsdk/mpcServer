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
	defer C.free(unsafe.Pointer(c_context1))
	var request = C.li17_p1_signature_send_signature_request(c_context1)

	return (C.GoString(request))
}

func SignRecvRequestP2(context2 string, request string) string {

	c_context2 := C.CString(context2)
	defer C.free(unsafe.Pointer(c_context2))
	c_request := C.CString(request)
	defer C.free(unsafe.Pointer(c_request))
	var p2 = C.li17_p2_signature_recv_signature_request(c_context2, c_request)

	return (C.GoString(p2))
}

func SignSendPartialP2(context2, msg string) string {
	c_context2 := C.CString(context2)
	defer C.free(unsafe.Pointer(c_context2))
	c_msg := C.CString(msg)
	defer C.free(unsafe.Pointer(c_msg))
	var p2sign = C.li17_p2_signature_send_signature_partial(c_context2, c_msg)

	return (C.GoString(p2sign))
}

func SignSendPartialP1(context1, sign2, msg string) string {
	c_context1 := C.CString(context1)
	defer C.free(unsafe.Pointer(c_context1))
	c_sign := C.CString(sign2)
	defer C.free(unsafe.Pointer(c_sign))
	c_msg := C.CString(msg)
	defer C.free(unsafe.Pointer(c_msg))
	var sign = C.li17_p1_signature_recv_signature_partial(c_context1, c_sign, c_msg)

	return (C.GoString(sign))
}
