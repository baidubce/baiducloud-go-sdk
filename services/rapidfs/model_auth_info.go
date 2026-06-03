package rapidfs

type AuthInfo struct {
	Cidr        *string `json:"cidr,omitempty"`
	Mode        *string `json:"mode,omitempty"`
	Description *string `json:"description,omitempty"`
}
