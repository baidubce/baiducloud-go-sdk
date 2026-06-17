package bcm

type Alarm struct {
	Id           *string             `json:"id,omitempty"`
	SeriesId     *string             `json:"seriesId,omitempty"`
	State        *string             `json:"state,omitempty"`
	InitState    *string             `json:"initState,omitempty"`
	CloseReason  *string             `json:"closeReason,omitempty"`
	StartTime    *string             `json:"startTime,omitempty"`
	EndTime      *string             `json:"endTime,omitempty"`
	BcmType      *string             `json:"type,omitempty"`
	Resource     *AlarmResource      `json:"resource,omitempty"`
	Policy       *AlarmPolicySummary `json:"policy,omitempty"`
	Actions      []*AlarmAction      `json:"actions,omitempty"`
	AlertMetrics []*AlertMetric      `json:"alertMetrics,omitempty"`
}
