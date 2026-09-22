package scs

type AddIpWhitelistRequest struct {
	InstanceId  *string   `json:"-"`
	SecurityIps []*string `json:"securityIps,omitempty"`
}
