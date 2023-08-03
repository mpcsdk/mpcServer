package sign

import "li17server/internal/logic/sign/util/li17"

func (a *sGenerator) SignSendRequestP1(context1 string) string {
	return li17.SignSendRequestP1(context1)
}

func (a *sGenerator) SignRecvRequestP2(context2 string, request string) string {
	return li17.SignRecvRequestP2(context2, request)
}
func (a *sGenerator) SignSendPartialP2(context2, msg string) string {
	return li17.SignSendPartialP2(context2, msg)
}
func (a *sGenerator) SignSendPartialP1(context1, sign2, msg string) string {
	return li17.SignSendPartialP1(context1, sign2, msg)
}
