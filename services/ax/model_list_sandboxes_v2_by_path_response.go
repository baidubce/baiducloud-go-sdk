package ax

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListSandboxesV2ByPathResponse struct {
	bce.BaseResponse
	Sandboxes []*ListedSandbox `json:"sandboxes,omitempty"`
}
