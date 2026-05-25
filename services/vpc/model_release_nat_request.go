package vpc

type ReleaseNatRequest struct {
	NatId       *string `json:"-"`
	ClientToken *string `json:"-"`
}
