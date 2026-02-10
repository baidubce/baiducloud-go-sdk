package eip

type BindTbspProtectionObjectRequest struct {
	Id          *string   `json:"-"`
	ClientToken *string   `json:"-"`
	IpList      []*string `json:"ipList,omitempty"`
}
