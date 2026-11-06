package rds

type OpenapiErrorlog struct {
	ErrorlogId          *string `json:"errorlogId,omitempty"`
	ErrorlogSizeInBytes *int64  `json:"errorlogSizeInBytes,omitempty"`
	ErrorlogStartTime   *string `json:"errorlogStartTime,omitempty"`
	ErrorlogEndTime     *string `json:"errorlogEndTime,omitempty"`
}
