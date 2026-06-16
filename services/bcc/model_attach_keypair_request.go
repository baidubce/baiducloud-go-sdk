package bcc

type AttachKeypairRequest struct {
	KeypairId   *string   `json:"-"`
	InstanceIds []*string `json:"instanceIds,omitempty"`
}
