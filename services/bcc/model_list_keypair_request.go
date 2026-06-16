package bcc

type ListKeypairRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
