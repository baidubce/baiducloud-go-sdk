package bls

type Error struct {
	RootCause []*Error `json:"root_cause,omitempty"`
	BlsType   *string  `json:"type,omitempty"`
	Reason    *string  `json:"reason,omitempty"`
}
