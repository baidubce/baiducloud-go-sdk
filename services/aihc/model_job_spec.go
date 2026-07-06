package aihc

type JobSpec struct {
	Image       *string      `json:"image,omitempty"`
	ImageConfig *ImageConfig `json:"imageConfig,omitempty"`
	Replicas    *int32       `json:"replicas,omitempty"`
	Resources   []*Resource  `json:"resources,omitempty"`
	Envs        []*Env       `json:"envs,omitempty"`
	EnableRDMA  *bool        `json:"enableRDMA,omitempty"`
	HostNetwork *bool        `json:"hostNetwork,omitempty"`
}
