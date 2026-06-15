package bcc

type ListRecycleInstanceRequest struct {
	Marker        *string `json:"marker,omitempty"`
	MaxKeys       *int32  `json:"maxKeys,omitempty"`
	InstanceId    *string `json:"instanceId,omitempty"`
	Name          *string `json:"name,omitempty"`
	PaymentTiming *string `json:"paymentTiming,omitempty"`
	RecycleBegin  *string `json:"recycleBegin,omitempty"`
	RecycleEnd    *string `json:"recycleEnd,omitempty"`
}
