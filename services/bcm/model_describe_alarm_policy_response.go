package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeAlarmPolicyResponse struct {
	bce.BaseResponse
	Success         *bool           `json:"success,omitempty"`
	Code            *string         `json:"code,omitempty"`
	Message         *string         `json:"message,omitempty"`
	Id              *string         `json:"id,omitempty"`
	Name            *string         `json:"name,omitempty"`
	Scope           *string         `json:"scope,omitempty"`
	ResourceType    *string         `json:"resourceType,omitempty"`
	SubResourceType *string         `json:"subResourceType,omitempty"`
	Target          *AlarmTarget    `json:"target,omitempty"`
	Rules           []*AlarmRule    `json:"rules,omitempty"`
	Actions         []*PolicyAction `json:"actions,omitempty"`
	CreatedTime     *string         `json:"createdTime,omitempty"`
	UpdatedTime     *string         `json:"updatedTime,omitempty"`
}
