package oos

type GetTaskDetailV2Request struct {
	DagId          *string `json:"-"`
	TaskId         *string `json:"-"`
	IgnoreChildren *string `json:"-"`
	Locale         *string `json:"-"`
}
