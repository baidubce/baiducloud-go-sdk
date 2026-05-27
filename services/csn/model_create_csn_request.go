package csn

type CreateCsnRequest struct {
	ClientToken *string     `json:"-"`
	Name        *string     `json:"name,omitempty"`
	Description *string     `json:"description,omitempty"`
	Tags        []*TagModel `json:"tags,omitempty"`
}
