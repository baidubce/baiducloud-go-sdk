package dbsc

type MysqlTableSpaceModel struct {
	DatabaseName  *string  `json:"databaseName,omitempty"`
	TableName     *string  `json:"tableName,omitempty"`
	Engine        *string  `json:"engine,omitempty"`
	TotalSpace    *int64   `json:"totalSpace,omitempty"`
	DataSpace     *int64   `json:"dataSpace,omitempty"`
	IndexSpace    *int64   `json:"indexSpace,omitempty"`
	FreeSpace     *int32   `json:"freeSpace,omitempty"`
	FreeRate      *float64 `json:"freeRate,omitempty"`
	UsageRate     *float64 `json:"usageRate,omitempty"`
	Rows          *int32   `json:"rows,omitempty"`
	AvgRowLength  *int32   `json:"avgRowLength,omitempty"`
	PhysicalSpace *int64   `json:"physicalSpace,omitempty"`
}
