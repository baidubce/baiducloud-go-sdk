package eip

type UpdateEipDeleteProtectRequest struct {
	Eip           *string `json:"-"`
	ClientToken   *string `json:"-"`
	DeleteProtect *bool   `json:"deleteProtect,omitempty"`
}
