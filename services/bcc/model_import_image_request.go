package bcc

type ImportImageRequest struct {
	OsName         *string `json:"osName,omitempty"`
	OsArch         *string `json:"osArch,omitempty"`
	OsType         *string `json:"osType,omitempty"`
	OsVersion      *string `json:"osVersion,omitempty"`
	Name           *string `json:"name,omitempty"`
	BosUrl         *string `json:"bosUrl,omitempty"`
	Detection      *bool   `json:"detection,omitempty"`
	GenerationType *string `json:"generationType,omitempty"`
}
