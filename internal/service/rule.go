// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "li17server/api/rules/v1"
	signv1 "li17server/api/sign/v1"
)

type (
	IRule interface {
		Exec(txs []*signv1.SignTxData) (*v1.RiskRes, error)
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
