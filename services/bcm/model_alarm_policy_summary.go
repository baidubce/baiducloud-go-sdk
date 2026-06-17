package bcm

type AlarmPolicySummary struct {
	Id      *string `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Content *string `json:"content,omitempty"`
	Level   *string `json:"level,omitempty"`
}
