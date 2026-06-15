package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeCacheRuleJobsResponse struct {
	bce.BaseResponse
	CacheJobInfos []*CacheJobInfo `json:"cacheJobInfos,omitempty"`
	Marker        *string         `json:"marker,omitempty"`
	IsTruncated   *bool           `json:"isTruncated,omitempty"`
	NextMarker    *string         `json:"nextMarker,omitempty"`
	MaxKeys       *int32          `json:"maxKeys,omitempty"`
}
