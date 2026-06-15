package bcc

type ModifyRelatedDeletePolicyRequest struct {
	InstanceId             *string `json:"-"`
	IsEipAutoRelatedDelete *bool   `json:"isEipAutoRelatedDelete,omitempty"`
}
