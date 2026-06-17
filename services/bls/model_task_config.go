package bls

type TaskConfig struct {
	SrcConfig  *SrcConfig  `json:"srcConfig,omitempty"`
	DestConfig *DestConfig `json:"destConfig,omitempty"`
}
