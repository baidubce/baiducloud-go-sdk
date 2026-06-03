package rapidfs

type DiskInfo struct {
	Dev          *string `json:"dev,omitempty"`
	MountPath    *string `json:"mountPath,omitempty"`
	DiskQuotaGiB *int64  `json:"diskQuotaGiB,omitempty"`
	Format       *bool   `json:"format,omitempty"`
	DiskType     *string `json:"diskType,omitempty"`
}
