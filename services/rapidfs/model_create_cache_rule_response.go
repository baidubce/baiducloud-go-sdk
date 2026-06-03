package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateCacheRuleResponse struct {
	bce.BaseResponse
	CacheRuleId *string `json:"cacheRuleId,omitempty"`
}
