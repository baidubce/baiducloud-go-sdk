package as

type MonitorObject struct {
	AsType    *string           `json:"type,omitempty"`
	Names     []*string         `json:"names,omitempty"`
	Resources []*PolicyResource `json:"resources,omitempty"`
	TypeName  *string           `json:"typeName,omitempty"`
}
