package db

import (
	"context"
	"fmt"
	"li17server/internal/model"
)

func (s *sDB) RecordTxs(ctx context.Context, data *model.AnalzyTx) error {

	fmt.Println(data)
	return nil
}
