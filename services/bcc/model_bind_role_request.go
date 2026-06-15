package bcc

type BindRoleRequest struct {
	RoleName  *string                  `json:"roleName,omitempty"`
	Instances []*InstancePassRoleModel `json:"instances,omitempty"`
}
