package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryAndModifyHistoryResponse struct {
	bce.BaseResponse
	Name        *string `json:"name,omitempty"`
	BeforeValue *string `json:"beforeValue,omitempty"`
	AfterValue  *string `json:"afterValue,omitempty"`
	Status      *string `json:"status,omitempty"`
	UpdateTime  *string `json:"updateTime,omitempty"`
}
