package vpc

type UpdateIpv6GatewayReleaseProtectionSwitchRequest struct {
	GatewayId     *string `json:"-"`
	ClientToken   *string `json:"-"`
	DeleteProtect *int32  `json:"deleteProtect,omitempty"`
}
