package rds

type Parameter struct {
	Name        *string `json:"name,omitempty"`
	Value       *string `json:"value,omitempty"`
	ApplyMethod *string `json:"applyMethod,omitempty"`
	Etag        *string `json:"etag,omitempty"`
}
