package bcc

type UpdateKeypairDescriptionRequest struct {
	KeypairId   *string `json:"-"`
	Description *string `json:"description,omitempty"`
}
