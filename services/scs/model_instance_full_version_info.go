package scs

type InstanceFullVersionInfo struct {
	ProxyFullVersion             *string `json:"proxyFullVersion,omitempty"`
	RedisOrPegaFullVerison       *string `json:"redisOrPegaFullVerison,omitempty"`
	ProxyLatestFullVersion       *string `json:"proxyLatestFullVersion,omitempty"`
	RedisOrPegaLatestFullVersion *string `json:"redisOrPegaLatestFullVersion,omitempty"`
	IsProxyCanUpgrade            *bool   `json:"isProxyCanUpgrade,omitempty"`
	IsRedisOrPegaCanUpgrade      *bool   `json:"isRedisOrPegaCanUpgrade,omitempty"`
	IsPegaCanRestart             *bool   `json:"isPegaCanRestart,omitempty"`
}
