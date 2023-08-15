package generator

import (
	"li17server/internal/service"
	"strconv"

	"github.com/dchest/captcha"
)

func (s *sGenerator) SendSms(sid string) error {

	//todo: get phone by  sid
	d := captcha.RandomDigits(6)
	code := ""
	for _, b := range d {
		code += strconv.Itoa(int(b))
	}
	////
	service.SmsCode().SendCode(sid, "+8613818637750", code)
	return nil
}
