package scs

type ProxyVersionUpgradeOrRestartRequest struct {
	InstanceId  *string   `json:"-"`
	ProxyList   []*string `json:"proxyList,omitempty"`
	UpgradeType *string   `json:"upgradeType,omitempty"`
	IsDefer     *bool     `json:"isDefer,omitempty"`
}
