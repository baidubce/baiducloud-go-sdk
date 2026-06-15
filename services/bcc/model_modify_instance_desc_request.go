package bcc

type ModifyInstanceDescRequest struct {
	InstanceId *string `json:"-"`
	Desc       *string `json:"desc,omitempty"`
}
