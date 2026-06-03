package pfs

type ExecuteInfo struct {
	JobId       *string `json:"jobId,omitempty"`
	TriggerTime *string `json:"triggerTime,omitempty"`
	Status      *int32  `json:"status,omitempty"`
	Errmsg      *string `json:"errmsg,omitempty"`
	Report      *string `json:"report,omitempty"`
}
