package vdb

type ModifyPublicAccessUsingPUTRequest struct {
	InstanceId   *string `json:"-"`
	PublicAccess *bool   `json:"publicAccess,omitempty"`
}
