package aihc

type DataSource struct {
	AihcType   *string `json:"type,omitempty"`
	Name       *string `json:"name,omitempty"`
	SourcePath *string `json:"sourcePath,omitempty"`
	MountPath  *string `json:"mountPath,omitempty"`
	Options    *Option `json:"options,omitempty"`
}
