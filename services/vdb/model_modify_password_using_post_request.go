package vdb

type ModifyPasswordUsingPOSTRequest struct {
	InstanceId *string `json:"-"`
	EngineType *string `json:"-"`
	From       *string `json:"from,omitempty"`
	Password   *string `json:"password,omitempty"`
	Username   *string `json:"username,omitempty"`
}
