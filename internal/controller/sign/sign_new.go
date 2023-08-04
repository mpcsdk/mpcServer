package sign

import (
	"li17server/api/sign"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type ControllerV1 struct{}

var tmp_privkey2 string = "0ac7d64995c6b4daac2688c0e40d25af50887ada5b7a4cbe197ada0bdef32375"

func NewV1() sign.ISignV1 {
	return &ControllerV1{}
}

func randPrivateKey() string {
	pk, _ := crypto.GenerateKey()
	pkBytes := crypto.FromECDSA(pk)
	pkstr := hexutil.Encode(pkBytes)[2:]
	return pkstr
}
