package cloudassistant

type FileUpload struct {
	Os            *string `json:"os,omitempty"`
	Content       *string `json:"content,omitempty"`
	Filename      *string `json:"filename,omitempty"`
	Filepath      *string `json:"filepath,omitempty"`
	BosBucketName *string `json:"bosBucketName,omitempty"`
	BosFilePath   *string `json:"bosFilePath,omitempty"`
	BosEtag       *string `json:"bosEtag,omitempty"`
	User          *string `json:"user,omitempty"`
	Group         *string `json:"group,omitempty"`
	Mode          *string `json:"mode,omitempty"`
	Overwrite     *bool   `json:"overwrite,omitempty"`
}
