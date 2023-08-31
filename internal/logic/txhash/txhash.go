package txhash

import (
	"context"
	"fmt"
	"li17server/internal/service"
	"os/exec"

	proto "li17server/api/txhash/v1"

	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/grpc"
)

type sTxHash struct {
	ctx    gctx.Ctx
	cmd    *exec.Cmd
	client proto.TransactionClient
}

func (s *sTxHash) DigestTxHash(ctx context.Context, msg string) string {
	rst, _ := s.client.DigestTxHash(ctx, &proto.TxRequest{
		Message: msg,
	})
	return rst.Message
}

func (s *sTxHash) start() {
	// hashserver
	s.cmd = exec.Command("node", "./txhash/dist/main.js")
	err := s.cmd.Start()
	if err != nil {
		panic(err)
	}
	s.connhash()
	err = s.cmd.Wait()
	//todo: exit
	panic(err)
}

func (s *sTxHash) connhash() {
	// conn, err := grpcx.Client.NewGrpcClientConn("localhost:50051")
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	conn.Connect()
	fmt.Println(conn.GetState().String())
	s.client = proto.NewTransactionClient(conn)
	// rst, _ := s.client.DigestTxHash(s.ctx, &proto.TxRequest{
	// 	Message: "msg",
	// })
	// fmt.Println(rst)
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
