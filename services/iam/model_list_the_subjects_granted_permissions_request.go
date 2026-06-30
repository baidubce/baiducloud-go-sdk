package iam

type ListTheSubjectsGrantedPermissionsRequest struct {
	PolicyId  *string `json:"-"`
	GrantType *string `json:"-"`
}
