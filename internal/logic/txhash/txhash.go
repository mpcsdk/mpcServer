package txhash

import (
	"context"
	"fmt"
	"mpcServer/internal/config"
	"mpcServer/internal/service"
	"os/exec"
	"sync"

	proto "mpcServer/api/txhash/v1"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/grpc"

	"github.com/mpcsdk/mpcCommon/mpccode"
)

type sTxHash struct {
	ctx  gctx.Ctx
	cmds []*exec.Cmd

	////
	clilock sync.Mutex
	maxcli  uint
	poscli  uint
	txclis  []proto.TransactionClient
}

func (s *sTxHash) client() proto.TransactionClient {
	s.clilock.Lock()
	defer s.clilock.Unlock()
	///
	i := s.poscli % s.maxcli
	s.poscli++

	return s.txclis[i]
}

func (s *sTxHash) DigestTxHash(ctx context.Context, msg string) (string, error) {
	rst, err := s.client().DigestTxHash(ctx, &proto.TxRequest{
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
	rst, err := s.client().HasDomain(ctx, &proto.TxRequest{
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
	rst, err := s.client().TypedDataEncoderHash(ctx, &proto.TxRequest{
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
func (s *sTxHash) start(i uint) {
	s.maxcli = i
	for x := uint(0); x < i; x++ {
		url := fmt.Sprintf("127.0.0.1:%d", 50000+x)
		fmt.Println("start txhash server:", url)
		// hashserver
		cmd := exec.Command("node", "./utility/txhash/dist/main.js", "--url", url)
		go func(cmd *exec.Cmd) {
			err := cmd.Run()
			if err != nil {
				panic(err)
			}
		}(cmd)
		s.cmds = append(s.cmds, cmd)
		s.connhash(url)
		// go func(cmd *exec.Cmd) {
		// 	err := cmd.Wait()
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// }(cmd)
		// err = s.cmd.Wait()
		//notice: need txhash service
		// panic(err)
	}
}

func (s *sTxHash) connhash(url string) {
	// conn, err := grpcx.Client.NewGrpcClientConn("localhost:50051")
	conn, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	conn.Connect()
	g.Log().Notice(s.ctx, "connhash server:", conn.GetState().String())
	s.txclis = append(s.txclis, proto.NewTransactionClient(conn))
}

func (s *sTxHash) daemon() {
	s.start(config.Config.Server.HashCore)
	for {
		select {
		case <-s.ctx.Done():
			for _, cmd := range s.cmds {
				cmd.Process.Kill()
			}
		default:
		}
	}
}

func new() *sTxHash {

	ctx := gctx.GetInitCtx()
	s := &sTxHash{
		ctx:    ctx,
		cmds:   []*exec.Cmd{},
		txclis: []proto.TransactionClient{},
		poscli: 0,
		maxcli: 0,
	}
	go s.daemon()
	return s
}
func init() {
	service.RegisterTxHash(new())
}
