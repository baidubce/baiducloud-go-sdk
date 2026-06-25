package cprom

type WebhookDetail struct {
	HookName       *string              `json:"hookName,omitempty"`
	HookMethod     *string              `json:"hookMethod,omitempty"`
	HookUrl        *string              `json:"hookUrl,omitempty"`
	Headers        *map[string]string   `json:"headers,omitempty"`
	Params         *map[string]string   `json:"params,omitempty"`
	MentionedUsers *MentionedUserConfig `json:"mentionedUsers,omitempty"`
}
