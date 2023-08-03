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
import (
	"unsafe"
)

func KeygenSendHashProofP1(context1 string) string {

	c_context := C.CString(context1)
	var proof = C.li17_p1_keygen_send_hash_proof(c_context)
	return C.GoString(proof)
}

func KeygenRecvHashProofP2(context2, proof1 string) string {
	c_context2 := C.CString(context2)
	c_proof1 := C.CString(proof1)
	p2_context := C.li17_p2_keygen_recv_hash_proof(c_context2, c_proof1)
	return C.GoString(p2_context)
}

func KeygenSendZkProofP1(context1 string) string {
	c_context1 := C.CString(context1)
	var p1_zk_proof = C.li17_p1_keygen_send_zk_proof(c_context1)

	return C.GoString(p1_zk_proof)
}

func KeygenRecvZkProofP1(context1, proof2 string) string {

	s_context1 := (*Li17String)(unsafe.Pointer(&context1)).Str
	s_proof2 := (*Li17String)(unsafe.Pointer(&proof2)).Str

	// C.Print((*C.char)(s_context1))
	// C.Print((*C.char)(s_proof2))

	p1_context := C.li17_p1_keygen_recv_zk_proof((*C.char)(s_context1), (*C.char)(s_proof2))
	return C.GoString(p1_context)

}

func KeygenSendZkProofP2(context1 string) string {
	c_context := C.CString(context1)
	var p1_zk_proof = C.li17_p2_keygen_send_zk_proof(c_context)
	return C.GoString(p1_zk_proof)
}

func KeygenRecvZkProofP2(context2, proof1 string) string {
	c_context := C.CString(context2)
	proof := C.CString(proof1)
	p2_context := C.li17_p2_keygen_recv_zk_proof(c_context, proof)
	return C.GoString(p2_context)
}

func PublicKeyP1(context1 string) string {
	c_context := C.CString(context1)
	var v1_public_key = C.li17_p1_public_key(c_context)
	return C.GoString(v1_public_key)
}
func PublicKeyP2(context2 string) string {
	c_context := C.CString(context2)
	var v2_public_key = C.li17_p2_public_key(c_context)
	return C.GoString(v2_public_key)
}
