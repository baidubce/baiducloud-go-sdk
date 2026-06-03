package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeCacheNodeResponse struct {
	bce.BaseResponse
	CacheNodeInfo *CacheNodeInfo `json:"cacheNodeInfo,omitempty"`
}
