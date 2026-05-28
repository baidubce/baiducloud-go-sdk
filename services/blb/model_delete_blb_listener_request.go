package blb

type DeleteBlbListenerRequest struct {
	BlbId        *string          `json:"-"`
	ClientToken  *string          `json:"-"`
	PortList     []*int32         `json:"portList,omitempty"`
	PortTypeList []*PortTypeModel `json:"portTypeList,omitempty"`
}
