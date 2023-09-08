package cmd

import (
	"context"
	"li17server/internal/controller/sign"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
func ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}
	var (
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if code == gcode.CodeNil {
		if err != nil {
			code = gcode.CodeInternalError
		} else {
			code = gcode.CodeOK
		}
	}
	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    code.Code(),
		Message: code.Message(),
		Data:    res,
	})
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(MiddlewareCORS)
				group.Middleware(ResponseHandler)
				group.Bind(
					sign.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
