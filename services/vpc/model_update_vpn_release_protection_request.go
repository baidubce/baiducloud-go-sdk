package vpc

type UpdateVpnReleaseProtectionRequest struct {
	VpnId         *string `json:"-"`
	ClientToken   *string `json:"-"`
	DeleteProtect *bool   `json:"deleteProtect,omitempty"`
}
