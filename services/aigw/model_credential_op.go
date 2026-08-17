package aigw

type CredentialOp struct {
	Operation      *string `json:"operation,omitempty"`
	CredentialName *string `json:"credentialName,omitempty"`
	Value          *string `json:"value,omitempty"`
}
