package blb

type CreateLbdcRequest struct {
	ClientToken      *string               `json:"-"`
	Name             *string               `json:"name,omitempty"`
	BlbType          *string               `json:"type,omitempty"`
	CcuCount         *int32                `json:"ccuCount,omitempty"`
	Desc             *string               `json:"desc,omitempty"`
	Billing          *BillingForCreate     `json:"billing,omitempty"`
	RenewReservation *ReservationForCreate `json:"renewReservation,omitempty"`
	Tags             []*TagModel           `json:"tags,omitempty"`
}
