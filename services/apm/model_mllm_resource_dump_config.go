package apm

type MllmResourceDumpConfig struct {
	RetentionDays *int32  `json:"retentionDays,omitempty"`
	Bucket        *string `json:"bucket,omitempty"`
}
