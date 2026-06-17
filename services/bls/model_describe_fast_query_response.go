package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeFastQueryResponse struct {
	bce.BaseResponse
	CreationDateTime *string `json:"creationDateTime,omitempty"`
	LastModifiedTime *string `json:"lastModifiedTime,omitempty"`
	FastQueryName    *string `json:"fastQueryName,omitempty"`
	Description      *string `json:"description,omitempty"`
	Query            *string `json:"query,omitempty"`
	Project          *string `json:"project,omitempty"`
	LogStoreName     *string `json:"logStoreName,omitempty"`
	LogStreamName    *string `json:"logStreamName,omitempty"`
}
