package oos

type GetTaskDetailV2Request struct {
	Namespace      *string `json:"-"`
	DagId          *string `json:"-"`
	TaskId         *string `json:"-"`
	IgnoreChildren *string `json:"-"`
	Locale         *string `json:"-"`
}
