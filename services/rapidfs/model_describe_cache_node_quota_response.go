package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeCacheNodeQuotaResponse struct {
	bce.BaseResponse
	CacheNodeQuotaInfo *CacheNodeQuotaInfo `json:"cacheNodeQuotaInfo,omitempty"`
}
