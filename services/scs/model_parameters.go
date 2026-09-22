package scs

type Parameters struct {
	ConfName         *string `json:"confName,omitempty"`
	ConfDefault      *string `json:"confDefault,omitempty"`
	ConfValue        *string `json:"confValue,omitempty"`
	ConfType         *int32  `json:"confType,omitempty"`
	ConfRange        *string `json:"confRange,omitempty"`
	ConfModule       *int32  `json:"confModule,omitempty"`
	ConfDesc         *string `json:"confDesc,omitempty"`
	NeedReboot       *int32  `json:"needReboot,omitempty"`
	ConfRedisVersion *string `json:"confRedisVersion,omitempty"`
	ConfCacheVersion *int32  `json:"confCacheVersion,omitempty"`
	ConfUserVisible  *int32  `json:"confUserVisible,omitempty"`
}
