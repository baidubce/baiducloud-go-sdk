package vpc

type DeleteIpGroupRequest struct {
	IpSetId     *string `json:"-"`
	ClientToken *string `json:"-"`
}
