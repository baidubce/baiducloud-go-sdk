package bcc

type SystemImageModel struct {
	ImageId      *string `json:"imageId,omitempty"`
	ImageName    *string `json:"imageName,omitempty"`
	OsType       *string `json:"osType,omitempty"`
	OsVersion    *string `json:"osVersion,omitempty"`
	OsArch       *string `json:"osArch,omitempty"`
	OsName       *string `json:"osName,omitempty"`
	OsLang       *string `json:"osLang,omitempty"`
	MinSizeInGiB *int32  `json:"minSizeInGiB,omitempty"`
}
