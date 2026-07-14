package as

type CreateRuleV2Request struct {
	RuleName        *string      `json:"ruleName,omitempty"`
	GroupId         *string      `json:"groupId,omitempty"`
	State           *string      `json:"state,omitempty"`
	AsType          *string      `json:"type,omitempty"`
	ActionType      *string      `json:"actionType,omitempty"`
	ActionNum       *int32       `json:"actionNum,omitempty"`
	CronTime        *string      `json:"cronTime,omitempty"`
	CooldownInSec   *int32       `json:"cooldownInSec,omitempty"`
	PeriodType      *string      `json:"periodType,omitempty"`
	PeriodValue     *int32       `json:"periodValue,omitempty"`
	PeriodStartTime *string      `json:"periodStartTime,omitempty"`
	PeriodEndTime   *string      `json:"periodEndTime,omitempty"`
	AsAlarmRule     *AsAlarmRule `json:"asAlarmRule,omitempty"`
}
