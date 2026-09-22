package scs

type GetDeploymentSetListRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
