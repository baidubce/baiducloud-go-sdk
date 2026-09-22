package face

type IdMatchStatusResult struct {
	VerifyStatus *int32 `json:"verify_status,omitempty"`
	IsNewest     *int32 `json:"isNewest,omitempty"`
	IsLosted     *int32 `json:"isLosted,omitempty"`
	IsExpired    *int32 `json:"isExpired,omitempty"`
	Hjzt         *int32 `json:"hjzt,omitempty"`
}
