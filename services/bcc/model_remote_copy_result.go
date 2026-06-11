package bcc

type RemoteCopyResult struct {
	Region  *string `json:"region,omitempty"`
	ImageId *string `json:"imageId,omitempty"`
	Code    *string `json:"code,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty"`
}
