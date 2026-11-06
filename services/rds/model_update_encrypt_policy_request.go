package rds

type UpdateEncryptPolicyRequest struct {
	EncryptStrategy []*EncryptPolicyRequest `json:"encryptStrategy,omitempty"`
}
