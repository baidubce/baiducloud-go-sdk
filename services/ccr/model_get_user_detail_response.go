package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetUserDetailResponse struct {
	bce.BaseResponse
	Name *string `json:"name,omitempty"`
}
