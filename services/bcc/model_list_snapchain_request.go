package bcc

type ListSnapchainRequest struct {
	OrderBy  *string `json:"-"`
	Order    *string `json:"-"`
	PageSize *int32  `json:"-"`
	PageNo   *int32  `json:"-"`
	VolumeId *string `json:"-"`
}
