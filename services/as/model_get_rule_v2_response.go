package as

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetRuleV2Response struct {
	bce.BaseResponse
	RuleId            *string      `json:"ruleId,omitempty"`
	RuleName          *string      `json:"ruleName,omitempty"`
	GroupId           *string      `json:"groupId,omitempty"`
	AccountId         *string      `json:"accountId,omitempty"`
	State             *string      `json:"state,omitempty"`
	AsType            *string      `json:"type,omitempty"`
	CronTime          *string      `json:"cronTime,omitempty"`
	ActionType        *string      `json:"actionType,omitempty"`
	ActionNum         *int32       `json:"actionNum,omitempty"`
	CooldownInSec     *int32       `json:"cooldownInSec,omitempty"`
	CreateTime        *string      `json:"createTime,omitempty"`
	LastExecutionTime *string      `json:"lastExecutionTime,omitempty"`
	LastScheduleTime  *string      `json:"lastScheduleTime,omitempty"`
	PeriodStartTime   *string      `json:"periodStartTime,omitempty"`
	PeriodEndTime     *string      `json:"periodEndTime,omitempty"`
	PeriodType        *string      `json:"periodType,omitempty"`
	PeriodValue       *int32       `json:"periodValue,omitempty"`
	AsAlarmRule       *AsAlarmRule `json:"asAlarmRule,omitempty"`
}
