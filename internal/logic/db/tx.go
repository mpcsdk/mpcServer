package db

import (
	"context"
	"encoding/json"
	"li17server/internal/dao"
	"li17server/internal/model"
	"li17server/internal/model/entity"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

func (s *sDB) RecordTxs(ctx context.Context, data *model.AnalzyTx) error {

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
		_, err = dao.EthTx.Ctx(ctx).Insert(d)
		if err != nil {
			g.Log().Error(ctx, "RecordTxs :", err, tx)
		}
	}

	return nil
}
