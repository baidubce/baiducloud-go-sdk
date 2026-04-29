package privatezone

import "github.com/baidubce/baiducloud-go-sdk/bce"

type SearchForDetailsOfPrivatzoneResponse struct {
	bce.BaseResponse
	ZoneId      *string `json:"zoneId,omitempty"`
	ZoneName    *string `json:"zoneName,omitempty"`
	RecordCount *int32  `json:"recordCount,omitempty"`
	CreateTime  *string `json:"createTime,omitempty"`
	UpdateTime  *string `json:"updateTime,omitempty"`
	BindVpcs    []*Vpc  `json:"bindVpcs,omitempty"`
}
