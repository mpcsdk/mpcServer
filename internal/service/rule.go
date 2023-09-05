// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "li17server/api/rules/v1"
	"li17server/internal/model"
)

type (
	IRule interface {
		Exec(from string, txs []*model.SignTxData) (*v1.RiskRes, error)
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
