package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	v1 "li17server/api/sign/v1"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/v2/os/gcmd"
)

func data2json(data interface{}) string {
	bf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(bf)
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)

	return bf.String()
}

var sid = "448212937646149"

var mainCmd = &gcmd.Command{
	Name:        "main",
	Brief:       "start http server",
	Description: "this is the command entry for starting your process",
}

var sendHashProof = &gcmd.Command{
	Name: "hashproof",
	Arguments: []gcmd.Argument{
		{
			Name:  "hashproof",
			Short: "p",
		},
	},
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		proof := parser.GetOpt("p").String()
		fmt.Println("hashproof:", proof)
		nproof := strings.ReplaceAll(proof, "\\", "")
		// sid := flag.Match("i", "sessionid").String()

		s := v1.SendHashProofReq{
			SessionId: sid,
			HashProof: nproof,
		}
		nb := data2json(s)
		res, err := g.Client().Post(
			"http://123.60.148.168:8000/SendHashProof",
			nb,
		)
		fmt.Println("send hashproof err:", err)

		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Response.Body)
		fmt.Println("send hashproof res:\n", buf.String())
		fmt.Println("============================================================================")
		return
	},
}

var sendZKProofP1 = &gcmd.Command{
	Name: "sendZKProofP1",
	Arguments: []gcmd.Argument{
		{
			Name:  "zkproofp1",
			Short: "z",
		},
	},
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		zkproofp1 := parser.GetOpt("z").String()
		fmt.Println(zkproofp1)
		nproof := strings.ReplaceAll(zkproofp1, "\\", "")

		s := v1.SendZKProofP1Req{
			SessionId: sid,
			ZKProofP1: nproof,
		}
		nb := data2json(s)
		res, err := g.Client().Post(
			"http://123.60.148.168:8000/SendZKProofP1",
			nb,
		)
		fmt.Println("send zkproofp1:", err)

		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Response.Body)
		fmt.Println("send zkproofp1 res:\n", buf.String())
		fmt.Println("============================================================================")
		return
	},
}

var signMsg = &gcmd.Command{
	Name: "signMsg",
	Arguments: []gcmd.Argument{
		{
			Name:  "msg",
			Short: "msg",
			IsArg: true,
		}, {
			Name:  "request",
			Short: "r",
			IsArg: true,
		},
	},
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		msg := parser.GetOpt("m").String()
		request := parser.GetOpt("r").String()

		s := v1.SignMsgReq{
			SessionId: sid,
			Msg:       msg,
			Request:   request,
		}
		nb := data2json(s)
		res, err := g.Client().Post(
			"http://123.60.148.168:8000/SignMsg",
			nb,
		)
		fmt.Println("send msg:", err)

		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Response.Body)
		fmt.Println("send msg res:\n", buf.String())
		fmt.Println("============================================================================")
		return
	},
}

func main() {
	err := mainCmd.AddCommand(sendHashProof, sendZKProofP1, signMsg)
	if err != nil {
		panic(err)
	}
	mainCmd.Run(gctx.New())
}
