package main

//#cgo CFLAGS: -I../libs/
//#cgo LDFLAGS: -L${SRCDIR}/../libs -lli17
//#include "li17.h"
//#include <stdlib.h>
import "C"

import (
	"fmt"
	"mpcServer/internal/logic/sign/util/li17"
	"sync"
)

func runtest(i int, private_key1, private_key2, msg32, public_key string, wg *sync.WaitGroup) {
	defer wg.Done()
	// refresh
	var p1 = li17.GenContextP1(private_key1, public_key)
	var p2 = li17.GenContextP2(private_key2, public_key)
	var zk_proof1 = li17.SendZKProofP1(p1)
	// p2 need zk_proof1
	var private_key2_ = li17.RecvZKProofP2(p2, zk_proof1)

	var zk_proof2 = li17.SendZKProofP2(p2)
	// p1 need zk_proof2
	var private_key1_ = li17.RecvZKProofP1(p1, zk_proof2)

	// context
	var p1_context = li17.GenContextP1(private_key1_, "")
	var context_p2 = li17.GenContextP2(private_key2_, "")

	// keygen
	var p1_hash_proof = li17.KeygenSendHashProofP1(p1_context)
	// p2 need p1_hash_proof
	context_p2 = li17.KeygenRecvHashProofP2(context_p2, p1_hash_proof)

	var p2_zk_proof = li17.KeygenSendZKProofP2(context_p2)
	// p1 need p2_zk_proof
	p1_context = li17.KeygenRecvZKProofP1(p1_context, p2_zk_proof)

	var p1_zk_proof = li17.KeygenSendZKProofP1(p1_context)
	// p2 need p1_zk_proof
	context_p2 = li17.KeygenRecvZKProofP2(context_p2, p1_zk_proof)

	///pubkey
	var v1_public_key = li17.PublicKeyP1(p1_context)
	var v2_public_key = li17.PublicKeyP2(context_p2)

	if v1_public_key == v2_public_key {
		// if public_key == v1_public_key {
		// 	fmt.Println("private_key1 : ", private_key1, " => ", private_key1_)
		// 	fmt.Println("private_key2 : ", private_key2, " => ", private_key2_)
		// 	fmt.Println("  public_key : ", public_key)
		// } else {
		// 	fmt.Println("private_key1 : ", private_key1_)
		// 	fmt.Println("private_key2 : ", private_key2_)
		// 	fmt.Println("  public_key : ", v1_public_key)
		// }

		// signature
		// var request = li17.li17_p1_signature_send_signature_request(p1_context)
		var request = li17.SignSendRequestP1(p1_context)

		// p2 need request and msg
		// context_p2 = li17.li17_p2_signature_recv_signature_request(context_p2, request)
		context_p2 = li17.SignRecvRequestP2(context_p2, request)
		// var p2_signature = li17.li17_p2_signature_send_signature_partial(context_p2, c_msg32)
		p2_sign := li17.SignSendPartialP2(context_p2, msg32)
		// var signature = li17.li17_p1_signature_recv_signature_partial(p1_context, p2_sign, c_msg32)
		var signature = li17.SignSendPartialP1(p1_context, p2_sign, msg32)

		if signature == "" {
			fmt.Println("           p1: ", p1)
			fmt.Println("           p2: ", p2)
			fmt.Println("    zk_proof1: ", zk_proof1)
			fmt.Println("private_key2_: ", private_key2_)
			fmt.Println("    zk_proof2: ", zk_proof2)
			fmt.Println("private_key1_: ", private_key1_)
			fmt.Println("   p1_context: ", p1_context)
			fmt.Println("   context_p2: ", context_p2)
			fmt.Println("p1_hash_proof: ", p1_hash_proof)
			fmt.Println("   context_p2: ", context_p2)
			fmt.Println("  p2_zk_proof: ", p2_zk_proof)
			fmt.Println("   p1_context: ", p1_context)
			fmt.Println("  p1_zk_proof: ", p1_zk_proof)
			fmt.Println("   context_p2: ", context_p2)
			fmt.Println("v1_public_key: ", v1_public_key)
			fmt.Println("v2_public_key: ", v2_public_key)
			fmt.Println("===============Sign==============")

			fmt.Println("   request: ", request)
			fmt.Println("context_p2: ", context_p2)
			fmt.Println("   p2_sign: ", p2_sign)
			fmt.Println(" signature: ", signature)
			fmt.Println("      msg : ", msg32)
			panic(i)
		}
		fmt.Println(">", i)
	} else {
		fmt.Println("pkey1 != pkey2")
	}

}
func main() {
	var msg32 = "4a2d6a86fc1bd9421f78ab5eb3805f7ebf9dc8480c25a86141e4712810ea0102"
	private_key1 := "df1a627fd5ec89eaed03fd1ab246c414b8e8d57538d330e8a281137c75b88d36"
	var private_key2 = "0ac7d64995c6b4daac2688c0e40d25af50887ada5b7a4cbe197ada0bdef32375"
	public_key := "045ae6d14d4934eeb004b818d687a1ea6efff0946d043dfb9338c0601a1ae0387fd00bfcefeff11961a48edc66f62ad87feed8a9ef157efa294c91466c70039bbe"

	wg := new(sync.WaitGroup)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go runtest(i, private_key1, private_key2, msg32, public_key, wg)
	}
	wg.Wait()
}
