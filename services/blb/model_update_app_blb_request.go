package blb

type UpdateAppBlbRequest struct {
	BlbId        *string `json:"-"`
	ClientToken  *string `json:"-"`
	Name         *string `json:"name,omitempty"`
	Desc         *string `json:"desc,omitempty"`
	AllowDelete  *bool   `json:"allowDelete,omitempty"`
	AllocateIpv6 *bool   `json:"allocateIpv6,omitempty"`
}
