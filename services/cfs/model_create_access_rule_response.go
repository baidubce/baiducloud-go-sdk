package cfs

type CreateAccessRuleResponse struct {
	Sucess *bool   `json:"sucess,omitempty"`
	ArIdx  *int32  `json:"ar_idx,omitempty"`
	ErrMsg *string `json:"err_msg,omitempty"`
}
