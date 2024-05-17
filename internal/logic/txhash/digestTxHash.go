package txhash

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Tx struct {
	ChainId *big.Int       `json:"chainId"`
	Address common.Address `json:"address"`
	Number  *big.Int       `json:"number"`
	Txs     []UnipassTx    `json:"txs"`
	TxHash  string         `json:"txHash,omitempty"`
}
type internalTx struct {
	ChainId int64          `json:"chainId"`
	Address common.Address `json:"address"`
	Number  int64          `json:"number"`
	Txs     []UnipassTx    `json:"txs"`
	TxHash  string         `json:"txHash,omitempty"`
}

func (s *Tx) UnmarshalJSON(data []byte) error {
	a := &internalTx{}
	err := json.Unmarshal(data, a)
	if err != nil {
		return err
	}
	s.ChainId = big.NewInt(int64(a.ChainId))
	s.Address = a.Address
	s.Number = big.NewInt(int64(a.Number))
	s.Txs = a.Txs
	s.TxHash = a.TxHash
	return nil
}

type unipassBigInt struct {
	Type string `json:"type"`
	Hex  string `json:"hex"`
}

func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

// isHex validates whether each byte is valid hexadecimal string.
func isHex(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	for _, c := range []byte(str) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}

type UnipassTx struct {
	CallType      uint8          `json:"callType"`
	RevertOnError bool           `json:"revertOnError"`
	GasLimit      *big.Int       `json:"gasLimit"`
	Target        common.Address `json:"target"`
	Value         *big.Int       `json:"value"`
	Data          []byte         `json:"data"`
}
type internalUnipassTx struct {
	CallType      uint8          `json:"callType"`
	RevertOnError bool           `json:"revertOnError"`
	GasLimit      unipassBigInt  `json:"gasLimit"`
	Target        common.Address `json:"target"`
	Value         unipassBigInt  `json:"value"`
	Data          string         `json:"data"`
}

func (s *UnipassTx) UnmarshalJSON(data []byte) error {
	a := &internalUnipassTx{}
	err := json.Unmarshal(data, a)
	if err != nil {
		return err
	}
	//////
	// b := hexutil.Big{}
	// hex := "0x01"
	// b := common.FromHex(hex)
	// b = common.TrimLeftZeroes(b)
	// err = b.UnmarshalText([]byte("0x01"))
	// if err != nil {
	// 	return err
	//

	s.CallType = a.CallType
	s.RevertOnError = a.RevertOnError
	b := common.FromHex(a.GasLimit.Hex)
	b = common.TrimLeftZeroes(b)
	s.GasLimit = big.NewInt(0).SetBytes(b)
	s.Target = a.Target
	b = common.FromHex(a.Value.Hex)
	b = common.TrimLeftZeroes(b)
	s.Value = big.NewInt(0).SetBytes(b)
	b = common.FromHex(a.Data)
	b = common.TrimLeftZeroes(b)
	s.Data = b
	return nil
}

// //
// //
func DigestTxHash(signData string) string {
	tx := Tx{}
	err := json.Unmarshal([]byte(signData), &tx)
	if err != nil {
		fmt.Println(err)
	}
	///
	Uint256, _ := abi.NewType("uint256", "", nil)
	TupleArr, _ := abi.NewType("tuple[]", "UnipassTx", []abi.ArgumentMarshaling{
		{Name: "CallType", Type: "uint8"},
		{Name: "RevertOnError", Type: "bool"},
		{Name: "Target", Type: "address"},
		{Name: "GasLimit", Type: "uint256"},
		{Name: "Value", Type: "uint256"},
		{Name: "Data", Type: "bytes"},
	})
	m := abi.NewMethod("foo", "foo", abi.Constructor, "", false, false, []abi.Argument{
		{"", Uint256, false},
		{"", TupleArr, false},
	}, nil)

	b, err := m.Inputs.Pack(&tx.Number, tx.Txs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(common.Bytes2Hex(b))
	b = crypto.Keccak256(b)
	////
	///
	arr := []byte{}
	arr = append(arr, []byte("\x19\x01")...)
	arr = append(arr, common.LeftPadBytes(tx.ChainId.Bytes(), 32)...)
	arr = append(arr, tx.Address.Bytes()...)
	arr = append(arr, b...)
	b = crypto.Keccak256(arr)
	return "0x" + common.Bytes2Hex(b)
}
