package eip

type DdosAttackRecordModel struct {
	Ip             *string   `json:"ip,omitempty"`
	StartTime      *string   `json:"startTime,omitempty"`
	EndTime        *string   `json:"endTime,omitempty"`
	AttackType     []*string `json:"attackType,omitempty"`
	AttackPeakMbps *int64    `json:"attackPeakMbps,omitempty"`
	AttackPeakPps  *int64    `json:"attackPeakPps,omitempty"`
	AttackPeakQps  *int64    `json:"attackPeakQps,omitempty"`
	AttackStatus   *string   `json:"attackStatus,omitempty"`
}
