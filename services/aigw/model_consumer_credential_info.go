package aigw

type ConsumerCredentialInfo struct {
	IdentityId   *int64  `json:"identityId,omitempty"`
	Name         *string `json:"name,omitempty"`
	Value        *string `json:"value,omitempty"`
	MaskedValue  *string `json:"maskedValue,omitempty"`
	GenerateMode *string `json:"generateMode,omitempty"`
	InHeader     *bool   `json:"inHeader,omitempty"`
	InQuery      *bool   `json:"inQuery,omitempty"`
}
