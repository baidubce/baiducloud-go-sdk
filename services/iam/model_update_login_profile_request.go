package iam

type UpdateLoginProfileRequest struct {
	UserName          *string `json:"-"`
	Password          *string `json:"password,omitempty"`
	NeedResetPassword *bool   `json:"needResetPassword,omitempty"`
	EnabledLogin      *bool   `json:"enabledLogin,omitempty"`
	EnabledLoginMfa   *bool   `json:"enabledLoginMfa,omitempty"`
	LoginMfaType      *string `json:"loginMfaType,omitempty"`
	ThirdPartyType    *string `json:"thirdPartyType,omitempty"`
	ThirdPartyAccount *string `json:"thirdPartyAccount,omitempty"`
}
