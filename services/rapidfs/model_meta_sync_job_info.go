package rapidfs

type MetaSyncJobInfo struct {
	MetaSyncJobId *string `json:"metaSyncJobId,omitempty"`
	Status        *string `json:"status,omitempty"`
	StartTime     *string `json:"startTime,omitempty"`
	EndTime       *string `json:"endTime,omitempty"`
}
