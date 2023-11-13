package nrpcclient

import (
	"context"
	v1 "mpcServer/api/tfa/nrpc/v1"

	"github.com/gogf/gf/v2/frame/g"
)

func (s *sNrpcClient) RpcSendMailCode(ctx context.Context, token, serial string) error {

	rst, err := s.tfacli.RpcSendMailCode(&v1.MailCodekReq{
		RiskSerial: serial,
		Token:      token,
	})

	if err != nil {
		if err.Error() == errDeadLine.Error() {
			g.Log().Warning(ctx, "RpcSendSmsCode TimeOut:")
			s.Flush()
			return nil
		}
		return err
	}
	g.Log().Notice(ctx, "RpcSendMailCode:", "rst:", rst)
	return nil
}

func (s *sNrpcClient) RpcSendSmsCode(ctx context.Context, token, serial string) error {

	rst, err := s.tfacli.RpcSendSmsCode(&v1.SmsCodeReq{
		RiskSerial: serial,
		Token:      token,
	})

	if err != nil {
		if err.Error() == errDeadLine.Error() {
			g.Log().Warning(ctx, "RpcSendSmsCode TimeOut:")
			s.Flush()
			return nil
		}
		return err
	}
	g.Log().Notice(ctx, "RpcSendMailCode:", "rst:", rst)
	return nil
}

func (s *sNrpcClient) RpcVerifyCode(ctx context.Context, token, serial, phoneCode, mailCode string) error {

	rst, err := s.tfacli.RpcSendVerifyCode(&v1.VerifyCodekReq{
		Token:      token,
		RiskSerial: serial,
		PhoneCode:  phoneCode,
		MailCode:   mailCode,
	})

	///
	if err != nil {
		if err.Error() == errDeadLine.Error() {
			g.Log().Warning(ctx, "RpcSendSmsCode TimeOut:")
			s.Flush()
			return nil
		}
		return err
	}
	///
	g.Log().Notice(ctx, "RpcVerifyCode:", "rst:", rst)
	return nil
}
