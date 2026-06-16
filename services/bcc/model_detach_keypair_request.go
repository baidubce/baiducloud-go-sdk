package bcc

type DetachKeypairRequest struct {
	KeypairId   *string   `json:"-"`
	InstanceIds []*string `json:"instanceIds,omitempty"`
}
