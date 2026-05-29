package bls

type LogRecord struct {
	Message   *string `json:"message,omitempty"`
	Timestamp *int64  `json:"timestamp,omitempty"`
}
