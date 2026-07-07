package bci

type BatchDeleteInstancesRequest struct {
	InstanceIds        []*string `json:"instanceIds,omitempty"`
	RelatedReleaseFlag *bool     `json:"relatedReleaseFlag,omitempty"`
}
