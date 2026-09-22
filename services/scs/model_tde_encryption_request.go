package scs

type TdeEncryptionRequest struct {
	InstanceId *string `json:"-"`
	Action     *string `json:"action,omitempty"`
}
