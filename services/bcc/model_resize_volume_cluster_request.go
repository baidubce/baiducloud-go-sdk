package bcc

type ResizeVolumeClusterRequest struct {
	ClusterId          *string `json:"-"`
	NewClusterSizeInGB *int32  `json:"newClusterSizeInGB,omitempty"`
}
