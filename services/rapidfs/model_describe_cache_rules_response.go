package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeCacheRulesResponse struct {
	bce.BaseResponse
	CacheRuleInfos []*CacheRuleInfo `json:"cacheRuleInfos,omitempty"`
	Marker         *string          `json:"marker,omitempty"`
	IsTruncated    *bool            `json:"isTruncated,omitempty"`
	NextMarker     *string          `json:"nextMarker,omitempty"`
	MaxKeys        *int32           `json:"maxKeys,omitempty"`
}
