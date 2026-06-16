package bcc

type ImportKeypairRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	PublicKey   *string `json:"publicKey,omitempty"`
}
