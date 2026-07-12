package oos

type KeyValuePair struct {
	Key   *string      `json:"key,omitempty"`
	Value *interface{} `json:"value,omitempty"`
}
