package cmd

// package main

// //#cgo CFLAGS: -I../libs/
// //#cgo LDFLAGS: -L${SRCDIR}/../libs -lli17
// //
// //#include "li17.h"
// //#include <stdlib.h>
// import "C"
// import (
// 	"fmt"
// 	li17 "li17server/internal/logic/li17/util/li17"
// )

// func main() {
// 	var msg32 = "4a2d6a86fc1bd9421f78ab5eb3805f7ebf9dc8480c25a86141e4712810ea0102"
// 	private_key1 := "df1a627fd5ec89eaed03fd1ab246c414b8e8d57538d330e8a281137c75b88d36"
// 	var private_key2 = "0ac7d64995c6b4daac2688c0e40d25af50887ada5b7a4cbe197ada0bdef32375"
// 	public_key := "045ae6d14d4934eeb004b818d687a1ea6efff0946d043dfb9338c0601a1ae0387fd00bfcefeff11961a48edc66f62ad87feed8a9ef157efa294c91466c70039bbe"

// 	// refresh
// 	var p1 = li17.GenContextP1(private_key1, public_key)
// 	var p2 = li17.GenContextP2(private_key2, public_key)
// 	var zk_proof1 = li17.SendZkProofP1(p1)
// 	// p2 need zk_proof1
// 	var private_key2_ = li17.RecvZkProofP2(p2, zk_proof1)

// 	var zk_proof2 = li17.SendZkProofP2(p2)
// 	// p1 need zk_proof2
// 	var private_key1_ = li17.RecvZkProofP1(p1, zk_proof2)

// 	// context
// 	var p1_context = li17.GenContextP1(private_key1_, "")
// 	var p2_context = li17.GenContextP2(private_key2_, "")

// 	// keygen
// 	var p1_hash_proof = li17.KeygenSendHashProofP1(p1_context)
// 	// p2 need p1_hash_proof
// 	p2_context = li17.KeygenRecvHashProofP2(p2_context, p1_hash_proof)

// 	var p2_zk_proof = li17.KeygenSendZkProofP2(p2_context)
// 	// p1 need p2_zk_proof
// 	p1_context = li17.KeygenRecvZkProofP1(p1_context, p2_zk_proof)

// 	var p1_zk_proof = li17.KeygenSendZkProofP1(p1_context)
// 	// p2 need p1_zk_proof
// 	p2_context = li17.KeygenRecvZkProofP2(p2_context, p1_zk_proof)

// 	///pubkey
// 	var v1_public_key = li17.PublicKeyP1(p1_context)
// 	var v2_public_key = li17.PublicKeyP2(p2_context)

// 	if v1_public_key == v2_public_key {
// 		if public_key == v1_public_key {
// 			fmt.Println("private_key1 : ", private_key1, " => ", private_key1_)
// 			fmt.Println("private_key2 : ", private_key2, " => ", private_key2_)
// 			fmt.Println("  public_key : ", public_key)
// 		} else {
// 			fmt.Println("private_key1 : ", private_key1_)
// 			fmt.Println("private_key2 : ", private_key2_)
// 			fmt.Println("  public_key : ", v1_public_key)
// 		}

// 		// signature
// 		// var request = li17.li17_p1_signature_send_signature_request(p1_context)
// 		var request = li17.SignSendRequestP1(p1_context)

// 		// p2 need request and msg
// 		// p2_context = li17.li17_p2_signature_recv_signature_request(p2_context, request)
// 		p2_context = li17.SignRecvRequestP2(p2_context, request)
// 		// var p2_signature = li17.li17_p2_signature_send_signature_partial(p2_context, c_msg32)
// 		p2_sign := li17.SignSendPartialP2(p2_context, msg32)
// 		// var signature = li17.li17_p1_signature_recv_signature_partial(p1_context, p2_sign, c_msg32)
// 		var signature = li17.SignSendPartialP1(p1_context, p2_sign, msg32)

// 		fmt.Println("         msg : ", msg32)
// 		fmt.Println("   signature : ", signature)
// 	} else {
// 		fmt.Println("pkey1 != pkey2")
// 	}

// }