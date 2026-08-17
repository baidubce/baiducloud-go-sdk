package aigw

type ConsumerCredentialLocation struct {
	InHeader *bool     `json:"inHeader,omitempty"`
	InQuery  *bool     `json:"inQuery,omitempty"`
	KeyNames []*string `json:"keyNames,omitempty"`
}
