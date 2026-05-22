package cfw

type UpdateCfwRequest struct {
	CfwId       *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
