package blb

type UpgradeLbdcRequest struct {
	Id          *string `json:"-"`
	ClientToken *string `json:"-"`
	CcuCount    *int32  `json:"ccuCount,omitempty"`
}
