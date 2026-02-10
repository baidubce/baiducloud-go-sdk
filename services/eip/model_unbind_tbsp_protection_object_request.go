package eip

type UnbindTbspProtectionObjectRequest struct {
	Id          *string   `json:"-"`
	ClientToken *string   `json:"-"`
	IpList      []*string `json:"ipList,omitempty"`
}
