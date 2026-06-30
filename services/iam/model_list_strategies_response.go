package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListStrategiesResponse struct {
	bce.BaseResponse
	Policies []*PolicyModel `json:"policies,omitempty"`
}
