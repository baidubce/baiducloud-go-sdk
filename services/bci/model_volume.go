package bci

type Volume struct {
	Nfs        []*NfsVolume        `json:"nfs,omitempty"`
	EmptyDir   []*EmptyDirVolume   `json:"emptyDir,omitempty"`
	ConfigFile []*ConfigFileVolume `json:"configFile,omitempty"`
}
