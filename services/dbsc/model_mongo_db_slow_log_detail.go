package dbsc

type MongoDBSlowLogDetail struct {
	Uuid           *string `json:"uuid,omitempty"`
	ClientIp       *string `json:"clientIp,omitempty"`
	ClientPort     *int32  `json:"clientPort,omitempty"`
	User           *string `json:"user,omitempty"`
	ConnectionId   *int32  `json:"connectionId,omitempty"`
	CurrentDB      *string `json:"currentDB,omitempty"`
	Duration       *int32  `json:"duration,omitempty"`
	Start          *string `json:"start,omitempty"`
	End            *string `json:"end,omitempty"`
	Fingerprint    *string `json:"fingerprint,omitempty"`
	FingerprintMd5 *string `json:"fingerprintMd5,omitempty"`
	Query          *string `json:"query,omitempty"`
	SqlCommand     *string `json:"sqlCommand,omitempty"`
	ScanRows       *int32  `json:"scanRows,omitempty"`
	ReturnRows     *int32  `json:"returnRows,omitempty"`
	KeyScanRows    *int32  `json:"keyScanRows,omitempty"`
	ResultLen      *int32  `json:"resultLen,omitempty"`
	PlanSummary    *string `json:"planSummary,omitempty"`
	Namespace      *string `json:"namespace,omitempty"`
}
