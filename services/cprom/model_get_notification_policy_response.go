package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetNotificationPolicyResponse struct {
	bce.BaseResponse
	NotifyRuleId       *string                 `json:"notifyRuleId,omitempty"`
	NotifyRuleName     *string                 `json:"notifyRuleName,omitempty"`
	StartTime          *string                 `json:"startTime,omitempty"`
	EndTime            *string                 `json:"endTime,omitempty"`
	Channel            []*string               `json:"channel,omitempty"`
	ReceiverType       *string                 `json:"receiverType,omitempty"`
	Users              []*User                 `json:"users,omitempty"`
	UserGroups         []*UserGroup            `json:"userGroups,omitempty"`
	WebhookConfigList  []*CallbackConfig       `json:"webhookConfigList,omitempty"`
	EscalateConfigList []*EscalateParam        `json:"escalateConfigList,omitempty"`
	RepeatNotifyConfig *map[string]interface{} `json:"repeatNotifyConfig,omitempty"`
	UpdateTime         *string                 `json:"updateTime,omitempty"`
}
