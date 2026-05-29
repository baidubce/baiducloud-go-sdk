package cfs

type ClientInfoModel struct {
	Zone     *string `json:"zone,omitempty"`
	VpcId    *string `json:"vpcId,omitempty"`
	MountId  *string `json:"mountId,omitempty"`
	ClientIp *string `json:"clientIp,omitempty"`
}
