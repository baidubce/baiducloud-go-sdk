package bls

type Notice struct {
	Id           *string        `json:"id,omitempty"`
	Name         *string        `json:"name,omitempty"`
	Members      []*string      `json:"members,omitempty"`
	Methods      []*string      `json:"methods,omitempty"`
	Callbacks    []*string      `json:"callbacks,omitempty"`
	DisableTimes []*DisableTime `json:"disableTimes,omitempty"`
}
