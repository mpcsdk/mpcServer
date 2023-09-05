package email

import (
	"context"
	"li17server/internal/service"
	"testing"
)

func Test_SendMail(t *testing.T) {
	service.RegisterMailCode(new())
	err := service.MailCode().SendMailCode(context.Background(), "xinwei.li@mixmarvel.com", "23")
	if err != nil {
		t.Error(err)
	}
}
