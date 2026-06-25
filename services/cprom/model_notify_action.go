package cprom

type NotifyAction struct {
	Channel           []*string         `json:"channel,omitempty"`
	ReceiverType      *string           `json:"receiverType,omitempty"`
	Users             []*User           `json:"users,omitempty"`
	UserGroups        []*UserGroup      `json:"userGroups,omitempty"`
	WebhookConfigList []*CallbackConfig `json:"webhookConfigList,omitempty"`
}
