package dbsc

type MySQLSlowLogDetail struct {
	ClientIp       *string `json:"clientIp,omitempty"`
	ClientPort     *int32  `json:"clientPort,omitempty"`
	User           *string `json:"user,omitempty"`
	ConnectionId   *int32  `json:"connectionId,omitempty"`
	CurrentDB      *string `json:"currentDB,omitempty"`
	Duration       *int32  `json:"duration,omitempty"`
	LockTime       *int32  `json:"lockTime,omitempty"`
	Start          *string `json:"start,omitempty"`
	End            *string `json:"end,omitempty"`
	Fingerprint    *string `json:"fingerprint,omitempty"`
	FingerprintMd5 *string `json:"fingerprintMd5,omitempty"`
	Method         *string `json:"method,omitempty"`
	Query          *string `json:"query,omitempty"`
	AffectedRows   *int32  `json:"affectedRows,omitempty"`
	ScanRows       *int32  `json:"scanRows,omitempty"`
	ReturnRows     *int32  `json:"returnRows,omitempty"`
	SqlType        *string `json:"sqlType,omitempty"`
}
