package scs

type EntranceItem struct {
	Ip   *string `json:"ip,omitempty"`
	Port *int32  `json:"port,omitempty"`
	Zone *string `json:"zone,omitempty"`
}
