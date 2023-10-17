package sign

import "mpcServer/internal/logic/sign/util/li17"

func (a *sSign) SignSendRequestP1(context1 string) string {
	return li17.SignSendRequestP1(context1)
}

func (a *sSign) SignRecvRequestP2(context2 string, request string) string {
	return li17.SignRecvRequestP2(context2, request)
}
func (a *sSign) SignSendPartialP2(context2, msg string) string {
	return li17.SignSendPartialP2(context2, msg)
}
func (a *sSign) SignSendPartialP1(context1, sign2, msg string) string {
	return li17.SignSendPartialP1(context1, sign2, msg)
}
