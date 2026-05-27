package vpc

type UnbindEniEipRequest struct {
	EniId           *string `json:"-"`
	ClientToken     *string `json:"-"`
	PublicIpAddress *string `json:"publicIpAddress,omitempty"`
}
