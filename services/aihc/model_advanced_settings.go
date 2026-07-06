package aihc

type AdvancedSettings struct {
	RuntimeEnv            *string `json:"runtimeEnv,omitempty"`
	SubmitterBackoffLimit *int32  `json:"SubmitterBackoffLimit,omitempty"`
}
