package bls

type CreateDownloadTaskRequest struct {
	Name           *string `json:"name,omitempty"`
	Project        *string `json:"project,omitempty"`
	LogStoreName   *string `json:"logStoreName,omitempty"`
	LogStreamName  *string `json:"logStreamName,omitempty"`
	Query          *string `json:"query,omitempty"`
	QueryStartTime *string `json:"queryStartTime,omitempty"`
	QueryEndTime   *string `json:"queryEndTime,omitempty"`
	Format         *string `json:"format,omitempty"`
	Limit          *int32  `json:"limit,omitempty"`
	Order          *string `json:"order,omitempty"`
	FileDir        *string `json:"fileDir,omitempty"`
}
