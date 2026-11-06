package rds

type Binlog struct {
	DownloadUrl     *string `json:"downloadUrl,omitempty"`
	DownloadExpires *string `json:"downloadExpires,omitempty"`
}
