package bcc

type OsModel struct {
	InstanceId     *string `json:"instanceId,omitempty"`
	OsArch         *string `json:"osArch,omitempty"`
	OsName         *string `json:"osName,omitempty"`
	OsVersion      *string `json:"osVersion,omitempty"`
	OsType         *string `json:"osType,omitempty"`
	OsLang         *string `json:"osLang,omitempty"`
	SpecialVersion *string `json:"specialVersion,omitempty"`
}
