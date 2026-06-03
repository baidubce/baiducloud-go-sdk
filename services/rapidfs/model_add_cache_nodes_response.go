package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AddCacheNodesResponse struct {
	bce.BaseResponse
	CacheNodeIds []*string `json:"cacheNodeIds,omitempty"`
}
