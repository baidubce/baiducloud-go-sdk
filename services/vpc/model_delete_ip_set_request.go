package vpc

type DeleteIpSetRequest struct {
	IpGroupId   *string `json:"-"`
	ClientToken *string `json:"-"`
}
