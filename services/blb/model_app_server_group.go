package blb

type AppServerGroup struct {
	Id                      *string               `json:"id,omitempty"`
	Name                    *string               `json:"name,omitempty"`
	Desc                    *string               `json:"desc,omitempty"`
	Status                  *string               `json:"status,omitempty"`
	PreserveClientIpEnabled *bool                 `json:"preserveClientIpEnabled,omitempty"`
	PortList                []*AppServerGroupPort `json:"portList,omitempty"`
}
