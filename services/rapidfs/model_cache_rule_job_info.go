package rapidfs

type CacheRuleJobInfo struct {
	CacheRuleJobId *string `json:"cacheRuleJobId,omitempty"`
	Status         *string `json:"status,omitempty"`
	Size           *int64  `json:"size,omitempty"`
	StartTime      *string `json:"startTime,omitempty"`
	EndTime        *string `json:"endTime,omitempty"`
}
