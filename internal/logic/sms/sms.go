package sms

import (
	"context"
	"errors"
	"fmt"
	"li17server/internal/service"
	"strconv"

	"github.com/dchest/captcha"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
)

type sSmsCode struct {
	ctx  gctx.Ctx
	sms  *huawei
	pool *grpool.Pool
}

func (s *sSmsCode) sendCode(sid, receiver, code string) {

	resp, status, err := s.sms.sendSms(receiver, code)
	//todo:
	fmt.Println(resp)
	fmt.Println(status)
	///
	if err != nil {
		service.Generator().RecordSid(s.ctx, sid, "smserr", err.Error())
	}
	if status != "" {
		service.Generator().RecordSid(s.ctx, sid, "smserr", status)
	}
	///
	service.Generator().RecordSid(s.ctx, sid, "smscode", code)
}

func (s *sSmsCode) SendCode(sid, receiver, code string) {

	//todo: get phone by  sid
	d := captcha.RandomDigits(6)
	code = ""
	for _, b := range d {
		code += strconv.Itoa(int(b))
	}
	////

	s.pool.Add(s.ctx, func(ctx context.Context) {
		s.sendCode(sid, receiver, code)
	})
}
func (s *sSmsCode) Verify(sid, code string) error {
	c, err := service.Cache().Get(s.ctx, sid+"smscode")
	if err == nil {
		if c.String() != code {
			return errors.New("verfiy fauild")
		}
	}
	//todo: faild stat
	stat, err := service.Cache().Get(s.ctx, sid+"sms")
	if err != nil {
		return err
	}
	if stat.String() == "err" {
		estr, err := service.Cache().Get(s.ctx, sid+"smserr")
		if err != nil {
			return err
		}
		return errors.New(estr.String())
	}

	status, err := service.Cache().Get(s.ctx, sid+"smsstatus")
	if err != nil {
		return err
	}
	///
	fmt.Println(status.String())
	return nil
}
func new() *sSmsCode {

	return &sSmsCode{
		ctx:  gctx.GetInitCtx(),
		pool: grpool.New(10),
		sms:  newhuawei(),
	}
}

func init() {
	service.RegisterSmsCode(new())
}
