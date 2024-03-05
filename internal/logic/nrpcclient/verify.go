package nrpcclient

import (
	"context"
	"mpcServer/api/riskctrl"

	"github.com/gogf/gf/v2/frame/g"
)

func (s *sNrpcClient) RpcSendMailCode(ctx context.Context, userId, serial string) error {

	rst, err := s.riskctrl.RpcSendMailCode(&riskctrl.SendMailCodeReq{
		RiskSerial: serial,
		UserId:     userId,
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

func (s *sNrpcClient) RpcSendSmsCode(ctx context.Context, userId, serial string) error {

	rst, err := s.riskctrl.RpcSendPhoneCode(&riskctrl.SendPhoneCodeReq{
		RiskSerial: serial,
		UserId:     userId,
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

func (s *sNrpcClient) RpcVerifyCode(ctx context.Context, userId, serial, phoneCode, mailCode string) error {

	rst, err := s.riskctrl.RpcVerifyCode(&riskctrl.VerifyCodeReq{
		UserId:     userId,
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
