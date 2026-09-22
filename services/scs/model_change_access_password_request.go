package scs

type ChangeAccessPasswordRequest struct {
	InstanceId *string `json:"-"`
	Password   *string `json:"password,omitempty"`
}
