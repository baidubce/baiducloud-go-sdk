package eip

type SelectiveReleaseOfEipRequest struct {
	Eip              *string `json:"-"`
	ReleaseToRecycle *bool   `json:"-"`
	ClientToken      *string `json:"-"`
}
