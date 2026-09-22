package scs

type LogItem struct {
	LogId           *string `json:"logId,omitempty"`
	LogSizeInBytes  *int32  `json:"logSizeInBytes,omitempty"`
	LogStartTime    *string `json:"logStartTime,omitempty"`
	LogEndTime      *string `json:"logEndTime,omitempty"`
	DownloadUrl     *string `json:"downloadUrl,omitempty"`
	DownloadExpires *string `json:"downloadExpires,omitempty"`
}
