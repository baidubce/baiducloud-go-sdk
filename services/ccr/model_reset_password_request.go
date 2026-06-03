package ccr

type ResetPasswordRequest struct {
	InstanceId *string `json:"-"`
	Password   *string `json:"password,omitempty"`
}
