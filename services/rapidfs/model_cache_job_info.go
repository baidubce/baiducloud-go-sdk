package rapidfs

type CacheJobInfo struct {
	CacheJobId *string `json:"cacheJobId,omitempty"`
	Status     *string `json:"status,omitempty"`
	Size       *int64  `json:"size,omitempty"`
	StartTime  *string `json:"startTime,omitempty"`
	EndTime    *string `json:"endTime,omitempty"`
}
