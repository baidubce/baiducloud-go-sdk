package ccr

type TriggerJob struct {
	CreationTime *string   `json:"creationTime,omitempty"`
	EventType    *string   `json:"eventType,omitempty"`
	Id           *int32    `json:"id,omitempty"`
	Images       []*string `json:"images,omitempty"`
	NotifyType   *string   `json:"notifyType,omitempty"`
	Operator     *string   `json:"operator,omitempty"`
	PolicyId     *int32    `json:"policyId,omitempty"`
	Status       *string   `json:"status,omitempty"`
	UpdateTime   *string   `json:"updateTime,omitempty"`
}
