package eip

type MoveInEipsRequest struct {
	Id          *string   `json:"-"`
	ClientToken *string   `json:"-"`
	Eips        []*string `json:"eips,omitempty"`
}
