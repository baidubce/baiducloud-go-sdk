package scs

type UpdateTlsEncryptionRequest struct {
	InstanceId *string `json:"-"`
	Action     *string `json:"action,omitempty"`
}
