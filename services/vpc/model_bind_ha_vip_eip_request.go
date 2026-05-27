package vpc

type BindHaVipEipRequest struct {
	HaVipId         *string `json:"-"`
	ClientToken     *string `json:"-"`
	PublicIpAddress *string `json:"publicIpAddress,omitempty"`
}
