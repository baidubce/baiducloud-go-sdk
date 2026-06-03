package rapidfs

type ModifyTokenRequest struct {
	ClientToken                 *string `json:"-"`
	InstanceId                  *string `json:"instanceId,omitempty"`
	TokenId                     *string `json:"tokenId,omitempty"`
	TokenRefreshIntervalMinutes *int32  `json:"tokenRefreshIntervalMinutes,omitempty"`
}
