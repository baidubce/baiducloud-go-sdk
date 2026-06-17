package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeUnplannedEventRecordsResponse struct {
	bce.BaseResponse
	RequestId                  *string                   `json:"requestId,omitempty"`
	IsTruncated                *bool                     `json:"isTruncated,omitempty"`
	Marker                     *string                   `json:"marker,omitempty"`
	MaxKeys                    *int32                    `json:"maxKeys,omitempty"`
	NextMarker                 *string                   `json:"nextMarker,omitempty"`
	UnplannedMaintenanceEvents []*UnplannedEventResponse `json:"unplannedMaintenanceEvents,omitempty"`
}
