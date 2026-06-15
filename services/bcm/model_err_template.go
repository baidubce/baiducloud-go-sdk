package bcm

type ErrTemplate struct {
	Index    *int32    `json:"index,omitempty"`
	Template *Template `json:"template,omitempty"`
	Message  *string   `json:"message,omitempty"`
}
