package bcc

type ReleaseMultipleInstanceByPostRequest struct {
	InstanceIds           []*string `json:"instanceIds,omitempty"`
	RelatedReleaseFlag    *bool     `json:"relatedReleaseFlag,omitempty"`
	DeleteRelatedCdsFlag  *bool     `json:"deleteRelatedCdsFlag,omitempty"`
	DeleteCdsSnapshotFlag *bool     `json:"deleteCdsSnapshotFlag,omitempty"`
	DeleteRelatedEnisFlag *bool     `json:"deleteRelatedEnisFlag,omitempty"`
	BccRecycleFlag        *bool     `json:"bccRecycleFlag,omitempty"`
}
