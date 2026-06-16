package cfw

type UpdateStatelessCfwRequest struct {
	CfwId       *string   `json:"-"`
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Protocol    *string   `json:"protocol,omitempty"`
	IpList      []*string `json:"ipList,omitempty"`
}
