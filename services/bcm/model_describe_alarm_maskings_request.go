package bcm

type DescribeAlarmMaskingsRequest struct {
	MaskingName *string `json:"maskingName,omitempty"`
	MaskingId   *string `json:"maskingId,omitempty"`
	Order       *string `json:"order,omitempty"`
	OrderBy     *string `json:"orderBy,omitempty"`
	PageNo      *int32  `json:"pageNo,omitempty"`
	PageSize    *int32  `json:"pageSize,omitempty"`
}
