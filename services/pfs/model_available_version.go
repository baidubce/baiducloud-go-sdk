package pfs

type AvailableVersion struct {
	VersionNumber     *string                      `json:"versionNumber,omitempty"`
	MountServiceInfos []*AvailableMountServiceInfo `json:"mountServiceInfos,omitempty"`
}
