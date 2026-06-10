package bcc

type PurchaseReservedVolumeRequest struct {
	VolumeId   *string  `json:"-"`
	Billing    *Billing `json:"billing,omitempty"`
	InstanceId *string  `json:"instanceId,omitempty"`
}
