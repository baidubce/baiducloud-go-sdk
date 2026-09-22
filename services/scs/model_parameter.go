package scs

type Parameter struct {
	ScsDefault   *string `json:"default,omitempty"`
	ForceRestart *int32  `json:"forceRestart,omitempty"`
	Name         *string `json:"name,omitempty"`
	Value        *string `json:"value,omitempty"`
}
