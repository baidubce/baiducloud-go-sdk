package bci

type NfsVolume struct {
	Name     *string `json:"name,omitempty"`
	Server   *string `json:"server,omitempty"`
	Path     *string `json:"path,omitempty"`
	ReadOnly *bool   `json:"readOnly,omitempty"`
}
