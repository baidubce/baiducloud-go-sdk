package privatezone

type PrivateZone struct {
	ZoneId      *string `json:"zoneId,omitempty"`
	ZoneName    *string `json:"zoneName,omitempty"`
	RecordCount *int32  `json:"recordCount,omitempty"`
	CreateTime  *string `json:"createTime,omitempty"`
	UpdateTime  *string `json:"updateTime,omitempty"`
}
