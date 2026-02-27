package eip

type OptionalReleaseEipRequest struct {
	Eip              *string `json:"-"`
	ReleaseToRecycle *bool   `json:"-"`
	ClientToken      *string `json:"-"`
}
