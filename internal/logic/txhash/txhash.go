package txhash

import (
	"context"
	"mpcServer/internal/service"
	"os/exec"

	proto "mpcServer/api/txhash/v1"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/grpc"

	"github.com/mpcsdk/mpcCommon/mpccode"
)

type sTxHash struct {
	ctx    gctx.Ctx
	cmd    *exec.Cmd
	client proto.TransactionClient
}

func (s *sTxHash) DigestTxHash(ctx context.Context, msg string) (string, error) {
	rst, err := s.client.DigestTxHash(ctx, &proto.TxRequest{
		Message: msg,
	})
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("msg", msg),
		))
		return "", err
	}
	return rst.Message, nil
}

func (s *sTxHash) HasDomain(ctx context.Context, msg string) (string, error) {
	rst, err := s.client.HasDomain(ctx, &proto.TxRequest{
		Message: msg,
	})
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("msg", msg),
		))
		return "", err
	}
	return rst.Message, nil
}

func (s *sTxHash) TypedDataEncoderHash(ctx context.Context, msg string) (string, error) {
	rst, err := s.client.TypedDataEncoderHash(ctx, &proto.TxRequest{
		Message: msg,
	})
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("msg", msg),
		))
		return "", err
	}
	return rst.Message, nil
}
func (s *sTxHash) start() {
	// hashserver
	s.cmd = exec.Command("node", "./utility/txhash/dist/main.js")
	err := s.cmd.Start()
	if err != nil {
		panic(err)
	}
	s.connhash()
	err = s.cmd.Wait()
	//notice: need txhash service
	panic(err)
}

func (s *sTxHash) connhash() {
	// conn, err := grpcx.Client.NewGrpcClientConn("localhost:50051")
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	conn.Connect()
	g.Log().Notice(s.ctx, "connhash server:", conn.GetState().String())
	s.client = proto.NewTransactionClient(conn)
	// rst, _ := s.client.DigestTxHash(s.ctx, &proto.TxRequest{
	// 	Message: "msg",
	// })
}

func (s *sTxHash) daemon() {
	for {
		s.start()
		select {
		case <-s.ctx.Done():
			s.cmd.Process.Kill()
		default:
		}
	}
}

func new() *sTxHash {

	ctx := gctx.GetInitCtx()
	s := &sTxHash{
		ctx: ctx,
	}
	go s.daemon()
	return s
}
func init() {
	service.RegisterTxHash(new())
}
