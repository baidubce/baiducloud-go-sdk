package cfw

type CreateStatelessCfwRequest struct {
	Name          *string   `json:"name,omitempty"`
	Description   *string   `json:"description,omitempty"`
	DefaultAction *string   `json:"defaultAction,omitempty"`
	Protocol      *string   `json:"protocol,omitempty"`
	IpList        []*string `json:"ipList,omitempty"`
}
