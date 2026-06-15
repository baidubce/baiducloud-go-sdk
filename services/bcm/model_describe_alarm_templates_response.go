package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeAlarmTemplatesResponse struct {
	bce.BaseResponse
	Success                       *bool            `json:"success,omitempty"`
	Code                          *string          `json:"code,omitempty"`
	Message                       *string          `json:"message,omitempty"`
	PageNo                        *int32           `json:"pageNo,omitempty"`
	PageSize                      *int32           `json:"pageSize,omitempty"`
	TotalCount                    *int32           `json:"totalCount,omitempty"`
	AlarmTemplates                []*AlarmTemplate `json:"alarmTemplates,omitempty"`
	AlarmTemplatesId              *string          `json:"alarmTemplates.id,omitempty"`
	AlarmTemplatesScope           *string          `json:"alarmTemplates.scope,omitempty"`
	AlarmTemplatesResourceType    *string          `json:"alarmTemplates.resourceType,omitempty"`
	AlarmTemplatesSubResourceType *string          `json:"alarmTemplates.subResourceType,omitempty"`
	AlarmTemplatesName            *string          `json:"alarmTemplates.name,omitempty"`
	AlarmTemplatesComment         *string          `json:"alarmTemplates.comment,omitempty"`
	AlarmTemplatesRules           []*AlarmRule     `json:"alarmTemplates.rules,omitempty"`
}
