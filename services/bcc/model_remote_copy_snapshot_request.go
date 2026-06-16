package bcc

type RemoteCopySnapshotRequest struct {
	SnapshotId      *string              `json:"-"`
	Uuid            *string              `json:"uuid,omitempty"`
	DestRegionInfos []*RemoteCopyRequest `json:"destRegionInfos,omitempty"`
}
