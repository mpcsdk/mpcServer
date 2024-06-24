package txhash

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/gogf/gf/v2/frame/g"
)

func typedDataEncoderHash(data apitypes.TypedData) (string, error) {
	domain, err := data.HashStruct("EIP712Domain", data.Domain.Map())
	if err != nil {
		return "", err
	}
	dataHash, err := data.HashStruct(data.PrimaryType, data.Message)
	if err != nil {
		return "", err
	}

	prefixedData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domain), string(dataHash)))
	prefixedDataHash := crypto.Keccak256(prefixedData)
	return "0x" + common.Bytes2Hex(prefixedDataHash), nil
}

func TypedDataEncoderHash(jsonData string) (string, error) {
	var typedData apitypes.TypedData
	if err := json.Unmarshal([]byte(jsonData), &typedData); err != nil {
		g.Log().Error(context.Background(), "DigestTxHash", err)
		return "", nil
	}

	data, err := typedDataEncoderHash(typedData)
	if err != nil {
		g.Log().Error(context.Background(), "DigestTxHash", err)
		return "", nil
	}

	return data, nil
}
