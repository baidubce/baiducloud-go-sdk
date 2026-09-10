package cce

type GetTaskListV2Request struct {
	TaskType      *string `json:"-"`
	TargetID      *string `json:"-"`
	OperationType *string `json:"-"`
	Phase         *string `json:"-"`
	Order         *string `json:"-"`
	OrderBy       *string `json:"-"`
	PageNo        *int32  `json:"-"`
	PageSize      *int32  `json:"-"`
}
