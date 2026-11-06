package rds

type MultiDownloads struct {
	DbName       *string `json:"dbName,omitempty"`
	DownloadName *string `json:"downloadName,omitempty"`
	DownloadSize *int64  `json:"downloadSize,omitempty"`
	DownloadUrl  *string `json:"downloadUrl,omitempty"`
}
