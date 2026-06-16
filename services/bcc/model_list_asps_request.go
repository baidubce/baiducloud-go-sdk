package bcc

type ListAspsRequest struct {
	Marker     *string `json:"-"`
	MaxKeys    *int32  `json:"-"`
	AspName    *string `json:"-"`
	VolumeName *string `json:"-"`
}
