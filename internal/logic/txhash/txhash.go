package txhash

import (
	"context"
	"li17server/internal/consts"
	"li17server/internal/service"
	"os/exec"

	proto "li17server/api/txhash/v1"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/grpc"
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
		g.Log().Warning(ctx, "DigestTxHash:", err)
		return "", gerror.NewCode(consts.CodeInternalError)
	}
	return rst.Message, nil
}

func (s *sTxHash) HasDomain(ctx context.Context, msg string) (string, error) {
	rst, err := s.client.HasDomain(ctx, &proto.TxRequest{
		Message: msg,
	})
	if err != nil {
		g.Log().Warning(ctx, "HasDomain:", err)
		return "", gerror.NewCode(consts.CodeInternalError)
	}
	return rst.Message, nil
}

func (s *sTxHash) TypedDataEncoderHash(ctx context.Context, msg string) (string, error) {
	rst, err := s.client.TypedDataEncoderHash(ctx, &proto.TxRequest{
		Message: msg,
	})
	if err != nil {
		g.Log().Warning(ctx, "TypedDataEncoderHash:", err)
		return "", gerror.NewCode(consts.CodeInternalError)
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
	g.Log().Info(s.ctx, "connhash server:", conn.GetState().String())
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
