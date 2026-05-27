package csn

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryCsnDetailResponse struct {
	bce.BaseResponse
	Name        *string     `json:"name,omitempty"`
	Description *string     `json:"description,omitempty"`
	CsnId       *string     `json:"csnId,omitempty"`
	Status      *string     `json:"status,omitempty"`
	InstanceNum *int32      `json:"instanceNum,omitempty"`
	CsnBpNum    *int32      `json:"csnBpNum,omitempty"`
	CreatedTime *string     `json:"createdTime,omitempty"`
	Tags        []*TagModel `json:"tags,omitempty"`
}
