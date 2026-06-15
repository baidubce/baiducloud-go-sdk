package bcc

type UnbindRoleRequest struct {
	RoleName  *string                  `json:"roleName,omitempty"`
	Instances []*InstancePassRoleModel `json:"instances,omitempty"`
}
