package oos

type DagActionModel struct {
	Pause  *bool `json:"pause,omitempty"`
	Resume *bool `json:"resume,omitempty"`
}
