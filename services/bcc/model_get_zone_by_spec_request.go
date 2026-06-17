package bcc

type GetZoneBySpecRequest struct {
	InstanceType *string `json:"-"`
	ProductType  *string `json:"-"`
	Spec         *string `json:"-"`
	SpecId       *string `json:"-"`
}
