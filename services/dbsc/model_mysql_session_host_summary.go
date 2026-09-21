package dbsc

type MysqlSessionHostSummary struct {
	Host                     *string  `json:"host,omitempty"`
	ActiveAverageExecuteTime *float64 `json:"activeAverageExecuteTime,omitempty"`
	ActiveMaxExecuteTime     *float64 `json:"activeMaxExecuteTime,omitempty"`
	ActiveTotalCount         *int32   `json:"activeTotalCount,omitempty"`
	ActiveTotalExecuteTime   *float64 `json:"activeTotalExecuteTime,omitempty"`
	AverageExecuteTime       *float64 `json:"averageExecuteTime,omitempty"`
	MaxExecuteTime           *float64 `json:"maxExecuteTime,omitempty"`
	TotalCount               *int32   `json:"totalCount,omitempty"`
	TotalExecuteTime         *float64 `json:"totalExecuteTime,omitempty"`
}
