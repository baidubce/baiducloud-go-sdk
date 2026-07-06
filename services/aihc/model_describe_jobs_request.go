package aihc

type DescribeJobsRequest struct {
	ResourcePoolId *string `json:"-"`
	QueueID        *string `json:"-"`
	Queue          *string `json:"queue,omitempty"`
	Status         *string `json:"status,omitempty"`
	KeywordType    *string `json:"keywordType,omitempty"`
	Keyword        *string `json:"keyword,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty"`
	Order          *string `json:"order,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty"`
}
