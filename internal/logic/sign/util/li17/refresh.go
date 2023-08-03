package li17

//#cgo CFLAGS: -I./libs/
//#cgo LDFLAGS: -L${SRCDIR}/libs -lli17
//
//#include "li17.h"
//#include <stdlib.h>
import "C"
import (
	"strings"
	"unsafe"
)

func CString2G(s *C.char) string {
	gstr := *(*string)(unsafe.Pointer(&s))
	len := strings.IndexByte(gstr, byte(0))
	ustr := *(*Li17String)(unsafe.Pointer(&s))
	ustr.Len = len
	gstr = *(*string)(unsafe.Pointer(&ustr))
	return gstr
}

func GenContextP1(preivateKey, publicKey string) string {
	c_private_key1 := C.CString(preivateKey)
	c_public_key := C.CString(publicKey)

	if publicKey == "" {
		var p1 = C.li17_p1_context(c_private_key1, nil)

		return C.GoString(p1)

	} else {
		var p1 = C.li17_p1_context(c_private_key1, c_public_key)

		return C.GoString(p1)
	}

}

func GenContextP2(preivateKey, publicKey string) string {
	c_private_key1 := C.CString(preivateKey)
	c_public_key := C.CString(publicKey)

	var p1 = C.li17_p2_context(c_private_key1, c_public_key)
	return (C.GoString(p1))
}

func SendZkProofP1(p1 string) string {
	c_p1 := C.CString(p1)
	var zk_proof1 = C.li17_p1_refresh_send_zk_proof(c_p1)

	return C.GoString(zk_proof1)
}

func RecvZkProofP1(p1, zkproof2 string) string {
	c_p1 := C.CString(p1)
	c_zkproof2 := C.CString(zkproof2)
	var priv_key1 = C.li17_p1_refresh_recv_zk_proof(c_p1, c_zkproof2)

	return C.GoString(priv_key1)
}

func SendZkProofP2(p2 string) string {
	c_p2 := C.CString(p2)
	var zk_proof2 = C.li17_p2_refresh_send_zk_proof(c_p2)

	return C.GoString(zk_proof2)
}

func RecvZkProofP2(p2, zkproof1 string) string {
	c_p2 := C.CString(p2)
	c_zkproof1 := C.CString(zkproof1)
	var priv_key2 = C.li17_p2_refresh_recv_zk_proof(c_p2, c_zkproof1)

	return C.GoString(priv_key2)
}
