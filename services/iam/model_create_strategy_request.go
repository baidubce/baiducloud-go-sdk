package iam

type CreateStrategyRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Document    *string `json:"document,omitempty"`
}
