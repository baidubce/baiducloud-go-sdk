package pfs

type FilesetModel struct {
	FilesetName      *string `json:"filesetName,omitempty"`
	FilesetId        *string `json:"filesetId,omitempty"`
	FilesetPath      *string `json:"filesetPath,omitempty"`
	Status           *int32  `json:"status,omitempty"`
	BlockQuota       *int32  `json:"blockQuota,omitempty"`
	BlockUsage       *int64  `json:"blockUsage,omitempty"`
	FilesQuota       *int64  `json:"filesQuota,omitempty"`
	FilesUsage       *int64  `json:"filesUsage,omitempty"`
	Allocatedinode   *int64  `json:"allocatedinode,omitempty"`
	Ctime            *string `json:"ctime,omitempty"`
	QpsLimit         *int64  `json:"qpsLimit,omitempty"`
	BandwidthLimitMb *int64  `json:"bandwidthLimitMb,omitempty"`
	ParentPath       *bool   `json:"parentPath,omitempty"`
}
