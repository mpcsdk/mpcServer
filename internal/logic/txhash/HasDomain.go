package txhash

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/gogf/gf/v2/frame/g"
)

func hashDomain(data apitypes.TypedData) (string, error) {
	domain, err := data.HashStruct("EIP712Domain", data.Domain.Map())
	if err != nil {
		return "", err
	}

	return domain.String(), err
}

func HashDomain(jsonData string) (string, error) {
	var typedData apitypes.TypedData
	if err := json.Unmarshal([]byte(jsonData), &typedData); err != nil {
		g.Log().Error(context.Background(), "DigestTxHash", err)
		return "", nil
	}

	data, err := hashDomain(typedData)
	if err != nil {
		g.Log().Error(context.Background(), "DigestTxHash", err)
		return "", nil
	}

	return data, nil
}
