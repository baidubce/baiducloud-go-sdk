package ax

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BatchReleaseSandboxesResponse struct {
	bce.BaseResponse
	Results []*ReleaseResult `json:"results,omitempty"`
}
