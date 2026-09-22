package scs

type FeatureSwitches struct {
	GroupModify            *bool `json:"groupModify,omitempty"`
	CrossAzNearest         *bool `json:"crossAzNearest,omitempty"`
	ProxyUpgradeSupport    *bool `json:"proxyUpgradeSupport,omitempty"`
	RecoverInOriginSupport *bool `json:"recoverInOriginSupport,omitempty"`
	RecoverInNewSupport    *bool `json:"recoverInNewSupport,omitempty"`
	SupportSentinelSwitch  *bool `json:"supportSentinelSwitch,omitempty"`
	WhitelistGroupSupport  *bool `json:"whitelistGroupSupport,omitempty"`
	BandwidthModify        *bool `json:"bandwidthModify,omitempty"`
	ShadowBackupSupport    *bool `json:"shadowBackupSupport,omitempty"`
}
