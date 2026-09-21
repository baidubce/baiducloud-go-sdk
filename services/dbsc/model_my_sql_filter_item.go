package dbsc

type MySQLFilterItem struct {
	FilterId     *string `json:"filterId,omitempty"`
	FilterKey    *string `json:"filterKey,omitempty"`
	FilterLimit  *int32  `json:"filterLimit,omitempty"`
	FilterType   *string `json:"filterType,omitempty"`
	FilterStatus *string `json:"filterStatus,omitempty"`
	CreateTime   *string `json:"createTime,omitempty"`
	UpdateTime   *string `json:"updateTime,omitempty"`
}
