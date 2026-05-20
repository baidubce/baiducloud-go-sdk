package dns

type UpdateLineGroupRequest struct {
	LineId      *int32    `json:"-"`
	ClientToken *string   `json:"-"`
	Name        *string   `json:"name,omitempty"`
	Lines       []*string `json:"lines,omitempty"`
}
