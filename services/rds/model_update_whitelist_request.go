package rds

type UpdateWhitelistRequest struct {
	InstanceId  *string   `json:"-"`
	Action      *string   `json:"-"`
	SecurityIps []*string `json:"securityIps,omitempty"`
	Tag         *string   `json:"tag,omitempty"`
}
