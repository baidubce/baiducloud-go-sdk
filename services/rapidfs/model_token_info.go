package rapidfs

type TokenInfo struct {
	TokenId                     *string `json:"tokenId,omitempty"`
	InstanceId                  *string `json:"instanceId,omitempty"`
	TokenValue                  *string `json:"tokenValue,omitempty"`
	TokenRefreshIntervalMinutes *int32  `json:"tokenRefreshIntervalMinutes,omitempty"`
	TokenExpireTime             *string `json:"tokenExpireTime,omitempty"`
}
