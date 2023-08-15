// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	ISmsCode interface {
		SendCode(sid, receiver, code string)
		Verify(sid, code string) error
	}
)

var (
	localSmsCode ISmsCode
)

func SmsCode() ISmsCode {
	if localSmsCode == nil {
		panic("implement not found for interface ISmsCode, forgot register?")
	}
	return localSmsCode
}

func RegisterSmsCode(i ISmsCode) {
	localSmsCode = i
}
