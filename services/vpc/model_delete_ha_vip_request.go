package vpc

type DeleteHaVipRequest struct {
	HaVipId     *string `json:"-"`
	ClientToken *string `json:"-"`
}
