package sms

import (
	"context"
	"errors"
	"li17server/internal/service"
	"strconv"

	"github.com/dchest/captcha"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
)

type sSmsCode struct {
	demestic *huawei
	foreign  *huawei
	pool     *grpool.Pool
}

func newforeign() *huawei {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	return &huawei{
		APIAddress:        cfg.MustGet(ctx, "sms.foreign.APIAddress").String(),
		ApplicationKey:    cfg.MustGet(ctx, "sms.foreign.ApplicationKey").String(),
		ApplicationSecret: cfg.MustGet(ctx, "sms.foreign.ApplicationSecret").String(),
		Sender:            cfg.MustGet(ctx, "sms.foreign.Sender").String(),
		TemplateID:        cfg.MustGet(ctx, "sms.foreign.TemplateID").String(),
		Signature:         cfg.MustGet(ctx, "sms.foreign.Signature").String(),
	}
}
func newdomestic() *huawei {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	return &huawei{
		APIAddress:        cfg.MustGet(ctx, "sms.domestic.APIAddress").String(),
		ApplicationKey:    cfg.MustGet(ctx, "sms.domestic.ApplicationKey").String(),
		ApplicationSecret: cfg.MustGet(ctx, "sms.domestic.ApplicationSecret").String(),
		Sender:            cfg.MustGet(ctx, "sms.domestic.Sender").String(),
		TemplateID:        cfg.MustGet(ctx, "sms.domestic.TemplateID").String(),
		Signature:         cfg.MustGet(ctx, "sms.domestic.Signature").String(),
	}
}

func (s *sSmsCode) sendCode(ctx context.Context, sid, receiver, code string) error {

	resp, status, err := s.foreign.sendSms(receiver, code)
	//todo: send smscode
	service.Generator().RecordSid(ctx, sid, "smscode", "123456")
	return nil
	///
	if err != nil {
		service.Generator().RecordSid(ctx, sid, "smserr", err.Error())
		return err
	}
	if status != "" {
		service.Generator().RecordSid(ctx, sid, "smserr", status)
		return errors.New(status)
	}
	g.Log().Debug(ctx, "sendCode:", receiver, code, resp)
	///
	service.Generator().RecordSid(ctx, sid, "smscode", code)
	return nil
}

func (s *sSmsCode) SendCode(ctx context.Context, sid, receiver, code string) error {

	//todo: get phone by  sid
	d := captcha.RandomDigits(6)
	code = ""
	for _, b := range d {
		code += strconv.Itoa(int(b))
	}
	////

	return s.sendCode(ctx, sid, receiver, code)
}
func (s *sSmsCode) Verify(ctx context.Context, sid, code string) error {
	c, err := service.Cache().Get(ctx, sid+"smscode")
	if err == nil {
		if c.String() != code {
			return errors.New("verfiy fauild")
		}
	}
	//todo: faild stat
	stat, err := service.Cache().Get(ctx, sid+"sms")
	if err != nil {
		return err
	}
	if stat.String() == "err" {
		estr, err := service.Cache().Get(ctx, sid+"smserr")
		if err != nil {
			return err
		}
		return errors.New(estr.String())
	}

	status, err := service.Cache().Get(ctx, sid+"smsstatus")
	if err != nil {
		return err
	}
	///
	g.Log().Debug(ctx, "Verify:", status)
	return nil

}
func new() *sSmsCode {

	return &sSmsCode{
		pool:     grpool.New(10),
		foreign:  newforeign(),
		demestic: newdomestic(),
	}
}

func init() {
	service.RegisterSmsCode(new())
}
