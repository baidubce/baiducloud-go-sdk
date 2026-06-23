package cprom

type ClaimAlertEventRequest struct {
	EventIds    []*string `json:"eventIds,omitempty"`
	ClaimReason *string   `json:"claimReason,omitempty"`
}
