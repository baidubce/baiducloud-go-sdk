package ccr

type BuildHistory struct {
	Created    *string `json:"created,omitempty"`
	Comment    *string `json:"comment,omitempty"`
	CreatedBy  *string `json:"createdBy,omitempty"`
	EmptyLayer *bool   `json:"emptyLayer,omitempty"`
}
