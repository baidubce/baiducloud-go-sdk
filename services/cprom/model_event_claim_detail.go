package cprom

type EventClaimDetail struct {
	EventId *string `json:"eventId,omitempty"`
	Success *bool   `json:"success,omitempty"`
	Message *string `json:"message,omitempty"`
}
