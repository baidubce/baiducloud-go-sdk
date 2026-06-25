package cprom

type EscalateParam struct {
	Rank         *int32          `json:"rank,omitempty"`
	Condition    *ClaimCondition `json:"condition,omitempty"`
	NotifyAction *NotifyAction   `json:"notifyAction,omitempty"`
}
