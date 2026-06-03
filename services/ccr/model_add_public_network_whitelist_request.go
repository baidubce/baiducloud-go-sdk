package ccr

type AddPublicNetworkWhitelistRequest struct {
	InstanceId  *string `json:"-"`
	IpAddr      *string `json:"ipAddr,omitempty"`
	Description *string `json:"description,omitempty"`
}
