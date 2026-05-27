package vpc

type BindEniEipRequest struct {
	EniId            *string `json:"-"`
	ClientToken      *string `json:"-"`
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`
	PublicIpAddress  *string `json:"publicIpAddress,omitempty"`
}
