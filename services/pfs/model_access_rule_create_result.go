package pfs

type AccessRuleCreateResult struct {
	Success      *bool   `json:"success,omitempty"`
	AccessRuleId *int32  `json:"accessRuleId,omitempty"`
	ErrorMsg     *string `json:"errorMsg,omitempty"`
	Ip           *string `json:"ip,omitempty"`
	Mask         *int32  `json:"mask,omitempty"`
}
