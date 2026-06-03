package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CheckBeforeAddCacheNodesResponse struct {
	bce.BaseResponse
	Pass     *bool      `json:"pass,omitempty"`
	ErrInfos []*ErrInfo `json:"errInfos,omitempty"`
}
