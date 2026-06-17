package bcc

type ListFlavorSpecRequest struct {
	ZoneName *string `json:"-"`
	Specs    *string `json:"-"`
	SpecIds  *string `json:"-"`
}
