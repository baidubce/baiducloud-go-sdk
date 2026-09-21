package vdb

type PasswordUsingGETRequest struct {
	InstanceId *string `json:"-"`
	Username   *string `json:"-"`
	EngineType *string `json:"-"`
}
