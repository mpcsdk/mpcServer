package sms

import "testing"

func Test_foreign(t *testing.T) {
	foreign := newforeign()
	resp, stat, err := foreign.sendSms("8613812345678", "123456")
	if err != nil {
		t.Error(err)
	}
	if stat != "" {
		t.Log(stat)
		t.Error(err)
	}
	t.Log(resp)
}
func Test_domestic(t *testing.T) {
	domestic := newdomestic()
	resp, stat, err := domestic.sendSms("+659035559", "123456")
	if err != nil {
		t.Error(err)
	}
	if stat != "" {
		t.Log(stat)
		t.Error(err)
	}
	t.Log(resp)
}
