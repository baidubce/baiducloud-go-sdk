package rapidfs

type DescribeTokenRequest struct {
	InstanceId *string `json:"instanceId,omitempty"`
	TokenId    *string `json:"tokenId,omitempty"`
}
