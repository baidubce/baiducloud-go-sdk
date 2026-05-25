package vpc

type HighlyAvailableVirtualIpBindingEipRequest struct {
	HaVipId         *string `json:"-"`
	ClientToken     *string `json:"-"`
	PublicIpAddress *string `json:"publicIpAddress,omitempty"`
}
