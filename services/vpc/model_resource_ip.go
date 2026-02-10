package vpc

type ResourceIp struct {
	Ip           *string `json:"ip,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}
