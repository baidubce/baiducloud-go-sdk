package ccr

type Whitelist struct {
	IpAddr      *string `json:"ipAddr,omitempty"`
	Description *string `json:"description,omitempty"`
}
