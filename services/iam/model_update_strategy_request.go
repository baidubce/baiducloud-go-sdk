package iam

type UpdateStrategyRequest struct {
	PolicyName  *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Document    *string `json:"document,omitempty"`
}
