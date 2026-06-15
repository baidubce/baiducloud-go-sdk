package bcc

type PurchaseReservedInstanceRequest struct {
	InstanceId       *string            `json:"-"`
	RelatedRenewFlag *string            `json:"-"`
	Billing          *Billing           `json:"billing,omitempty"`
	CdsCustomPeriod  []*CdsCustomPeriod `json:"cdsCustomPeriod,omitempty"`
}
