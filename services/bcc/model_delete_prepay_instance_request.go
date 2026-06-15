package bcc

type DeletePrepayInstanceRequest struct {
	InstanceId            *string `json:"-"`
	RelatedReleaseFlag    *bool   `json:"relatedReleaseFlag,omitempty"`
	DeleteCdsSnapshotFlag *bool   `json:"deleteCdsSnapshotFlag,omitempty"`
	DeleteRelatedEnisFlag *bool   `json:"deleteRelatedEnisFlag,omitempty"`
}
