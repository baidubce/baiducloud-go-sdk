package ax

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QuerySandboxesResponse struct {
	bce.BaseResponse
	Sandboxes []*QueriedSandbox `json:"sandboxes,omitempty"`
	NextToken *string           `json:"nextToken,omitempty"`
}
