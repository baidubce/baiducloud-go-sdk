package bcc

type ListImagesRequest struct {
	Marker    *string `json:"-"`
	MaxKeys   *int32  `json:"-"`
	ImageType *string `json:"-"`
	ImageName *string `json:"-"`
}
