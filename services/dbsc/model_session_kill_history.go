package dbsc

type SessionKillHistory struct {
	SessionId          *int32  `json:"sessionId,omitempty"`
	SessionUser        *string `json:"sessionUser,omitempty"`
	SessionHost        *string `json:"sessionHost,omitempty"`
	SessionDb          *string `json:"sessionDb,omitempty"`
	SessionCommand     *string `json:"sessionCommand,omitempty"`
	SessionExecuteTime *int32  `json:"sessionExecuteTime,omitempty"`
	SessionState       *string `json:"sessionState,omitempty"`
	SessionSql         *string `json:"sessionSql,omitempty"`
	Status             *int32  `json:"status,omitempty"`
	StatusDesc         *string `json:"statusDesc,omitempty"`
	StatusInfo         *string `json:"statusInfo,omitempty"`
	OperateTime        *string `json:"operateTime,omitempty"`
}
