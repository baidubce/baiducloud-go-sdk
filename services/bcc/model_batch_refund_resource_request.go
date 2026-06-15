package bcc

type BatchRefundResourceRequest struct {
	InstanceIds           []*string `json:"instanceIds,omitempty"`
	RelatedReleaseFlag    *bool     `json:"relatedReleaseFlag,omitempty"`
	DeleteCdsSnapshotFlag *bool     `json:"deleteCdsSnapshotFlag,omitempty"`
	DeleteRelatedEnisFlag *bool     `json:"deleteRelatedEnisFlag,omitempty"`
}
