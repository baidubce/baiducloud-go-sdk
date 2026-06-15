package bcc

type ReleaseInstanceByPostRequest struct {
	InstanceId            *string `json:"-"`
	RelatedReleaseFlag    *bool   `json:"relatedReleaseFlag,omitempty"`
	DeleteCdsSnapshotFlag *bool   `json:"deleteCdsSnapshotFlag,omitempty"`
	DeleteRelatedEnisFlag *bool   `json:"deleteRelatedEnisFlag,omitempty"`
	BccRecycleFlag        *bool   `json:"bccRecycleFlag,omitempty"`
	CdsAttributeActive    *bool   `json:"cdsAttributeActive,omitempty"`
}
