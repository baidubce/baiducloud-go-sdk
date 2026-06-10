package bcc

type ModifyCdsAttributeRequest struct {
	VolumeId           *string `json:"-"`
	CdsName            *string `json:"cdsName,omitempty"`
	Desc               *string `json:"desc,omitempty"`
	DeleteWithInstance *bool   `json:"deleteWithInstance,omitempty"`
	DeleteAutoSnapshot *bool   `json:"deleteAutoSnapshot,omitempty"`
}
