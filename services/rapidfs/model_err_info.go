package rapidfs

type ErrInfo struct {
	ErrCode *string `json:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty"`
}
