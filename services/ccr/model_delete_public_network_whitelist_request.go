package ccr

type DeletePublicNetworkWhitelistRequest struct {
	InstanceId *string   `json:"-"`
	Items      []*string `json:"items,omitempty"`
}
