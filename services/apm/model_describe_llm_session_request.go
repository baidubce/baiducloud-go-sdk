package apm

type DescribeLLMSessionRequest struct {
	SessionID     *string `json:"sessionID,omitempty"`
	BeginDatetime *string `json:"beginDatetime,omitempty"`
	EndDatetime   *string `json:"endDatetime,omitempty"`
}
