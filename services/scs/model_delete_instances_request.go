package scs

type DeleteInstancesRequest struct {
	InstanceIds []*string `json:"instanceIds,omitempty"`
}
