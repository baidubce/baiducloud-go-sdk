package apm

type StatementDetail struct {
	Id        *string `json:"id,omitempty"`
	Statement *string `json:"statement,omitempty"`
	Service   *string `json:"service,omitempty"`
}
