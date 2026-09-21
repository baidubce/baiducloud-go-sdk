package dbsc

type MysqlDatabaseSpaceModel struct {
	DatabaseName  *string  `json:"databaseName,omitempty"`
	TotalSpace    *int32   `json:"totalSpace,omitempty"`
	DataSpace     *int32   `json:"dataSpace,omitempty"`
	IndexSpace    *int32   `json:"indexSpace,omitempty"`
	FreeSpace     *int32   `json:"freeSpace,omitempty"`
	FreeRate      *float64 `json:"freeRate,omitempty"`
	UsageRate     *float64 `json:"usageRate,omitempty"`
	Rows          *int32   `json:"rows,omitempty"`
	PhysicalSpace *int32   `json:"physicalSpace,omitempty"`
	TableCount    *int32   `json:"tableCount,omitempty"`
}
