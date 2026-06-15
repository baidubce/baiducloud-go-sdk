package bcm

type AlarmRule struct {
	Conditions           []*AlarmCondition `json:"conditions,omitempty"`
	PendingCount         *int32            `json:"pendingCount,omitempty"`
	CheckIntervalSeconds *int32            `json:"checkIntervalSeconds,omitempty"`
	Content              *string           `json:"content,omitempty"`
}
