package ethtx

import (
	"context"
	"encoding/hex"
	v1 "li17server/api/rules/v1"
	"li17server/internal/service"

	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func tidyTx(tx *v1.RiskReq) *v1.RiskReq {
	tx.Data = strings.Replace(tx.Data, "0x", "", -1)
	return tx
}

type sEthTx struct {
}

func (s *sEthTx) tidy() {

}

func (s *sEthTx) AnalzyContractData(ctx context.Context, target string, data string) (map[string]interface{}, error) {
	target = strings.ToLower(target)
	data = strings.TrimPrefix(data, "0x")
	///
	contractabi, err := service.Db().GetAbi(ctx, target)
	if err != nil {
		return nil, err
	}
	//
	contract, err := abi.JSON(strings.NewReader(contractabi))
	if err != nil {
		return nil, err
	}
	//data
	dataByte, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}
	////
	method, err := contract.MethodById(dataByte[:4])
	if err != nil {
		return nil, err
	}
	//todo: check method
	//
	args := make(map[string]interface{})
	err = method.Inputs.UnpackIntoMap(args, dataByte[4:])
	if err != nil {
		return nil, err
	}
	///todo: args
	return args, nil
}

func new() *sEthTx {
	return &sEthTx{}
}
func init() {
	service.RegisterEthTx(new())
}
