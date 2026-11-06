package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryBackupListResponse struct {
	bce.BaseResponse
	Marker      *string          `json:"marker,omitempty"`
	IsTruncated *bool            `json:"isTruncated,omitempty"`
	NextMarker  *string          `json:"nextMarker,omitempty"`
	Backups     []*OpenApiBackup `json:"backups,omitempty"`
}
