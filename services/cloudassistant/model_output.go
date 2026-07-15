package cloudassistant

type Output struct {
	ExitCode    *int32  `json:"exitCode,omitempty"`
	Stderr      *string `json:"stderr,omitempty"`
	Stdout      *string `json:"stdout,omitempty"`
	IsTruncated *bool   `json:"isTruncated,omitempty"`
}
