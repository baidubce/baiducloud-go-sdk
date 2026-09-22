package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetBackUpUsageResponse struct {
	bce.BaseResponse
	LogicalLogBackupBillingSizeBytes *int64  `json:"logicalLogBackupBillingSizeBytes,omitempty"`
	SnapshotDataBackupSizeBytes      *int64  `json:"snapshotDataBackupSizeBytes,omitempty"`
	PhysicalDataBackupSizeBytes      *int64  `json:"physicalDataBackupSizeBytes,omitempty"`
	LogicalLogBackupSizeBytes        *int64  `json:"logicalLogBackupSizeBytes,omitempty"`
	LogicalDataBackupSizeBytes       *int64  `json:"logicalDataBackupSizeBytes,omitempty"`
	PhysicalLogBackupSizeBytes       *int64  `json:"physicalLogBackupSizeBytes,omitempty"`
	DataType                         *string `json:"dataType,omitempty"`
}
