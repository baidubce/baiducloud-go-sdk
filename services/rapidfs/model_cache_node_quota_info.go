package rapidfs

type CacheNodeQuotaInfo struct {
	Used  *int32 `json:"used,omitempty"`
	Quota *int32 `json:"quota,omitempty"`
}
