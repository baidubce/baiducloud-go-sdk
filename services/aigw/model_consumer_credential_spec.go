package aigw

type ConsumerCredentialSpec struct {
	Name         *string   `json:"name,omitempty"`
	GenerateMode *string   `json:"generateMode,omitempty"`
	Value        *string   `json:"value,omitempty"`
	InHeader     *bool     `json:"inHeader,omitempty"`
	InQuery      *bool     `json:"inQuery,omitempty"`
	KeyNames     []*string `json:"keyNames,omitempty"`
	Description  *string   `json:"description,omitempty"`
}
