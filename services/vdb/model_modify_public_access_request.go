package vdb

type ModifyPublicAccessRequest struct {
	InstanceId   *string `json:"-"`
	PublicAccess *bool   `json:"publicAccess,omitempty"`
}
