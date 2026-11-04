package dns

type UpdateLineGroupRequest struct {
	LineId      *string   `json:"-"`
	ClientToken *string   `json:"-"`
	Name        *string   `json:"name,omitempty"`
	Lines       []*string `json:"lines,omitempty"`
}
