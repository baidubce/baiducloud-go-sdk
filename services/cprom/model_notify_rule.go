package cprom

type NotifyRule struct {
	NotifyRuleId       *string             `json:"notifyRuleId,omitempty"`
	NotifyRuleName     *string             `json:"notifyRuleName,omitempty"`
	StartTime          *string             `json:"startTime,omitempty"`
	EndTime            *string             `json:"endTime,omitempty"`
	Channel            []*string           `json:"channel,omitempty"`
	ReceiverType       *string             `json:"receiverType,omitempty"`
	Users              []*User             `json:"users,omitempty"`
	UserGroups         []*UserGroup        `json:"userGroups,omitempty"`
	WebhookConfigList  []*CallbackConfig   `json:"webhookConfigList,omitempty"`
	EscalateConfigList []*EscalateParam    `json:"escalateConfigList,omitempty"`
	RepeatNotifyConfig *RepeatNotifyConfig `json:"repeatNotifyConfig,omitempty"`
	CreateTime         *string             `json:"createTime,omitempty"`
	UpdateTime         *string             `json:"updateTime,omitempty"`
}
