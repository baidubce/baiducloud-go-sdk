package rds

type OpenapiSlowlog struct {
	SlowlogId          *string `json:"slowlogId,omitempty"`
	SlowlogSizeInBytes *int64  `json:"slowlogSizeInBytes,omitempty"`
	SlowlogStartTime   *string `json:"slowlogStartTime,omitempty"`
	SlowlogEndTime     *string `json:"slowlogEndTime,omitempty"`
}
