package bcc

type FileSystemModel struct {
	FsId     *string `json:"fsId,omitempty"`
	MountAds *string `json:"mountAds,omitempty"`
	Path     *string `json:"path,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
}
