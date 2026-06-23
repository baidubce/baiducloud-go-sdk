package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListRelatedCloudProductsResponse struct {
	bce.BaseResponse
	Scopes []*string `json:"scopes,omitempty"`
}
