package bci

type ConfigFileVolume struct {
	Name        *string             `json:"name,omitempty"`
	DefaultMode *int32              `json:"defaultMode,omitempty"`
	ConfigFiles []*ConfigFileDetail `json:"configFiles,omitempty"`
}
