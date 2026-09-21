package dbsc

type SCSSlowLogInfo struct {
	Content     *string `json:"content,omitempty"`
	LogDuration *int64  `json:"logDuration,omitempty"`
	LogKey      *string `json:"logKey,omitempty"`
	LogSql      *string `json:"logSql,omitempty"`
	LogTime     *string `json:"logTime,omitempty"`
	SlowLogId   *int64  `json:"slowLogId,omitempty"`
	ClientIp    *string `json:"clientIp,omitempty"`
	ClientIP    *string `json:"ClientIP,omitempty"`
}
