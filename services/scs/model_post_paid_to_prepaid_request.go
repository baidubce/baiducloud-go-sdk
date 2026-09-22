package scs

type PostPaidToPrepaidRequest struct {
	Duration    *int32    `json:"duration,omitempty"`
	InstanceIds []*string `json:"instanceIds,omitempty"`
}
