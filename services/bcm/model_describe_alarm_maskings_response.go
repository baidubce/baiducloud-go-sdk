package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeAlarmMaskingsResponse struct {
	bce.BaseResponse
	Maskings  []*AlarmMaskingModel `json:"maskings,omitempty"`
	PageNo    *int32               `json:"pageNo,omitempty"`
	PageSize  *int32               `json:"pageSize,omitempty"`
	TotalSize *int32               `json:"totalSize,omitempty"`
}
