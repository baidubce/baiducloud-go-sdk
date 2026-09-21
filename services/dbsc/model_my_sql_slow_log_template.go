package dbsc

type MySQLSlowLogTemplate struct {
	Fringerprint    *string `json:"fringerprint,omitempty"`
	FringerprintMD5 *string `json:"fringerprintMD5,omitempty"`
	DbName          *string `json:"dbName,omitempty"`
	ExecuteTimes    *int32  `json:"executeTimes,omitempty"`
	DurationSum     *int32  `json:"durationSum,omitempty"`
	DurationMax     *int32  `json:"durationMax,omitempty"`
	DurationMin     *int32  `json:"durationMin,omitempty"`
	DurationAvg     *int32  `json:"durationAvg,omitempty"`
	LockTimeSum     *int32  `json:"lockTimeSum,omitempty"`
	LockTimeMax     *int32  `json:"lockTimeMax,omitempty"`
	LockTimeMin     *int32  `json:"lockTimeMin,omitempty"`
	LockTimeAvg     *int32  `json:"lockTimeAvg,omitempty"`
	ScanRowsSum     *int32  `json:"scanRowsSum,omitempty"`
	ScanRowsMax     *int32  `json:"scanRowsMax,omitempty"`
	ScanRowsMin     *int32  `json:"scanRowsMin,omitempty"`
	ScanRowsAvg     *int32  `json:"scanRowsAvg,omitempty"`
	ReturnRowsSum   *int32  `json:"returnRowsSum,omitempty"`
	ReturnRowsMax   *int32  `json:"returnRowsMax,omitempty"`
	ReturnRowsMin   *int32  `json:"returnRowsMin,omitempty"`
	ReturnRowsAvg   *int32  `json:"returnRowsAvg,omitempty"`
}
