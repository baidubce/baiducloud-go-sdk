package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetBackUpUrlResponse struct {
	bce.BaseResponse
	Url           *string `json:"url,omitempty"`
	UrlExpiration *int32  `json:"urlExpiration,omitempty"`
}
