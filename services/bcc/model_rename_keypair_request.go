package bcc

type RenameKeypairRequest struct {
	KeypairId *string `json:"-"`
	Name      *string `json:"name,omitempty"`
}
