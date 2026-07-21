package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ObtainAListOfPermanentlyValidApikeysResponse struct {
	bce.BaseResponse
	Success *bool  `json:"success,omitempty"`
	Status  *int32 `json:"status,omitempty"`
	Page    *Page  `json:"page,omitempty"`
}
