package bci

type Capabilities struct {
	Add  []*string `json:"add,omitempty"`
	Drop []*string `json:"drop,omitempty"`
}
