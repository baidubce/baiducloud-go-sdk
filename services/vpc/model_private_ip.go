package vpc

type PrivateIP struct {
	PublicIpAddress  *string `json:"publicIpAddress,omitempty"`
	Primary          *bool   `json:"primary,omitempty"`
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`
}
