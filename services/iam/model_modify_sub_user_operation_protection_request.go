package iam

type ModifySubUserOperationProtectionRequest struct {
	UserName   *string `json:"userName,omitempty"`
	EnabledMfa *bool   `json:"enabledMfa,omitempty"`
	MfaType    *string `json:"mfaType,omitempty"`
}
