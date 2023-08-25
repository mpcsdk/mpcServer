// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package rules

import (
	"context"
	
	"li17server/api/rules/v1"
)

type IRulesV1 interface {
	Risk(ctx context.Context, req *v1.RiskReq) (res *v1.RiskRes, err error)
}


