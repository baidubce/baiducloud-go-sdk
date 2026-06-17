package bls

type ListAlarmRecordRequest struct {
	PolicyNamePattern   *string `json:"policyNamePattern,omitempty"`
	PolicyIdPattern     *string `json:"policyIdPattern,omitempty"`
	LogStoreNamePattern *string `json:"logStoreNamePattern,omitempty"`
	Level               *string `json:"level,omitempty"`
	State               *string `json:"state,omitempty"`
	StartDateTime       *string `json:"startDateTime,omitempty"`
	EndDateTime         *string `json:"endDateTime,omitempty"`
	OrderBy             *string `json:"orderBy,omitempty"`
	Order               *string `json:"order,omitempty"`
	PageNo              *int32  `json:"pageNo,omitempty"`
	PageSize            *int32  `json:"pageSize,omitempty"`
}
