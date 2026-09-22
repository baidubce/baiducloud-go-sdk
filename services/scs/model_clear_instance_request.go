package scs

type ClearInstanceRequest struct {
	InstanceId     *string `json:"-"`
	Password       *string `json:"password,omitempty"`
	DbIndex        *int32  `json:"dbIndex,omitempty"`
	IsFlushExpired *bool   `json:"isFlushExpired,omitempty"`
	IsDefer        *bool   `json:"isDefer,omitempty"`
}
