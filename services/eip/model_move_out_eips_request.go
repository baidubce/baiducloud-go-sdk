package eip

type MoveOutEipsRequest struct {
	Id          *string            `json:"-"`
	ClientToken *string            `json:"-"`
	MoveOutEips []*EipMoveOutModel `json:"moveOutEips,omitempty"`
}
