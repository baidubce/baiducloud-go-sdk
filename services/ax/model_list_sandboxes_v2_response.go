package ax

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListSandboxesV2Response struct {
	bce.BaseResponse
	Sandboxes []*ListedSandbox `json:"sandboxes,omitempty"`
}
