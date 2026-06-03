package ccr

type CreateTemporaryPasswordRequest struct {
	InstanceId *string `json:"-"`
	Duration   *int32  `json:"duration,omitempty"`
}
