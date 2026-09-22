package scs

type LogListRequest struct {
	InstanceId *string `json:"-"`
	FileType   *string `json:"-"`
	StartTime  *string `json:"-"`
	EndTime    *string `json:"-"`
}
