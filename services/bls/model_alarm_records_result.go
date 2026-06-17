package bls

type AlarmRecordsResult struct {
	PageNo     *int32   `json:"pageNo,omitempty"`
	PageSize   *int32   `json:"pageSize,omitempty"`
	TotalCount *int32   `json:"totalCount,omitempty"`
	Alarms     []*Alarm `json:"alarms,omitempty"`
}
