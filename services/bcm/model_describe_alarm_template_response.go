package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeAlarmTemplateResponse struct {
	bce.BaseResponse
	Success         *bool        `json:"success,omitempty"`
	Code            *string      `json:"code,omitempty"`
	Message         *string      `json:"message,omitempty"`
	Id              *string      `json:"id,omitempty"`
	Scope           *string      `json:"scope,omitempty"`
	ResourceType    *string      `json:"resourceType,omitempty"`
	SubResourceType *string      `json:"subResourceType,omitempty"`
	Name            *string      `json:"name,omitempty"`
	Comment         *string      `json:"comment,omitempty"`
	Rules           []*AlarmRule `json:"rules,omitempty"`
}
