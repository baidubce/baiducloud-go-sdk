package bcc

type PurchaseReservedVolumeClusterRequest struct {
	ClusterId *string  `json:"-"`
	Billing   *Billing `json:"billing,omitempty"`
}
