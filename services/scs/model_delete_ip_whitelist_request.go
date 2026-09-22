package scs

type DeleteIpWhitelistRequest struct {
	InstanceId  *string   `json:"-"`
	SecurityIps []*string `json:"securityIps,omitempty"`
}
