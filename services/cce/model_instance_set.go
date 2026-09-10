package cce

type InstanceSet struct {
	InstanceSpec *interface{} `json:"instanceSpec,omitempty"`
	Count        *int32       `json:"count,omitempty"`
}
