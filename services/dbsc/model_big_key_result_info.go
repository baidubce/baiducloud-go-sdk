package dbsc

type BigKeyResultInfo struct {
	Db           *int32  `json:"db,omitempty"`
	ElementCount *int32  `json:"elementCount,omitempty"`
	Encoding     *string `json:"encoding,omitempty"`
	ExpireTime   *string `json:"expireTime,omitempty"`
	Key          *string `json:"key,omitempty"`
	Size         *int32  `json:"size,omitempty"`
	DbscType     *string `json:"type,omitempty"`
}
