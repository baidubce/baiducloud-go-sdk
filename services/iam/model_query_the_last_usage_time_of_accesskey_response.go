package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryTheLastUsageTimeOfAccesskeyResponse struct {
	bce.BaseResponse
	AccessKeyId  *string `json:"accessKeyId,omitempty"`
	LastUsedTime *string `json:"lastUsedTime,omitempty"`
}
