package dbsc

type MongoDBSlowLogTemplate struct {
	FingerprintMd5 *string `json:"fingerprintMd5,omitempty"`
	Fingerprint    *string `json:"fingerprint,omitempty"`
	Namespace      *string `json:"namespace,omitempty"`
	ExecuteTimes   *string `json:"executeTimes,omitempty"`
	DurationSum    *int32  `json:"durationSum,omitempty"`
	DurationMax    *int32  `json:"durationMax,omitempty"`
	DurationMin    *int32  `json:"durationMin,omitempty"`
	DurationAvg    *int32  `json:"durationAvg,omitempty"`
	KeyScanRowsSum *int32  `json:"keyScanRowsSum,omitempty"`
	KeyScanRowsMax *int32  `json:"keyScanRowsMax,omitempty"`
	KeyScanRowsMin *int32  `json:"keyScanRowsMin,omitempty"`
	KeyScanRowsAvg *int32  `json:"keyScanRowsAvg,omitempty"`
	ScanRowsSum    *int32  `json:"scanRowsSum,omitempty"`
	ScanRowsMax    *int32  `json:"scanRowsMax,omitempty"`
	ScanRowsMin    *int32  `json:"scanRowsMin,omitempty"`
	ScanRowsAvg    *int32  `json:"scanRowsAvg,omitempty"`
	ReturnRowsSum  *int32  `json:"returnRowsSum,omitempty"`
	ReturnRowsMax  *int32  `json:"returnRowsMax,omitempty"`
	ReturnRowsMin  *int32  `json:"returnRowsMin,omitempty"`
	ReturnRowsAvg  *int32  `json:"returnRowsAvg,omitempty"`
}
