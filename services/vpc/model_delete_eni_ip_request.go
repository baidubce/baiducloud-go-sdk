package vpc

type DeleteEniIpRequest struct {
	EniId            *string `json:"-"`
	PrivateIpAddress *string `json:"-"`
	ClientToken      *string `json:"-"`
}
