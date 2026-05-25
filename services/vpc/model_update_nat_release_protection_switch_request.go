package vpc

type UpdateNatReleaseProtectionSwitchRequest struct {
	NatId         *string `json:"-"`
	ClientToken   *string `json:"-"`
	DeleteProtect *bool   `json:"deleteProtect,omitempty"`
}
