package eip

type DisableTbspIpCleanRequest struct {
	Id              *string `json:"-"`
	ClientToken     *string `json:"-"`
	Ip              *string `json:"ip,omitempty"`
	TurnOffDuration *int32  `json:"turnOffDuration,omitempty"`
}
