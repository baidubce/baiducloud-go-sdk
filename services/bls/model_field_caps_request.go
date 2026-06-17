package bls

type FieldCapsRequest struct {
	Name   *string   `json:"-"`
	Fields []*string `json:"fields,omitempty"`
}
