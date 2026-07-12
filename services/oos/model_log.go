package oos

type Log struct {
	Timestamp *string      `json:"timestamp,omitempty"`
	Level     *string      `json:"level,omitempty"`
	Msg       *string      `json:"msg,omitempty"`
	Tags      *interface{} `json:"tags,omitempty"`
}
