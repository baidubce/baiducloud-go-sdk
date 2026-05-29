package blb

type AppIpGroup struct {
	Id                *string                    `json:"id,omitempty"`
	Name              *string                    `json:"name,omitempty"`
	Desc              *string                    `json:"desc,omitempty"`
	BackendPolicyList []*AppIpGroupBackendPolicy `json:"backendPolicyList,omitempty"`
}
