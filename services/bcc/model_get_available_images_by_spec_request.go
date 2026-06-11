package bcc

type GetAvailableImagesBySpecRequest struct {
	Spec    *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
	OsName  *string `json:"-"`
}
