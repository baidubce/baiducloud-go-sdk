package ccr

type TagScanOverview struct {
	EndTime    *string           `json:"endTime,omitempty"`
	Fixable    *int32            `json:"fixable,omitempty"`
	ReportId   *string           `json:"reportId,omitempty"`
	ScanStatus *string           `json:"scanStatus,omitempty"`
	Severity   *string           `json:"severity,omitempty"`
	StartTime  *string           `json:"startTime,omitempty"`
	Summary    *map[string]int32 `json:"summary,omitempty"`
	Total      *int32            `json:"total,omitempty"`
}
