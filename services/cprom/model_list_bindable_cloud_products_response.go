package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListBindableCloudProductsResponse struct {
	bce.BaseResponse
	Scopes []*ScopeDetail `json:"scopes,omitempty"`
}
