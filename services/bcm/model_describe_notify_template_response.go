package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeNotifyTemplateResponse struct {
	bce.BaseResponse
	Success        *bool             `json:"success,omitempty"`
	Code           *string           `json:"code,omitempty"`
	Message        *string           `json:"message,omitempty"`
	Id             *string           `json:"id,omitempty"`
	Name           *string           `json:"name,omitempty"`
	Source         *string           `json:"source,omitempty"`
	SilencePeriods []*SilencePeriod  `json:"silencePeriods,omitempty"`
	Receivers      []*NotifyReceiver `json:"receivers,omitempty"`
	Callbacks      []*Callback       `json:"callbacks,omitempty"`
}
