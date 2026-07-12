package oos

type Option struct {
	Label *string      `json:"label,omitempty"`
	Value *interface{} `json:"value,omitempty"`
}
