package rds

type RenewalExampleRequest struct {
	Duration    *int32    `json:"duration,omitempty"`
	InstanceIds []*string `json:"instanceIds,omitempty"`
}
