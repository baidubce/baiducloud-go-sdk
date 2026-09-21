package dbsc

type APIPGSlowLogTemplateItem struct {
	FingerprintMD5 *string `json:"fingerprintMD5,omitempty"`
	Fingerprint    *string `json:"fingerprint,omitempty"`
	DbName         *string `json:"dbName,omitempty"`
	ExecuteTimes   *int64  `json:"executeTimes,omitempty"`
	DurationSum    *int64  `json:"durationSum,omitempty"`
	DurationMax    *int64  `json:"durationMax,omitempty"`
	DurationMin    *int64  `json:"durationMin,omitempty"`
	DurationAvg    *int64  `json:"durationAvg,omitempty"`
}
