package bcc

type ModifyInstanceHostnameRequest struct {
	InstanceId            *string `json:"-"`
	Reboot                *bool   `json:"reboot,omitempty"`
	IsOpenHostnameDomain  *bool   `json:"isOpenHostnameDomain,omitempty"`
	Hostname              *string `json:"hostname,omitempty"`
	IsAllowHostnameRepeat *bool   `json:"isAllowHostnameRepeat,omitempty"`
}
