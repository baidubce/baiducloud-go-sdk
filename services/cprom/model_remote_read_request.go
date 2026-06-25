package cprom

type RemoteReadRequest struct {
	RemoteReadUrl *string `json:"-"`
	Query         *string `json:"query,omitempty"`
	Step          *int32  `json:"step,omitempty"`
	Start         *int64  `json:"start,omitempty"`
	End           *int64  `json:"end,omitempty"`
}
