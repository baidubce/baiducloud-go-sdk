package eip

type RecycleEipModel struct {
	Name                *string `json:"name,omitempty"`
	Eip                 *string `json:"eip,omitempty"`
	EipId               *string `json:"eipId,omitempty"`
	Status              *string `json:"status,omitempty"`
	RouteType           *string `json:"routeType,omitempty"`
	BandwidthInMbps     *int32  `json:"bandwidthInMbps,omitempty"`
	PaymentTiming       *string `json:"paymentTiming,omitempty"`
	BillingMethod       *string `json:"billingMethod,omitempty"`
	RecycleTime         *string `json:"recycleTime,omitempty"`
	ScheduledDeleteTime *string `json:"scheduledDeleteTime,omitempty"`
}
