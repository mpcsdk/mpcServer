package ethtx

import (
	"context"
	"encoding/hex"
	v1 "li17server/api/rules/v1"
	"li17server/internal/model"
	"li17server/internal/service"

	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/gogf/gf/v2/os/gcache"
)

func tidyTx(tx *v1.RiskReq) *v1.RiskReq {
	tx.Data = strings.Replace(tx.Data, "0x", "", -1)
	return tx
}

type sEthTx struct {
	abicache *gcache.Cache
}

func (s *sEthTx) analzyTx(ctx context.Context, tx *model.SignTxData) (*model.AnalzyTxData, error) {

	target := strings.ToLower(tx.Target)
	data := strings.TrimPrefix(tx.Data, "0x")
	///
	contractabi := ""
	if a, err := s.abicache.Get(ctx, target); !a.IsEmpty() {
		contractabi = a.String()
	} else {
		contractabi, err = service.DB().GetAbi(ctx, target)
		if err != nil {
			return nil, err
		}
		s.abicache.Set(ctx, target, contractabi, 0)
	}
	///
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
	///
	atx := &model.AnalzyTxData{
		Target:     target,
		MethodId:   hex.EncodeToString(method.ID),
		MethodName: method.RawName,
		Sig:        method.Sig,
		Data:       data,
		Args:       args,
	}
	return atx, nil
}

func (s *sEthTx) tidy(signtxs *model.SignTx) {
	signtxs.Address = strings.ToLower(signtxs.Address)
}
func (s *sEthTx) AnalzyTxs(ctx context.Context, signtxs *model.SignTx) (*model.AnalzyTx, error) {
	// s.tidy(signtxs)
	atx := &model.AnalzyTx{}
	atx.Address = strings.ToLower(signtxs.Address)
	///
	for _, tx := range signtxs.Txs {
		adata, err := s.analzyTx(ctx, tx)
		if err != nil {
			return nil, err
		}
		atx.Txs = append(atx.Txs, adata)
	}
	return atx, nil
}

func new() *sEthTx {
	return &sEthTx{
		abicache: gcache.New(),
	}
}
func init() {
	service.RegisterEthTx(new())
}
