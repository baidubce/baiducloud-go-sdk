package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeDownloadTaskResponse struct {
	bce.BaseResponse
	Success *bool               `json:"success,omitempty"`
	Code    *string             `json:"code,omitempty"`
	Message *string             `json:"message,omitempty"`
	Result  *DownloadTaskResult `json:"result,omitempty"`
}
