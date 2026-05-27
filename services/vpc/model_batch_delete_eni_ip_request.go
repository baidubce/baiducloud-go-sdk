package vpc

type BatchDeleteEniIpRequest struct {
	EniId              *string   `json:"-"`
	ClientToken        *string   `json:"-"`
	PrivateIpAddresses []*string `json:"privateIpAddresses,omitempty"`
}
