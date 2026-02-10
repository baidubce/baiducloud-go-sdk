package eip

type ListTbspRequest struct {
	Id      *string `json:"-"`
	Name    *string `json:"-"`
	Status  *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
