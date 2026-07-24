package bcc

type RemoteCopySnapshotRequest struct {
	SnapshotId      *string              `json:"-"`
	DestRegionInfos []*RemoteCopyRequest `json:"destRegionInfos,omitempty"`
}
