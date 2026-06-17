package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeAlarmsResponse struct {
	bce.BaseResponse
	Alarms     []*Alarm `json:"alarms,omitempty"`
	PageNo     *int32   `json:"pageNo,omitempty"`
	PageSize   *int32   `json:"pageSize,omitempty"`
	TotalCount *int32   `json:"totalCount,omitempty"`
}
