package bls

type ListAlarmPolicyRequest struct {
	PolicyNamePattern   *string `json:"policyNamePattern,omitempty"`
	PolicyIdPattern     *string `json:"policyIdPattern,omitempty"`
	LogStoreNamePattern *string `json:"logStoreNamePattern,omitempty"`
	State               *string `json:"state,omitempty"`
	NoticeState         *string `json:"noticeState,omitempty"`
	OrderBy             *string `json:"orderBy,omitempty"`
	Order               *string `json:"order,omitempty"`
	PageNo              *int32  `json:"pageNo,omitempty"`
	PageSize            *int32  `json:"pageSize,omitempty"`
}
