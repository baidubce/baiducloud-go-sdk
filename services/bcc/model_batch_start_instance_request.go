package bcc

type BatchStartInstanceRequest struct {
	InstanceIds []*string `json:"instanceIds,omitempty"`
}
