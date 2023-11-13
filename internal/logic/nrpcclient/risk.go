package nrpcclient

import (
	"context"
	"errors"
	"mpcServer/internal/consts"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "mpcServer/api/risk/nrpc/v1"
)

var errDeadLine = errors.New("nats: timeout")

func (s *sNrpcClient) RpcRiskTxs(ctx context.Context, userId string, signTxData string) (*v1.TxRiskRes, error) {

	rst, err := s.riskcli.RpcRiskTxs(&v1.TxRiskReq{
		UserId: userId,
		// Address:    ,
		SignTxData: signTxData,
		// Txs:     risktxs,
	})

	///
	if err != nil {
		if err.Error() == errDeadLine.Error() {
			g.Log().Warning(ctx, "RpcRiskTxs TimeOut:")
			s.Flush()
			return &v1.TxRiskRes{
				Ok: consts.RiskCodePass,
			}, nil
		}
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("useid", userId),
			mpccode.ErrDetail("signtx", signTxData),
		))
		return nil, err
	}
	///
	g.Log().Notice(ctx, "RpcRiskTxs:", "rst:", rst)
	return rst, nil
}
func (s *sNrpcClient) RpcAlive(ctx context.Context) error {

	_, err := s.riskcli.RpcAlive(&emptypb.Empty{})
	if err != nil {
		if err.Error() == errDeadLine.Error() {
			g.Log().Warning(ctx, "RpcAlive TimeOut:")
			s.Flush()
			return nil
		}
	}
	return err
}
