package vdb

type BindEipUsingPOSTRequest struct {
	InstanceId *string `json:"-"`
	EngineType *string `json:"-"`
	Eip        *string `json:"eip,omitempty"`
}
