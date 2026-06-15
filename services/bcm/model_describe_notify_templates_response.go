package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeNotifyTemplatesResponse struct {
	bce.BaseResponse
	Success                       *bool             `json:"success,omitempty"`
	Code                          *string           `json:"code,omitempty"`
	Message                       *string           `json:"message,omitempty"`
	NotifyTemplates               []*NotifyTemplate `json:"notifyTemplates,omitempty"`
	NotifyTemplatesId             *string           `json:"notifyTemplates.id,omitempty"`
	NotifyTemplatesName           *string           `json:"notifyTemplates.name,omitempty"`
	NotifyTemplatesSource         *string           `json:"notifyTemplates.source,omitempty"`
	NotifyTemplatesSilencePeriods []*SilencePeriod  `json:"notifyTemplates.silencePeriods,omitempty"`
	NotifyTemplatesReceivers      []*NotifyReceiver `json:"notifyTemplates.receivers,omitempty"`
	NotifyTemplatesCallbacks      []*Callback       `json:"notifyTemplates.callbacks,omitempty"`
}
