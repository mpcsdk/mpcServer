package nrpcclient

import (
	"context"
	"errors"
	"mpcServer/api/riskctrl"
	"mpcServer/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/nats-rpc/nrpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var errDeadLine = errors.New("nats: timeout")

func (s *sNrpcClient) RpcRiskTxs(ctx context.Context, userId string, signTxData string) (*riskctrl.TxRequestRes, error) {

	rst, err := s.riskctrl.RpcTxsRequest(&riskctrl.TxRequestReq{
		UserId:     userId,
		SignTxData: signTxData,
	})

	///
	if err != nil {
		if err.Error() == errDeadLine.Error() {
			g.Log().Warning(ctx, "RpcRiskTxs TimeOut:", "userId:", userId, "signTxData:", signTxData)
			s.Flush()
			return &riskctrl.TxRequestRes{
				Ok: consts.RiskCodeError,
			}, mpccode.CodePerformRiskError()
		}
		if nerr, ok := err.(*nrpc.Error); ok {
			return rst, mpccode.CodePerformRiskError(nerr.Message)
		} else {
			return rst, mpccode.CodePerformRiskError(gtrace.GetTraceID(ctx))
		}
	}
	///
	g.Log().Notice(ctx, "RpcRiskTxs:", "rst:", rst)
	return rst, nil
}
func (s *sNrpcClient) RpcAlive(ctx context.Context) error {

	_, err := s.riskctrl.RpcAlive(&emptypb.Empty{})
	if err != nil {
		if err.Error() == errDeadLine.Error() {
			g.Log().Warning(ctx, "RpcAlive TimeOut:")
			s.Flush()
			return nil
		}
	}
	return err
}
