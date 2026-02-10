package eip

type RenewTbspRequest struct {
	Id            *string `json:"-"`
	ClientToken   *string `json:"-"`
	RenewTime     *int32  `json:"renewTime,omitempty"`
	RenewTimeUnit *string `json:"renewTimeUnit,omitempty"`
}
