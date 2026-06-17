package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribePlannedEventsResponse struct {
	bce.BaseResponse
	RequestId                *string                 `json:"requestId,omitempty"`
	IsTruncated              *bool                   `json:"isTruncated,omitempty"`
	Marker                   *string                 `json:"marker,omitempty"`
	MaxKeys                  *int32                  `json:"maxKeys,omitempty"`
	NextMarker               *string                 `json:"nextMarker,omitempty"`
	PlannedMaintenanceEvents []*PlannedEventResponse `json:"plannedMaintenanceEvents,omitempty"`
}
