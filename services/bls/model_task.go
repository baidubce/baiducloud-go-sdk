package bls

type Task struct {
	Uuid           *string `json:"uuid,omitempty"`
	Name           *string `json:"name,omitempty"`
	Project        *string `json:"project,omitempty"`
	LogStoreName   *string `json:"logStoreName,omitempty"`
	Query          *string `json:"query,omitempty"`
	QueryStartTime *string `json:"queryStartTime,omitempty"`
	QueryEndTime   *string `json:"queryEndTime,omitempty"`
	Format         *string `json:"format,omitempty"`
	Limit          *int32  `json:"limit,omitempty"`
	Order          *string `json:"order,omitempty"`
	State          *string `json:"state,omitempty"`
	FailedCode     *string `json:"failedCode,omitempty"`
	FailedMessage  *string `json:"failedMessage,omitempty"`
	WrittenRows    *int32  `json:"writtenRows,omitempty"`
	FileDir        *string `json:"fileDir,omitempty"`
	FileName       *string `json:"fileName,omitempty"`
	ExecStartTime  *string `json:"execStartTime,omitempty"`
	ExecEndTime    *string `json:"execEndTime,omitempty"`
	CreatedTime    *string `json:"createdTime,omitempty"`
	UpdatedTime    *string `json:"updatedTime,omitempty"`
}
