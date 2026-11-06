package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryParameterListResponse struct {
	bce.BaseResponse
	Etag       *string              `json:"etag,omitempty"`
	Parameters []*OpenApiConfigItem `json:"parameters,omitempty"`
}
