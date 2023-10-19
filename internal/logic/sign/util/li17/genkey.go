package li17

//#cgo CFLAGS: -I./libs/
//#cgo LDFLAGS: -L${SRCDIR}/libs -lli17
//
//#include "li17.h"
//#include <stdlib.h>
//#include <string.h>
// #include <stdio.h>
// void Print(const char *p) {
// 	printf("printf:%s\n", p);
// }
import "C"
import "unsafe"

func KeygenSendHashProofP1(context1 string) string {

	c_context := C.CString(context1)
	var cptr = C.li17_p1_keygen_send_hash_proof(c_context)
	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}

func KeygenRecvHashProofP2(context2, proof1 string) string {
	c_context2 := C.CString(context2)
	c_proof1 := C.CString(proof1)
	cptr := C.li17_p2_keygen_recv_hash_proof(c_context2, c_proof1)
	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}

func KeygenSendZKProofP1(context1 string) string {
	c_context1 := C.CString(context1)
	var cptr = C.li17_p1_keygen_send_zk_proof(c_context1)

	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}

func KeygenRecvZKProofP1(context1, proof2 string) string {

	s_context1 := C.CString(context1)
	s_proof2 := C.CString(proof2)

	cptr := C.li17_p1_keygen_recv_zk_proof(s_context1, s_proof2)
	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}

func KeygenSendZKProofP2(context1 string) string {
	c_context := C.CString(context1)
	var cptr = C.li17_p2_keygen_send_zk_proof(c_context)
	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}

func KeygenRecvZKProofP2(context2, proof1 string) string {
	c_context := C.CString(context2)
	proof := C.CString(proof1)
	cptr := C.li17_p2_keygen_recv_zk_proof(c_context, proof)
	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}

func PublicKeyP1(context1 string) string {
	c_context := C.CString(context1)
	var cptr = C.li17_p1_public_key(c_context)
	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}
func PublicKeyP2(context2 string) string {
	c_context := C.CString(context2)
	var cptr = C.li17_p2_public_key(c_context)
	rst := C.GoString(cptr)
	C.free(unsafe.Pointer((cptr)))
	return rst
}
