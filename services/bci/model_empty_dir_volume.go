package bci

type EmptyDirVolume struct {
	Name      *string  `json:"name,omitempty"`
	Medium    *string  `json:"medium,omitempty"`
	SizeLimit *float32 `json:"sizeLimit,omitempty"`
}
