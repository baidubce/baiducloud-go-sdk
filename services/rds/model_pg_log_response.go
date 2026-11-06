package rds

type PgLogResponse struct {
	PglogId          *string `json:"pglogId,omitempty"`
	PglogSizeInBytes *int64  `json:"pglogSizeInBytes,omitempty"`
	PglogStartTime   *string `json:"pglogStartTime,omitempty"`
	PglogEndTime     *string `json:"pglogEndTime,omitempty"`
}
