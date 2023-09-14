package db

import (
	"context"
	"encoding/json"
	"li17server/internal/dao"
	"li17server/internal/model"
	"li17server/internal/model/entity"
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
		d := &entity.EthTx{
			Address:    addr,
			Target:     tx.Target,
			MethodId:   tx.MethodId,
			MethodName: tx.MethodName,
			Sig:        tx.Sig,
			Data:       tx.Data,
			Args:       string(args),
		}
		///todo: specific methdo
		if tx.MethodName == "safeTransferFrom" {
			d.From = tx.Args["_from"].(common.Address).Hex()
			d.To = tx.Args["_to"].(common.Address).Hex()
			d.Value = tx.Args["_tokenIndex"].(string)
		} else if tx.MethodName == "transfer" {
			d.From = addr
			if to, ok := tx.Args["_to"]; ok {
				d.To = to.(common.Address).String()
			}
			if val, ok := tx.Args["_value"]; ok {
				d.Value = val.(*big.Int).String()
			}
			if wad, ok := tx.Args["wad"]; ok {
				d.Value = wad.(*big.Int).String()
			}
		} else {
			g.Log().Error(ctx, "UnRecognized methhod:", tx.MethodName)
		}

		///
		d.From = strings.ToLower(d.From)
		d.To = strings.ToLower(d.To)

		///
		_, err = dao.EthTx.Ctx(ctx).Insert(d)
		if err != nil {
			g.Log().Error(ctx, "RecordTxs :", err, tx)
		}
	}

	return nil
}
