package vpc

type UpdateDeleteProtectRequest struct {
	GatewayId     *string `json:"-"`
	ClientToken   *string `json:"-"`
	DeleteProtect *int32  `json:"deleteProtect,omitempty"`
}
