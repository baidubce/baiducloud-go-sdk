package cprom

type CallbackConfig struct {
	WebhookType *string          `json:"webhookType,omitempty"`
	WebhookList []*WebhookDetail `json:"webhookList,omitempty"`
}
