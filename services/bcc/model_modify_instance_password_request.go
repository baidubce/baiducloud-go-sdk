package bcc

type ModifyInstancePasswordRequest struct {
	InstanceId *string `json:"-"`
	AdminPass  *string `json:"adminPass,omitempty"`
}
