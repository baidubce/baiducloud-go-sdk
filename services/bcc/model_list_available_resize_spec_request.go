package bcc

type ListAvailableResizeSpecRequest struct {
	Spec           *string   `json:"spec,omitempty"`
	SpecId         *string   `json:"specId,omitempty"`
	Zone           *string   `json:"zone,omitempty"`
	InstanceIdList []*string `json:"instanceIdList,omitempty"`
}
