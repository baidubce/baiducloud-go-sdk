package cfw

type BindCfwRequest struct {
	CfwId        *string    `json:"-"`
	InstanceType *string    `json:"instanceType,omitempty"`
	Instances    []*CfwBind `json:"instances,omitempty"`
}
