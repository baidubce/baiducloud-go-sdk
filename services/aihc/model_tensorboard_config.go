package aihc

type TensorboardConfig struct {
	Enable  *bool   `json:"enable,omitempty"`
	LogPath *string `json:"logPath,omitempty"`
}
