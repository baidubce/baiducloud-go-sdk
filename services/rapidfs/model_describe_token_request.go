package rapidfs

type DescribeTokenRequest struct {
	Action     *string `json:"-"`
	InstanceId *string `json:"instanceId,omitempty"`
	TokenId    *string `json:"tokenId,omitempty"`
}
