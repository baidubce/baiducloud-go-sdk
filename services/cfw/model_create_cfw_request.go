package cfw

type CreateCfwRequest struct {
	Name        *string       `json:"name,omitempty"`
	CfwType     *int32        `json:"type,omitempty"`
	Border      *int32        `json:"border,omitempty"`
	Description *string       `json:"description,omitempty"`
	CfwRules    []*CreateRule `json:"cfwRules,omitempty"`
}
