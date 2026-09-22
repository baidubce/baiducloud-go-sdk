package scs

type RenewInstanceRequest struct {
	InstanceIds []*string `json:"instanceIds,omitempty"`
	Duration    *int32    `json:"duration,omitempty"`
}
