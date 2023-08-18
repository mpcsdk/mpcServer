// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	IRule interface {
		Exec() error
	}
)

var (
	localRule IRule
)

func Rule() IRule {
	if localRule == nil {
		panic("implement not found for interface IRule, forgot register?")
	}
	return localRule
}

func RegisterRule(i IRule) {
	localRule = i
}
