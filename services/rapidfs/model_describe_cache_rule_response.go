package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeCacheRuleResponse struct {
	bce.BaseResponse
	CacheRuleInfo *CacheRuleInfo `json:"cacheRuleInfo,omitempty"`
}
