package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetTagBuildHistoryResponse struct {
	bce.BaseResponse
	Items []*BuildHistory `json:"items,omitempty"`
}
