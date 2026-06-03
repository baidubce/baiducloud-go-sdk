package rapidfs

type CreateCacheRuleRequest struct {
	ClientToken     *string `json:"-"`
	InstanceId      *string `json:"instanceId,omitempty"`
	DataSrcId       *string `json:"dataSrcId,omitempty"`
	CacheRuleName   *string `json:"cacheRuleName,omitempty"`
	RapidfsType     *string `json:"type,omitempty"`
	Directory       *string `json:"directory,omitempty"`
	ExecuteOnCreate *bool   `json:"executeOnCreate,omitempty"`
	Description     *string `json:"description,omitempty"`
}
