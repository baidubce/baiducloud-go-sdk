package ax

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListSandboxesResponse struct {
	bce.BaseResponse
	Sandboxes []*ListedSandbox `json:"sandboxes,omitempty"`
}
