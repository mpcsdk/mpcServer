package db

import (
	"context"
	"encoding/json"
	"li17server/internal/dao"
	"li17server/internal/model"
	"li17server/internal/model/do"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sDB) RecordTxs(ctx context.Context, data *model.AnalzyTx) error {
	g.Log().Debug(ctx, "record txs", data)
	///
	addr := strings.ToLower(data.Address)
	for _, tx := range data.Txs {
		//dao写eth_tx表
		args, err := json.Marshal(tx.Args)
		if err != nil {
			g.Log().Error(ctx, err, tx)
			continue
		}
		d := &do.EthTx{
			Address:    addr,
			Target:     tx.Target,
			MethodId:   tx.MethodId,
			MethodName: tx.MethodName,
			Sig:        tx.Sig,
			Data:       tx.Data,
			Args:       string(args),
		}
		///todo: specific method
		if tx.MethodName == "transferFrom" {
			d.From = strings.ToLower(tx.Args["from"].(common.Address).Hex())
			d.To = strings.ToLower(tx.Args["to"].(common.Address).Hex())
			if val, ok := tx.Args["tokenId"]; ok {
				d.Value = val.(*big.Int).String()
			}
		} else if tx.MethodName == "transfer" {
			d.From = addr
			if to, ok := tx.Args["_to"]; ok {
				d.To = strings.ToLower(to.(common.Address).String())
			}
			if val, ok := tx.Args["_value"]; ok {
				d.Value = val.(*big.Int).String()
			}
			if wad, ok := tx.Args["wad"]; ok {
				d.Value = wad.(*big.Int).String()
			}
			if amount, ok := tx.Args["amount"]; ok {
				d.Value = amount.(*big.Int).String()
			}
			if recipient, ok := tx.Args["recipient"]; ok {
				d.To = strings.ToLower(recipient.(common.Address).String())
			}

		} else {
			g.Log().Error(ctx, "UnRecognized methhod:", tx.MethodName)
		}

		///
		_, err = dao.EthTx.Ctx(ctx).Insert(d)
		if err != nil {
			g.Log().Error(ctx, "RecordTxs :", err, tx)
		}
	}

	return nil
}
