package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetParameterTemplateDetailsResponse struct {
	bce.BaseResponse
	Success *bool        `json:"success,omitempty"`
	Message *interface{} `json:"message,omitempty"`
	Status  *int32       `json:"status,omitempty"`
	Result  *Result      `json:"result,omitempty"`
}
