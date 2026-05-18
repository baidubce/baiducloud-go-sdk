package eip

type ListEipBpRequest struct {
	Marker   *string `json:"-"`
	MaxKeys  *int32  `json:"-"`
	Id       *string `json:"-"`
	Name     *string `json:"-"`
	BindType *string `json:"-"`
	Type     *string `json:"-"`
}
