// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	INats interface{}
)

var (
	localNats INats
)

func Nats() INats {
	if localNats == nil {
		panic("implement not found for interface INats, forgot register?")
	}
	return localNats
}

func RegisterNats(i INats) {
	localNats = i
}
