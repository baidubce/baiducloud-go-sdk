package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QuerySummaryOfMainAccountResponse struct {
	bce.BaseResponse
	AccountId *string           `json:"accountId,omitempty"`
	LimitInfo *AccountLimitInfo `json:"limitInfo,omitempty"`
	CountInfo *AccountCountInfo `json:"countInfo,omitempty"`
}
