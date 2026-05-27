package vpc

type UnbindHaVipEipRequest struct {
	HaVipId     *string `json:"-"`
	ClientToken *string `json:"-"`
}
