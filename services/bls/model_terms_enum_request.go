package bls

type TermsEnumRequest struct {
	Name        *string      `json:"-"`
	Field       *string      `json:"field,omitempty"`
	BlsString   *string      `json:"string,omitempty"`
	Size        *int32       `json:"size,omitempty"`
	IndexFilter *interface{} `json:"index_filter,omitempty"`
}
