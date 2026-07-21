package iam

type ACL struct {
	Id                *string     `json:"id,omitempty"`
	Version           *string     `json:"version,omitempty"`
	AccessControlList []*ACLEntry `json:"accessControlList,omitempty"`
}
