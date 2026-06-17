package bcc

type DescribeUnplannedEventsRequest struct {
	Action                   *string   `json:"-"`
	ServerEventStatus        *string   `json:"serverEventStatus,omitempty"`
	ServerEventIds           []*string `json:"serverEventIds,omitempty"`
	InstanceIds              []*string `json:"instanceIds,omitempty"`
	ProductCategory          *string   `json:"productCategory,omitempty"`
	ServerEventType          *string   `json:"serverEventType,omitempty"`
	ServerEventLogTimeFilter *string   `json:"serverEventLogTimeFilter,omitempty"`
	PeriodStartTime          *string   `json:"periodStartTime,omitempty"`
	PeriodEndTime            *string   `json:"periodEndTime,omitempty"`
	MaxKeys                  *int32    `json:"maxKeys,omitempty"`
	Marker                   *string   `json:"marker,omitempty"`
}
