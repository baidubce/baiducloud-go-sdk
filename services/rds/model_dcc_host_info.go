package rds

type DccHostInfo struct {
	Master *DccHostId `json:"master,omitempty"`
	Backup *DccHostId `json:"backup,omitempty"`
}
