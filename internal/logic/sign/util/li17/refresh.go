package li17

//#cgo CFLAGS: -I./libs/
//#cgo LDFLAGS: -L${SRCDIR}/libs -lli17
//
//#include "li17.h"
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

// func CString2G(s *C.char) string {
// 	gstr := *(*string)(unsafe.Pointer(&s))
// 	len := strings.IndexByte(gstr, byte(0))
// 	ustr := *(*Li17String)(unsafe.Pointer(&s))
// 	ustr.Len = len
// 	gstr = *(*string)(unsafe.Pointer(&ustr))
// 	return gstr
// }

func GenContextP1(preivateKey, publicKey string) string {
	c_private_key1 := C.CString(preivateKey)
	c_public_key := C.CString(publicKey)

	if publicKey == "" {
		var cptr = C.li17_p1_context(c_private_key1, nil)

		rst := C.GoString(cptr)
		C.free(unsafe.Pointer((cptr)))
		return rst

	} else {
		var cptr = C.li17_p1_context(c_private_key1, c_public_key)

		rst := C.GoString(cptr)
		C.free(unsafe.Pointer((cptr)))
		return rst
	}

}

func GenContextP2(preivateKey, publicKey string) string {
	c_private_key1 := C.CString(preivateKey)
	c_public_key := C.CString(publicKey)

	var cptr = C.li17_p2_context(c_private_key1, c_public_key)
	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}

func SendZKProofP1(p1 string) string {
	c_p1 := C.CString(p1)
	var cptr = C.li17_p1_refresh_send_zk_proof(c_p1)

	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}

func RecvZKProofP1(p1, ZKProof2 string) string {
	c_p1 := C.CString(p1)
	c_ZKProof2 := C.CString(ZKProof2)
	var cptr = C.li17_p1_refresh_recv_zk_proof(c_p1, c_ZKProof2)

	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}

func SendZKProofP2(p2 string) string {
	c_p2 := C.CString(p2)
	var cptr = C.li17_p2_refresh_send_zk_proof(c_p2)

	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}

func RecvZKProofP2(p2, ZKProof1 string) string {
	c_p2 := C.CString(p2)
	c_ZKProof1 := C.CString(ZKProof1)
	var cptr = C.li17_p2_refresh_recv_zk_proof(c_p2, c_ZKProof1)

	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}
