package dns

type AddLineGroupRequest struct {
	ClientToken *string   `json:"-"`
	Name        *string   `json:"name,omitempty"`
	Lines       []*string `json:"lines,omitempty"`
}
