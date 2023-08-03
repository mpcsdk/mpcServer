package li17

import (
	"testing"
)

var msg32 = "4a2d6a86fc1bd9421f78ab5eb3805f7ebf9dc8480c25a86141e4712810ea0102"
var private_key1 = "df1a627fd5ec89eaed03fd1ab246c414b8e8d57538d330e8a281137c75b88d36"
var private_key2 = "0ac7d64995c6b4daac2688c0e40d25af50887ada5b7a4cbe197ada0bdef32375"
var public_key = "045ae6d14d4934eeb004b818d687a1ea6efff0946d043dfb9338c0601a1ae0387fd00bfcefeff11961a48edc66f62ad87feed8a9ef157efa294c91466c70039bbe"

func Test_Refresh(t *testing.T) {
	p1 := GenContextP1(private_key1, public_key)
	p2 := GenContextP2(private_key2, public_key)

	proof1 := SendZKProofP1(p1)
	pkey2 := RecvZKProofP2(p2, proof1)

	proof2 := SendZKProofP2(p2)
	pkey1 := RecvZKProofP1(p1, proof2)

	context1 := GenContextP1(pkey1, "")
	context2 := GenContextP2(pkey2, "")

	////gen
	hashProof1 := KeygenSendHashProofP1(context1)
	context2 = KeygenRecvHashProofP2(context2, hashProof1)
	ZKProof2 := KeygenSendZKProofP2(context2)

	context1 = KeygenRecvZKProofP1(context1, ZKProof2)
	ZKProof1 := KeygenSendZKProofP1(context1)
	context2 = KeygenRecvZKProofP2(context2, ZKProof1)

	pk_v1 := PublicKeyP1(context1)
	pk_v2 := PublicKeyP2(context2)

	if pk_v1 != pk_v2 {
		t.Error("pkey1 != pkey2")
	}
}
