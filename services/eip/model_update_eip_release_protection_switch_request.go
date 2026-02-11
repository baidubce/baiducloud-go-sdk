package eip

type UpdateEipReleaseProtectionSwitchRequest struct {
	Eip           *string `json:"-"`
	ClientToken   *string `json:"-"`
	DeleteProtect *bool   `json:"deleteProtect,omitempty"`
}
