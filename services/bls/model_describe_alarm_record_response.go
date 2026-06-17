package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeAlarmRecordResponse struct {
	bce.BaseResponse
	Success *bool   `json:"success,omitempty"`
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Alarm   *Alarm  `json:"alarm,omitempty"`
}
