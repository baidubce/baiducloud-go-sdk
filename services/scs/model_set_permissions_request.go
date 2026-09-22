package scs

type SetPermissionsRequest struct {
	InstanceId *string `json:"-"`
	UserName   *string `json:"userName,omitempty"`
	UserType   *int32  `json:"userType,omitempty"`
}
