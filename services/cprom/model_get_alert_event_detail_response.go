package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetAlertEventDetailResponse struct {
	bce.BaseResponse
	EventId           *string                 `json:"eventId,omitempty"`
	MonitorInstanceId *string                 `json:"monitorInstanceId,omitempty"`
	AlertingRuleId    *string                 `json:"alertingRuleId,omitempty"`
	AlertingRuleName  *string                 `json:"alertingRuleName,omitempty"`
	NotifyRuleId      *string                 `json:"notifyRuleId,omitempty"`
	NotifyRuleName    *string                 `json:"notifyRuleName,omitempty"`
	Severity          *string                 `json:"severity,omitempty"`
	Status            *string                 `json:"status,omitempty"`
	StartTime         *int32                  `json:"startTime,omitempty"`
	EndTime           *int32                  `json:"endTime,omitempty"`
	Duration          *int32                  `json:"duration,omitempty"`
	AlarmValue        *string                 `json:"alarmValue,omitempty"`
	Expr              *string                 `json:"expr,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	AlarmTags         *map[string]interface{} `json:"alarmTags,omitempty"`
	Labels            *map[string]interface{} `json:"labels,omitempty"`
	Annotations       *map[string]interface{} `json:"annotations,omitempty"`
	ClaimedInfo       *ClaimedInfo            `json:"claimedInfo,omitempty"`
}
